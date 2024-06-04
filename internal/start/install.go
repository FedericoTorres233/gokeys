package start

import (
	"log"

	"github.com/federicotorres233/gokeys/pkg/utils"
)

var (
	master string
	dbdir  string
)

func Install() {
	dbdir = "bin/passwd.db"

	err := setupDB(dbdir)
	if err != nil {
		log.Println("[ERROR] ", err)
	}


	master, err = utils.ReadPassword()
	if err != nil {
		log.Println("[ERROR] ", err)
	}

	// Encrypt db
	utils.DbEncrypt(master, dbdir)
}
