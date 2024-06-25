package db

import (
	"database/sql"
	"fmt"

	"github.com/federicotorres233/gokeys/internal/types"
	_ "github.com/mattn/go-sqlite3"
)

// Function to query and print rows for a specific website
func GetRecords(db *sql.DB, record *types.Record) error {
	var id int

	whereSQL := fmt.Sprintf("WHERE website = '%s' OR username = '%s'", record.Website, record.Username)
	querySQL := fmt.Sprintf("SELECT id, password, website, username, notes, tag, url, favorite, status FROM gokeys %s;", whereSQL)
	rows, err := db.Query(querySQL)
	if err != nil {
		return err
	}
	defer rows.Close()

	// fmt.Printf("Results for website: %s\n", record.Website)
	for rows.Next() {
		err = rows.Scan(
			&id,
			&record.Password,
			&record.Website,
			&record.Username,
			&record.Notes,
			&record.Tag,
			&record.Url,
			&record.Favorite,
			&record.Status,
		)
		if err != nil {
			panic(err)
		}
		// log.Printf("ID: %d, Password: %s, Website: %s, Username: %s\n", id, password, website, username)
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}
