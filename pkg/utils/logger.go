package utils

import (
	"log"
	"os"
	"path/filepath"
)

func SetupLogger(logdir string) {
	// Create the directory if it doesn't exist
	err := os.MkdirAll(logdir, 0o755)
	if err != nil {
		log.Fatal(err)
	}

	// Open the log file in append mode, create if it doesn't exist
	filePath := filepath.Join(logdir, "error.log")
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		log.Fatal(err)
	}

	// Set the output of the default logger to the file
	log.SetOutput(file)
}
