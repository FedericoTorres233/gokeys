package db

import (
	"database/sql"

	"github.com/federicotorres233/gokeys/internal/types"
	_ "github.com/mattn/go-sqlite3"
)

// Function to query and print rows for a specific website
func AddRecord(db *sql.DB, record *types.Record) error {
	querySQL := `INSERT INTO gokeys
	(website, username, password, notes, tag, url, favorite, status)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(querySQL,
		&record.Website,
		&record.Username,
		&record.Password,
		&record.Notes,
		&record.Tag,
		&record.Url,
		&record.Favorite,
		&record.Status)
	if err != nil {
		return err
	}

	return nil
}
