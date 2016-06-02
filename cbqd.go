package cbqd

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

type AccessCreds struct {
	Dkey  string `json:"username"`
	Dpass string `json:"password"`
}

type Database struct {
	Ukey       AccessCreds
	Host, Port string
}

var (
	dbflag  = flag.String("db", "mysql", "Database type to dump.")
	csflag  = flag.String("cs", "aws", "S3 storage repository to use.")
	kvflag  = flag.Bool("kv", false, "Access vault to acquire secrets.")
	dhflag  = flag.String("dh", "127.0.0.1", "Host IP for the database to be backuped up.")
	dpflag  = flag.String("dp", "3306", "Database port for access.")
	vpflag  = flag.Bool("V", false, "Print the version number.")
	version = formattedVersion()
)

func (a AccessCreds) GetCreds(vbackend string, inout string, kvault bool) (AccessCreds, error) {
	un := inout + "_KEY"
	up := inout + "_PASS"
	ac := AccessCreds{}

	if kvault == false {
		ac.Dkey = os.Getenv(un)
		ac.Dpass = os.Getenv(up)
		if ac.Dkey == "" && ac.Dpass == "" {
			log.Fatalln(OS_ENVIRONMENT_UNSET)
		}
		return ac, nil
	}
	return ac, VAULT_CREDENTIAL_ERROR
}

func usage() {
	//
}

func init() {
	flag.Parse()
}

func Cbqd() {
	increds, err := new(AccessCreds).GetCreds(*dbflag, "CBQD_IN", *kvflag)
	if err != nil {
		log.Fatalln(err)
	}

	db := Database{increds, *dhflag, *dpflag}
	outcreds, err := new(AccessCreds).GetCreds(*csflag, "CBQD_OUT", *kvflag)
	if err != nil {
		log.Fatalln(err)
	}

	topdir, err := ioutil.TempDir("", "cbqd_state")
	if err != nil {
		log.Fatalln(err)
	}

	defer os.RemoveAll(topdir)
	tmpfhandle := filepath.Join(topdir, "tmpfile")

	bname, err := MYSQL{}.DBdump(db, tmpfhandle)
	if err != nil {
		log.Fatalln(err)
	}

	err0 := AWS{}.CloudSend(outcreds, bname, tmpfhandle)
	if err0 != nil {
		log.Fatalln(BACKUP_UPLOAD_ERROR)
	}

}
