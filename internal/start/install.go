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

	// Set up DB
	err := setupDB(dbdir)
	if err != nil {
		log.Println("[ERROR] ", err)
	}

	// Read master password
	master, err = utils.ReadPassword()
	if err != nil {
		log.Println("[ERROR] ", err)
	}

	// Generate key & salt
	key, err := utils.GenerateKey(master)
	if err != nil {
		log.Println(err)
	}

	// Encrypt db
	utils.DbEncrypt(key, dbdir)
	utils.DbEncrypt(master, dbdir)
}
