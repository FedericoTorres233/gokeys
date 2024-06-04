package manager

import (
	"errors"

	"database/sql"

	"github.com/federicotorres233/gokeys/internal/db"
	"github.com/federicotorres233/gokeys/internal/types"
	_ "github.com/mattn/go-sqlite3"
)

// GetPassword retrieves a password for a specific site and username
func (pm *PasswordManager) GetPassword(record *types.Record) error {

	// Open a connection to the SQLite database
	database, err := sql.Open("sqlite3", "bin/passwd.db")
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
