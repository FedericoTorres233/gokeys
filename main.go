package main

import (
	"github.com/federicotorres233/gokeys/cmd"
	"github.com/federicotorres233/gokeys/internal/utils"
)

func main() {

	_ = utils.NewLogger("tmp/logs", "gokeys.log")

	cmd.Execute()
}
