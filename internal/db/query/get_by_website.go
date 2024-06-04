package query

import (
	"database/sql"

	"github.com/federicotorres233/gokeys/internal/types"
	_ "github.com/mattn/go-sqlite3"
)

// Function to query and print rows for a specific website
func QueryByWebsite(db *sql.DB, record *types.Record) error {
	var id int

	querySQL := `SELECT id, password, website, username FROM gokeys WHERE website = ?`
	rows, err := db.Query(querySQL, record.Website)
	if err != nil {
		return err
	}
	defer rows.Close()

	//fmt.Printf("Results for website: %s\n", record.Website)
	for rows.Next() {
		err = rows.Scan(&id, &record.Password, &record.Website, &record.Username)
		if err != nil {
			panic(err)
		}
		//log.Printf("ID: %d, Password: %s, Website: %s, Username: %s\n", id, password, website, username)
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}
