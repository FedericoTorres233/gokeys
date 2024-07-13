package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/federicotorres233/gokeys/internal/manager"
	"github.com/federicotorres233/gokeys/internal/utils"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new password",
	Run: func(cmd *cobra.Command, args []string) {
		// Get default configs
		def := utils.LoadConfigs()

		// Get password from password manager
		enc_db := filepath.Join(utils.GetBaseDir(), def["encrypted_db"])
		tmp_db := def["temporary_db"]
		pm := manager.NewPasswordManager(tmp_db, enc_db)

		// Add password to db
		err := pm.AddPassword(&record)
		if err != nil {
			utils.LogError(err)
			return
		}

		utils.LogInfo("Added new record to database")

		if record.Username == "" {
			fmt.Printf("Password for website %s added successfully!\n", record.Website)
			return
		}
		fmt.Printf("Password for %s at %s added successfully!\n", record.Username, record.Website)
	},
}

func init() {
	// Add command
	rootCmd.AddCommand(addCmd)

	// Set up flags
	addCmd.Flags().StringVarP(&record.Website, "website", "w", "", "add website [required]")
	// addCmd.Flags().StringVarP(&record.Password, "password", "p", "", "Password [required]")
	addCmd.Flags().StringVarP(&record.Username, "username", "u", "", "add username [required]")
	addCmd.Flags().StringVarP(&record.Notes, "notes", "o", "", "add notes")
	addCmd.Flags().StringVarP(&record.Notes, "email", "e", "", "add email")
	addCmd.Flags().StringVarP(&record.Tag, "tag", "t", "", "add a tag")
	addCmd.Flags().StringVar(&record.Url, "url", "", "add url")
	addCmd.Flags().BoolVarP(&record.Status, "status", "s", false, "set status")
	addCmd.Flags().BoolVar(&record.Favorite, "favorite", false, "set as favorite")

	// Mark required flags
	addCmd.MarkFlagRequired("website")
	addCmd.MarkFlagRequired("username")
}
