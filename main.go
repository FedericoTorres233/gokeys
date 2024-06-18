package main

import (
	"github.com/federicotorres233/gokeys/cmd"
	"github.com/federicotorres233/gokeys/internal/utils"
)

func main() {
	utils.SetupLogger()
	cmd.Execute()
}
