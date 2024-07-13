package setup

import (
	"os"
	"path/filepath"

	"github.com/federicotorres233/gokeys/internal/crypto"
	"github.com/federicotorres233/gokeys/internal/utils"
)

func Install() {

	// Read master password
	master, err := utils.ReadMasterPassword()
	if err != nil {
		utils.LogError(err)
		os.Exit(1)
	}

	// Generate key & salt
	key, err := crypto.GenerateNewKey(master)
	if err != nil {
		utils.LogError(err)
		os.Exit(1)
	}

	// Get default filepaths for databases
	def := utils.LoadConfigs()
	enc_db := filepath.Join(utils.GetBaseDir(), def["encrypted_db"])
	tmp_db := def["temporary_db"]

	// Set up db
	err = setupDB(key, tmp_db, enc_db)
	if err != nil {
		utils.LogError(err)
		os.Exit(1)
	}
}
