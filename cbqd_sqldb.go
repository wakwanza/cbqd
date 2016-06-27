package cbqd

import (
	gp "github.com/maxwellhealth/go-gpg"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
)

var (
	dbname = os.Getenv("CBQD_DB_NAME")
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
	return "nil"
}

//Encrypyt the database dump
func EncryptDBdump(dbtxt string, dbgpg string) error {
	pubkey, err0 := ioutil.ReadFile(*enflag)
	sourcedata, err := os.OpenFile(dbtxt, os.O_RDONLY, 0660)
	if err != nil {
		log.Error(err)
	}
	encryptdata, err := os.OpenFile(dbgpg, os.O_RDWR|os.O_CREATE, 0660)
	if err != nil {
		log.Error(err)
	}
	if err0 != nil {
		return GPG_KEY_ERROR
	}
	err2 := gp.Encode(pubkey, sourcedata, encryptdata)
	if err2 != nil {
		return err2
	}
	return nil
}

//Take a data snapshot from the specified database
func (a SQLDB) DBdump(d Database, tmpdir string) (string, error) {
	tstamp := time.Now().UTC().Format(time.RFC3339)
	objname := "CBQD_DB_" + tstamp + ".sql"
	gpgname := "CBQD_DB_" + tstamp + ".gpg"
	err := os.Chdir(tmpdir)
	if err != nil {
		return "", BACKUP_FOLDER_ERROR
	}
	log.Info("initiate database data dump process....")
	if err = exec.Command(MakeCommandString(d), " > ", objname).Run(); err != nil {
		_, errm := exec.LookPath("mysqldump")
		if errm != nil {
			return "", DB_DUMP_ERROR_EXEC
		}
		return "", DB_DUMP_ERROR
	}
	log.Info("database dump complete......begin data encryption process.")

	err0 := EncryptDBdump(objname, gpgname)
	if err0 == nil {
		log.Info("data encryption complete")
	} else {
		log.Info("data encryption could not be completed.")
	}
	return gpgname, err0
}
