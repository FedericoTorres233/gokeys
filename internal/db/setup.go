package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func SetupDB() {
    //os.Remove("bin/passwd.db")

    // Open a connection to the SQLite database
    db, err := sql.Open("sqlite3", "bin/passwd.db")
    if err != nil {
        panic(err)
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
        log.Println(err)
        os.Exit(1)
    }
}
