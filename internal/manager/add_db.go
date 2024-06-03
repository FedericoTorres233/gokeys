package manager

import (
	"database/sql"

	"github.com/federicotorres233/gokeys/internal/db"
	"github.com/federicotorres233/gokeys/types"
	_ "github.com/mattn/go-sqlite3"
)

// AddPassword adds a password for a specific site and username
func (pm *PasswordManager) AddPassword(record types.Record) error {
	
	// Open a connection to the SQLite database
	database, err := sql.Open("sqlite3", "bin/passwd.db")
	if err != nil {
		return err
	}
	defer database.Close()
	
	// Get the id depending on the site
	err = db.AddRecord(database, record)
	if err != nil {
		return err
	}

	return nil
}
