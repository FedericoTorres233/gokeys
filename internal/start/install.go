package start

import (
	"log"

	"github.com/federicotorres233/gokeys/pkg/utils"
)

func Install() {
	dbdir := "bin/passwd.db"
	tmpdir := "/tmp/db.db"

	// Read master password
	master, err := utils.ReadPassword()
	if err != nil {
		log.Println("[ERROR] ", err)
	}

	// Generate key & salt
	key, err := utils.GenerateKey(master)
	if err != nil {
		log.Println("[ERROR] ", err)
	}

	// Set up DB
	err = setupDB(key, tmpdir, dbdir)
	if err != nil {
		log.Println("[ERROR] ", err)
	}

	// Decrypt db
	err = utils.DbDecrypt()
	if err != nil {
		log.Println("[ERROR] ", err)
	}

}
