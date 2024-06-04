package main

import (
	"github.com/federicotorres233/gokeys/cmd/cli"
	"github.com/federicotorres233/gokeys/internal/db"
	"github.com/federicotorres233/gokeys/pkg/utils"
)

const logdir string = "bin/logs"

func main() {
	utils.SetupLogger(logdir)
	db.SetupDB()
	cli.Execute()
}
