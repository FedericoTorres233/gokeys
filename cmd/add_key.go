package cmd

import (
	"fmt"

	"github.com/federicotorres233/gokeys/internal/manager"
	"github.com/federicotorres233/gokeys/types"
	"github.com/spf13/cobra"
)

var record types.Record

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new password",
	Run: func(cmd *cobra.Command, args []string) {
		pm := manager.NewPasswordManager()
		pm.AddPassword(record)
		fmt.Printf("Password for %s at %s added successfully!\n", record.Username, record.Website)
	},
}

func init() {
	// Add command
	rootCmd.AddCommand(addCmd)

	// Set up flags
	addCmd.Flags().StringVarP(&record.Website, "website", "w", "", "Site [required]")
	addCmd.Flags().StringVarP(&record.Username, "username", "u", "", "Username [required]")
	addCmd.Flags().StringVarP(&record.Password, "password", "p", "", "Password [required]")

	// Mark required flags
	addCmd.MarkFlagRequired("site")
	addCmd.MarkFlagRequired("username")
	addCmd.MarkFlagRequired("password")
}
