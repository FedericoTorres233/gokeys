package manager

import (
	"database/sql"
	"errors"
	"os"

	"github.com/federicotorres233/gokeys/internal/db"
	"github.com/federicotorres233/gokeys/internal/types"
	"github.com/federicotorres233/gokeys/pkg/utils"
	_ "github.com/mattn/go-sqlite3"
)

// GetPassword retrieves a password for a specific site and username
func (pm *PasswordManager) GetPassword(record *types.Record) error {

	// Decrypt database
	err := utils.DbDecrypt(pm.tmpdb)
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

	// Open a connection to the SQLite database
	database, err = sql.Open("sqlite3", pm.tmpdb)
	if err != nil {
		return err
	}
	defer database.Close()

	// Get the id depending on the site (passing the pointer to record)
	err = db.GetRecordByWebsite(database, record)
	if err != nil {
		return err
	}

	// Check if record is not empty
	if record.Password != "" {
		return nil
	}

	return errors.New("sorry! password not found")
}
