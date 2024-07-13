package manager

import (
	"database/sql"
	"os"

	"github.com/federicotorres233/gokeys/internal/crypto"
	db "github.com/federicotorres233/gokeys/internal/database"
	"github.com/federicotorres233/gokeys/internal/types"
	"github.com/federicotorres233/gokeys/internal/utils"
	_ "github.com/mattn/go-sqlite3"
)

// AddPassword adds a password for a specific site and username
func (pm *PasswordManager) AddPassword(record *types.Record) error {

	// Decrypt database
	err := crypto.DbDecrypt(pm.tmpdb)
	if err != nil {
		return err
	}

	// Read password from stdin
	p, err := utils.ReadPassword()
	if err != nil {
		utils.LogError(err)
		return err
	}
	record.Password = p

	// Open a connection to the SQLite database
	database, err := sql.Open("sqlite3", pm.tmpdb)
	if err != nil {
		return err
	}
	defer os.Remove(pm.tmpdb)
	defer database.Close()

	// Get the id depending on the site
	err = db.AddRecord(database, record)
	if err != nil {
		return err
	}

	// Copy an encrypted version of the database
	key := make([]byte, 32)
	crypto.GetKey(&key)
	err = crypto.DbEncrypt(key, pm.encdb, pm.tmpdb)
	if err != nil {
		return err
	}

	return nil
}
