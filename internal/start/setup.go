package start

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func setupDB(dbdir string) error {
	//os.Remove("bin/passwd.db")

	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", dbdir)
	if err != nil {
		return err
	}
	defer db.Close()

	// Create a table
	createTable := `CREATE TABLE IF NOT EXISTS gokeys (
        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "website" TEXT NOT NULL UNIQUE,
        "password" TEXT NOT NULL,
        "username" TEXT
    );`
	_, err = db.Exec(createTable)
	if err != nil {
		return err
	}

	log.Println("INFO: db set up successfully at", dbdir)
    return nil
}
