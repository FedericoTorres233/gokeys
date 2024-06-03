package main

import (
	"github.com/federicotorres233/gokeys/cmd/cli"
	"github.com/federicotorres233/gokeys/internal/db"
)

func main() {
	db.SetupDB()
	cli.Execute()
}
