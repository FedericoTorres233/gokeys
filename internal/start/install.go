package start

import (
	"log"

	"github.com/federicotorres233/gokeys/pkg/utils"
)

func Install() {

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

	// Set up DB & get default configs
	def := utils.LoadConfigs()
	err = setupDB(key, def["temporary_db"], def["encrypted_db"])
	if err != nil {
		log.Println("[ERROR] ", err)
	}
}
