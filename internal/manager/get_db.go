package manager

import (
	"errors"

	"database/sql"

	"github.com/federicotorres233/gokeys/internal/db"
	_ "github.com/mattn/go-sqlite3"
)

// GetPassword retrieves a password for a specific site and username
func (pm *PasswordManager) GetPassword(website, username string) (string, error) {

	// Open a connection to the SQLite database
	database, err := sql.Open("sqlite3", "bin/passwd.db")
	if err != nil {
		panic(err)
	}
	defer database.Close()

	// Get the id depending on the site
	record, err := db.QueryByWebsite(database, website)
	if err != nil {
		return "", err
	}

	// Check if record is not empty
	if record.Password != "" {
		return record.Password, nil
	}

	return "", errors.New("sorry! password not found")
}
