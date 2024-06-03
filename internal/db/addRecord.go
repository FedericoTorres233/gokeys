package db

import (
	"database/sql"

	"github.com/federicotorres233/gokeys/types"
	_ "github.com/mattn/go-sqlite3"
)

// Function to query and print rows for a specific website
func AddRecord(db *sql.DB, record types.Record) error {

	querySQL := `INSERT INTO gokeys (website, username, password) VALUES (?, ?, ?)`
	_, err := db.Exec(querySQL, record.Website, record.Username, record.Password)
	if err != nil {
		return err
	}

	return nil
}
