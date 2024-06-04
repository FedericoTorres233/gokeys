package cli

import (
	"fmt"
	"log"

	"github.com/federicotorres233/gokeys/internal/manager"
	"github.com/spf13/cobra"
)

var (
  getWebsite  string
  getUsername string
)

// getCmd represents the get command
var getCmd = &cobra.Command{
  Use:   "get",
  Short: "Get a password",
  Run: func(cmd *cobra.Command, args []string) {
    pm := manager.NewPasswordManager()
    password, err := pm.GetPassword(getWebsite, getUsername)
    if err != nil {
      log.Println("ERROR: ", err)
      return
    }

    if getUsername == ""{
      fmt.Printf("Password for website %s: %s\n", getWebsite, password)
      return
    }
    fmt.Printf("Password for %s at %s: %s\n", getUsername, getWebsite, password)

  },
}

func init() {
  // Add command
  rootCmd.AddCommand(getCmd)

  // Set up flags
  getCmd.Flags().StringVarP(&getWebsite, "website", "w", "", "Website [required]")
  getCmd.Flags().StringVarP(&getUsername, "username", "u", "", "Username")

  // Mark required flags
  getCmd.MarkFlagRequired("website")
}
