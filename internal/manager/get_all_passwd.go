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
func (pm *PasswordManager) GetAllPasswords() ([]types.Record, error) {

	// Decrypt database
	err := crypto.DbDecrypt(pm.tmpdb)
	if err != nil {
		return nil, err
	}

	// Open a connection to the SQLite database
	database, err := sql.Open("sqlite3", pm.tmpdb)
	if err != nil {
		return nil, err
	}
	defer os.Remove(pm.tmpdb)
	defer database.Close()

	// Get the id depending on the site (passing the pointer to record)
	r, err := db.GetAllRecords(database)
	if err != nil {
		return nil, err
	}

	return r, errors.New("sorry! password not found")
}
