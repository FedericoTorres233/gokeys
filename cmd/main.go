package main

import (
	"github.com/federicotorres233/gokeys/cmd/cli"
	"github.com/federicotorres233/gokeys/pkg/utils"
)

func main() {
	utils.SetupLogger()
	cli.Execute()
}
