package db

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func SetDB() {
    os.Remove("bin/passwd.db")

    // Open a connection to the SQLite database
    db, err := sql.Open("sqlite3", "bin/passwd.db")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    // Create a table
    createTable := `CREATE TABLE IF NOT EXISTS gokeys (
        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "password" TEXT NOT NULL,
        "website" TEXT NOT NULL,
        "username" TEXT
    );`
    _, err = db.Exec(createTable)
    if err != nil {
        panic(err)
    }

    // Insert data into the table
    insertUserSQL := `INSERT INTO gokeys (password, website) VALUES (?, ?)`
    _, err = db.Exec(insertUserSQL, "1234", "github.com")
    if err != nil {
        panic(err)
    }
    _, err = db.Exec(insertUserSQL, "5678", "minecraft.net")
    if err != nil {
        panic(err)
    }
}
