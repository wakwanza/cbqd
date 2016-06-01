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

type MYSQL struct {
}

//Create command string to initiate backup of data
func MakeCommandString(a Database) string {
	if dbname == "" {
		dbname = " -A "
	}
	return "mysqldump --single-transaction -q  -u " + a.Ukey.Dkey + " -p" + a.Ukey.Dpass + " -h " + a.Host + " -P " + a.Port + " --database " + dbname + " "
}

//Create command string to encrypt data dump and remove unecncrypted data
func MakeEncryptString(gpgid string) string {
	return "gpg -W -e --yes -r " + gpgid + " "
}

//Take a data snapshot from the specified database
func (a MYSQL) DBdump(d Database, tmpdir string) (string, error) {
	objname := "CBQD_DB_" + time.Now().UTC().Format(time.RFC3339) + ".sql"
	err := os.Chdir(tmpdir)
	if err != nil {
		return "", BACKUP_FOLDER_ERROR
	}
	if err := exec.Command(MakeCommandString(d), " > ", objname).Run(); err != nil {
		return "", err
	}
	if identity != "" {
		if err := exec.Command(MakeEncryptString(identity), objname).Run(); err != nil {
			return "", err
		}
		objnameg := objname + ".gpg"
		return objnameg, nil
	}
	return objname, nil
}
