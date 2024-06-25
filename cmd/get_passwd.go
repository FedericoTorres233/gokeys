package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/federicotorres233/gokeys/internal/manager"
	"github.com/federicotorres233/gokeys/internal/utils"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a password",
	Run: func(cmd *cobra.Command, args []string) {
		// Get default configs
		def := utils.LoadConfigs()

		// Get password from password manager
		enc_db := filepath.Join(utils.GetBaseDir(), def["encrypted_db"])
		tmp_db := def["temporary_db"]
		pm := manager.NewPasswordManager(tmp_db, enc_db)

		// Check flags
		if record.Website != "" {
			err := pm.GetPassword(&record)
			if err != nil {
				utils.LogError(err)
				return
			}
			fmt.Printf("Password for website %s: %s\n", record.Website, record.Password)
		} else if record.Username != "" {
			err := pm.GetPassword(&record)
			if err != nil {
				utils.LogError(err)
				return
			}
			fmt.Printf("Password for %s at %s: %s\n", record.Username, record.Website, record.Password)
		}

		// Copy to clipboard
		if setClipboard {
			fmt.Println("Copied to clipboard!")
			err := utils.SetClip(record.Password)
			utils.LogError(err)
		}
		return

	},
}

func init() {
	// Add command
	rootCmd.AddCommand(getCmd)

	// Set up flags
	getCmd.Flags().StringVarP(&record.Website, "website", "w", "", "website")
	getCmd.Flags().StringVarP(&record.Username, "username", "u", "", "username")

	getCmd.Flags().BoolVarP(&setClipboard, "clipboard", "c", false, "copy to clipboard")

}
