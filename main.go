package main

import (
	"github.com/federicotorres233/gokeys/cmd"
	"github.com/federicotorres233/gokeys/internal/db"
)

func main() {
	db.SetupDB()
	cmd.Execute()
}
