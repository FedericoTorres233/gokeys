package start

import (
	"database/sql"
	"log"
	"os"

	"github.com/federicotorres233/gokeys/pkg/crypto"
	_ "github.com/mattn/go-sqlite3"
)

func setupDB(key []byte, tmpdb string, encdb string) error {
	// os.Remove("bin/passwd.db")

	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", tmpdb)
	if err != nil {
		return err
	}
	defer db.Close()

	// Create a table
	createTable := `CREATE TABLE IF NOT EXISTS gokeys (
        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "website" TEXT NOT NULL UNIQUE,
        "password" TEXT NOT NULL,
        "username" TEXT,
		"notes" TEXT,
		"tag" TEXT,
		"url" TEXT,
		"favorite" BOOL,
		"status" BOOL
    );`
	_, err = db.Exec(createTable)
	if err != nil {
		return err
	}

	log.Println("[INFO] DB set up successfully at", tmpdb)

	// DB encryption
	log.Println("[INFO] Encrypting database")
	err = crypto.DbEncrypt(key, encdb, tmpdb)
	if err != nil {
		return err
	}

	// Remove temporary DB
	os.Remove(tmpdb)

	return nil
}
