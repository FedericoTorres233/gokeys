package query

import (
	"database/sql"
	"log"

	"github.com/federicotorres233/gokeys/internal/types"
	_ "github.com/mattn/go-sqlite3"
)

// Function to query and print rows for a specific website
func QueryByWebsite(db *sql.DB, website string) (types.Record, error) {
	var id int
	var password, username string
	var record types.Record

	querySQL := `SELECT id, password, website, username FROM gokeys WHERE website = ?`
	rows, err := db.Query(querySQL, website)
	if err != nil {
		return types.Record{}, err
	}
	defer rows.Close()

	log.Printf("Results for website: %s\n", website)
	for rows.Next() {
		err = rows.Scan(&id, &password, &website, &username)
		if err != nil {
			panic(err)
		}
		log.Printf("ID: %d, Password: %s, Website: %s, Username: %s\n", id, password, website, username)
		record.Password = password
		record.Website = website
		record.Username = username
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		return record, err
	}

	return record, nil
}
