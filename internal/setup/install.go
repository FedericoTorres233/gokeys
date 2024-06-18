package setup

import (
	"log"
	"path/filepath"

	"github.com/federicotorres233/gokeys/internal/crypto"
	"github.com/federicotorres233/gokeys/internal/utils"
)

func Install() {

	// Read master password
	master, err := utils.ReadPassword()
	if err != nil {
		log.Println("[ERROR]", err)
	}

	// Generate key & salt
	key, err := crypto.GenerateKey(master)
	if err != nil {
		log.Println("[ERROR]", err)
	}

	// Get default filepaths for databases
	def := utils.LoadConfigs()
	enc_db := filepath.Join(utils.GetBaseDir(), def["encrypted_db"])
	tmp_db := def["temporary_db"]

	// Set up db
	err = setupDB(key, tmp_db, enc_db)
	if err != nil {
		log.Println("[ERROR]", err)
	}
}
