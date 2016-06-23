package cbqd

import (
	"os"
	"os/exec"
	"time"
)

var (
	dbname   = os.Getenv("CBQD_DB_NAME")
	identity = os.Getenv("CBQD_GPG_ID")
)

type SQLDB struct {
}

//Create command string to initiate backup of data
func MakeCommandString(a Database) string {
	if dbname == "" {
		dbname = "-A"
	}
	switch a.Dtype {
	case "mysql":
		return "mysqldump --single-transaction -q  -u " + a.Ukey.Dkey + " -p" + a.Ukey.Dpass + " -h " + a.Host + " -P " + a.Port + " --database " + dbname + " "
	default:
		log.Error(DB_TYPE_ERROR)
	}

}

//Create command string to encrypt data dump and remove unecncrypted data
func MakeEncryptString(gpgid string) string {
	return "gpg2 -e --trust-model always -R " + gpgid + " "
}

//Take a data snapshot from the specified database
func (a SQLDB) DBdump(d Database, tmpdir string) (string, error) {
	objname := "CBQD_DB_" + time.Now().UTC().Format(time.RFC3339) + ".sql"
	err := os.Chdir(tmpdir)
	if err != nil {
		return "", BACKUP_FOLDER_ERROR
	}
	log.Info("initiate database data dump process.")
	if err = exec.Command(MakeCommandString(d), " > ", objname).Run(); err != nil {
		_, errm := exec.LookPath("mysqldump")
		if errm != nil {
			return "", DB_DUMP_ERROR_EXEC
		}
		return "", DB_DUMP_ERROR
	}
	if identity != "" {
		log.Info("begin data encryption process.")
		if err = exec.Command(MakeEncryptString(identity), objname).Run(); err != nil {
			return "", err
		}
		objnameg := objname + ".gpg"
		return objnameg, nil
	}
	return objname, nil
}
