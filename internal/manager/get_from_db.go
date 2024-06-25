package manager

import (
	"database/sql"
	"errors"
	"os"

	"github.com/federicotorres233/gokeys/internal/crypto"
	db "github.com/federicotorres233/gokeys/internal/database"
	"github.com/federicotorres233/gokeys/internal/types"
	_ "github.com/mattn/go-sqlite3"
)

// GetPassword retrieves a password for a specific site and username
func (pm *PasswordManager) GetPassword(record *types.Record) error {

	// Decrypt database
	err := crypto.DbDecrypt(pm.tmpdb)
	if err != nil {
		return err
	}

	// Open a connection to the SQLite database
	database, err := sql.Open("sqlite3", pm.tmpdb)
	if err != nil {
		return err
	}
	defer os.Remove(pm.tmpdb)
	defer database.Close()

	// Get the id depending on the site (passing the pointer to record)
	err = db.GetRecords(database, record)
	if err != nil {
		return err
	}

	// Check if record is not empty
	if record.Password != "" {
		return nil
	}

	return errors.New("sorry! password not found")
}
