package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func LoadConfigs() map[string]string {
	// Specify the path to the file you want to load
	filePath := filepath.Join(".", "public", "config")

	// Read the contents of the file
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	// Convert the file content to a string
	fileContent := string(content)

	// Split the content into lines
	lines := strings.Split(fileContent, "\n")

	// Initialize an empty map
	data := make(map[string]string)

	// Parse each line and populate the map
	for _, line := range lines {
		parts := strings.SplitN(line, "=", 2) // Assuming key-value pairs are separated by "="
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			data[key] = value
		}
	}

	return data
}
