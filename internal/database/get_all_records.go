package db

import (
	"database/sql"

	"github.com/federicotorres233/gokeys/internal/types"
	_ "github.com/mattn/go-sqlite3"
)

// Function to query and print rows for a specific website
func GetAllRecords(db *sql.DB) ([]types.Record, error) {

	rows, err := db.Query("SELECT id, password, website, username, notes, tag, url, favorite, status FROM gokeys;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []types.Record
	for rows.Next() {
		var id int
		var record types.Record

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
		records = append(records, record)
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return records, nil
}
