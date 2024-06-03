package cli

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gokeys",
	Short: "Simple password/key manager written in Go",
	Long:  `Simple password/key manager written in Go. Your credentials are stored using post-quantum encryption`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func init() {
	/*
		// Define flags and configs. Cobra supports persistent flags, which, if defined here,
		// will be global for your application.

		// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gokeys.yaml)")

		// Cobra also supports local flags, which will only run
		// when this action is called directly.
		rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	*/
}
