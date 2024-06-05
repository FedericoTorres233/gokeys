package utils

import "os"

func IsProduction() bool {
	// Get the environment variable
	mode := os.Getenv("MODE")
	
	// Check if it is set to "production"
	return mode == "production"
}
