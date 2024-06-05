package cli

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/federicotorres233/gokeys/internal/manager"
	"github.com/federicotorres233/gokeys/pkg/utils"
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
		err := pm.GetPassword(&record)
		if err != nil {
			log.Println("[ERROR] ", err)
			return
		}

		// Copy to clipboard
		if setClipboard {
			fmt.Println("Copied to clipboard!")
			err := utils.SetClip(record.Password)
			log.Println("[ERROR] Could not set clipboard.", err)
		}

		// Output password
		if record.Username == "" {
			fmt.Printf("Password for website %s: %s\n", record.Website, record.Password)
			return
		}
		fmt.Printf("Password for %s at %s: %s\n", record.Username, record.Website, record.Password)
	},
}

func init() {
	// Add command
	rootCmd.AddCommand(getCmd)

	// Set up flags
	getCmd.Flags().StringVarP(&record.Website, "website", "w", "", "website [required]")
	getCmd.Flags().StringVarP(&record.Username, "username", "u", "", "username")
	getCmd.Flags().BoolVarP(&setClipboard, "clipboard", "c", false, "copy to clipboard")

	// Mark required flags
	getCmd.MarkFlagRequired("website")
}
