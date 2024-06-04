package start

import (
	"database/sql"
	"log"
	"os"

	"github.com/federicotorres233/gokeys/pkg/utils"
	_ "github.com/mattn/go-sqlite3"
)

func setupDB(key []byte, tmpdir string, dbdir string) error {
	//os.Remove("bin/passwd.db")

	// Open a connection to the SQLite database
	db, err := sql.Open("sqlite3", tmpdir)
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

	log.Println("[INFO] DB set up successfully at", tmpdir)

	// DB encryption
	log.Println("[INFO] Encrypting database")
	err = utils.DbEncrypt(key, dbdir, tmpdir)
	if err != nil {
		log.Println("[ERROR] ", err)
	}

	// Remove temporary DB
	os.Remove(tmpdir)

    return nil
}
