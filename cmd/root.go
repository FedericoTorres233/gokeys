package cmd

import (
	"fmt"
	"os"

	"github.com/federicotorres233/gokeys/internal/setup"
	"github.com/federicotorres233/gokeys/internal/types"
	"github.com/federicotorres233/gokeys/internal/utils"
	"github.com/spf13/cobra"
)

var (
	record       types.Record
	setClipboard bool
	installed    bool
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gokeys",
	Short: "Simple password/key manager written in Go",
	Long:  `Simple password/key manager written in Go. To start using it, first run the command using the --install flag, which generates an encrypted database`,
	Run: func(cmd *cobra.Command, args []string) {

		// Load ascii logo
		fmt.Println(string(utils.LoadAsciiArt()))
		fmt.Println("Welcome to gokeys!")

		if installed {
			fmt.Println("Performing installation of GoKeys...")
			setup.Install()
			return
		} else {
			cmd.Help()
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		utils.LogError(err)
		os.Exit(1)
	}
}

func init() {
	// Define flags and configs. Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gokeys.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolVarP(&installed, "install", "i", false, "install database")
}
