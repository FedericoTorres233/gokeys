package cmd

import (
	"fmt"
	"path/filepath"

	"github.com/federicotorres233/gokeys/internal/manager"
	"github.com/federicotorres233/gokeys/internal/utils"
	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getAllCmd = &cobra.Command{
	Use:   "get-all",
	Short: "Get all passwords",
	Run: func(cmd *cobra.Command, args []string) {
		// Get default configs
		def := utils.LoadConfigs()

		// Get password from password manager
		enc_db := filepath.Join(utils.GetBaseDir(), def["encrypted_db"])
		tmp_db := def["temporary_db"]
		pm := manager.NewPasswordManager(tmp_db, enc_db)
		
		records, err := pm.GetAllPasswords()
		if err != nil {
			utils.LogError(err)
			return
		}
		if len(records) == 0 {
			fmt.Println("No passwords found!")
			return
		}

		// Format passwords
		fmt.Println("")
		for _, r := range records {
			fmt.Println("<>--------------<>")
			fmt.Println("Username:", r.Username)
			fmt.Println("Password:", r.Password)
			fmt.Println("Website:", r.Website)
			fmt.Println("Status:", r.Status)
			//fmt.Println("Url:", r.Url)
			//fmt.Println("Tag:", r.Tag)
			//fmt.Println("IsFavorite:", r.Favorite)
		}
		fmt.Println("<>--------------<>")

		if err != nil {
			utils.LogError(err)
			return
		}

		// Copy to clipboard
		if setClipboard {
			fmt.Println("Copied to clipboard!")
			err := utils.SetClip(record.Password)
			utils.LogError(err)
		}

	},
}

func init() {
	// Add command
	rootCmd.AddCommand(getAllCmd)

	// Set up flags
	getAllCmd.Flags().BoolVarP(&setClipboard, "clipboard", "c", false, "copy to clipboard")

}
