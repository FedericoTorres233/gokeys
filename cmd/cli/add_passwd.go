package cli

import (
	"fmt"
	"log"

	"github.com/federicotorres233/gokeys/internal/manager"
	"github.com/federicotorres233/gokeys/pkg/utils"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new password",
	Run: func(cmd *cobra.Command, args []string) {

		// Password manager instance
		pm := manager.NewPasswordManager()
		
		// Read password from stdin
		p, err := utils.ReadPassword()
		if err != nil {
			log.Println("[ERROR] ", err)
			return
		}
		record.Password = p

		// Add password to db
		err = pm.AddPassword(&record)
		if err != nil {
			log.Println("[ERROR] ", err)
			return
		}

		log.Println("[INFO] Added new record to database")

		if record.Username == ""{
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
	addCmd.Flags().StringVarP(&record.Website, "website", "w", "", "Website [required]")
	//addCmd.Flags().StringVarP(&record.Password, "password", "p", "", "Password [required]")
	addCmd.Flags().StringVarP(&record.Username, "username", "u", "", "Username")

	// Mark required flags
	addCmd.MarkFlagRequired("website")
	//addCmd.MarkFlagRequired("password")
}
