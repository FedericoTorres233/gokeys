package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

)

func SetupLogger() {
	// Get the $HOME
	home, err := GetUserHome()
	if err != nil {
		fmt.Println("Could not set the log directory")
		os.Exit(1)
	}

	// Create the directory if it doesn't exist
	dir := filepath.Join(home, ".cache", "gokeys", "logs")

	err = os.MkdirAll(dir, 0o744)
	if err != nil {
		log.Fatal(err)
	}

	// Open the log file in append mode, create if it doesn't exist
	filePath := filepath.Join(dir, "gokeys.log")
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		log.Fatal(err)
	}

	// Set the output of the default logger to the file
	log.SetOutput(file)
}
