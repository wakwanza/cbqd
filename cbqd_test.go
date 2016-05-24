package cbqd

import (
	"fmt"
	"testing"
)

func TestCreds(t *testing.T) {
	_, err := new(AccessCreds).GetCreds(*dbflag, "CBQD_IN", *kvflag)
	if err != nil {
		fmt.Println(err)
	}

	_, err1 := new(AccessCreds).GetCreds(*csflag, "CBQD_OUT", *kvflag)
	if err1 != nil {
		fmt.Println(err1)
	}
}
