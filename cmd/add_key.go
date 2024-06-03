package cmd

import (
  "fmt"
  "github.com/federicotorres233/gokeys/internal/manager"
  "github.com/spf13/cobra"
)

var (
  site     string
  username string
  password string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
  Use:   "add",
  Short: "Add a new password",
  Run: func(cmd *cobra.Command, args []string) {
    pm := manager.NewPasswordManager()
    pm.AddPassword(site, username, password)
    fmt.Printf("Password for %s at %s added successfully!\n", username, site)
  },
}

func init() {
  // Add command
  rootCmd.AddCommand(addCmd)

  // Set up flags
  addCmd.Flags().StringVarP(&site, "site", "s", "", "Site [required]")
  addCmd.Flags().StringVarP(&username, "username", "u", "", "Username [required]")
  addCmd.Flags().StringVarP(&password, "password", "p", "", "Password [required]")

  // Mark required flags
  addCmd.MarkFlagRequired("site")
  addCmd.MarkFlagRequired("username")
  addCmd.MarkFlagRequired("password")
}
