package cbqd

import (
	"flag"
	"github.com/wakwanza/go-logging"
	"io/ioutil"
	"os"
)

type AccessCreds struct {
	Dkey  string `json:"username"`
	Dpass string `json:"password"`
}

type Database struct {
	Ukey       AccessCreds
	Host, Port string
	Dtype      string
}

var (
	dbflag = flag.String("db", "mysql", "`Database type` to dump.")
	csflag = flag.String("cs", "aws", "S3 `storage repository` to use.")
	kvflag = flag.Bool("kv", false, "Access `vault` to acquire secrets.")
	dhflag = flag.String("dh", "127.0.0.1", "`Host IP` for the database to be backed up.")
	dpflag = flag.String("dp", "3306", "`Database port` for access.")
	log    = logging.MustGetLogger("cbqd")
	bl1    = logging.NewLogBackend(os.Stderr, "", 0)
	blf    = logging.NewBackendFormatter(bl1, lgform)
	lgform = logging.MustStringFormatter(`%{color}%{time:15:04:05.000} %{shortpkg} ▶ %{level:.4s} %{id:03x}%{color:reset} %{message}`)
)

func (a AccessCreds) GetCreds(vbackend string, inout string, kvault bool) (AccessCreds, error) {
	un := inout + "_KEY"
	up := inout + "_PASS"
	ac := AccessCreds{}
	log.Info("acquire access credentials for databse and s3 store.")
	if kvault == false {
		ac.Dkey = os.Getenv(un)
		ac.Dpass = os.Getenv(up)
		if ac.Dkey == "" && ac.Dpass == "" {
			log.Fatal(OS_ENVIRONMENT_UNSET)
		}
		return ac, nil
	}
	return ac, VAULT_CREDENTIAL_ERROR
}

func init() {
	flag.Parse()
	logging.SetBackend(bl1, blf)
}

func Cbqd() {
	increds, err := new(AccessCreds).GetCreds(*dbflag, "CBQD_IN", *kvflag)
	if err != nil {
		log.Fatal(err)
	}

	db := Database{increds, *dhflag, *dpflag, *dbflag}
	outcreds, err0 := new(AccessCreds).GetCreds(*csflag, "CBQD_OUT", *kvflag)
	if err0 != nil {
		log.Fatal(err0)
	}

	topdir, err1 := ioutil.TempDir("", "cbqd_db_state")
	if err1 != nil {
		log.Fatal(err1)
	}

	defer os.RemoveAll(topdir)

	bname, err2 := SQLDB{}.DBdump(db, topdir)
	if err2 != nil {
		log.Fatal(err2)
	}

	cstore := S3C{*csflag}
	err3 := cstore.CloudSend(outcreds, bname, topdir)
	if err3 != nil {
		log.Fatal(BACKUP_UPLOAD_ERROR)
	}

}
