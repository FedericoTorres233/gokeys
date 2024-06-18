package utils

import (
	"log"
	"os"
	"path/filepath"
)

type Logger struct {
	dir      string
	filepath string
}

func NewLogger(d string, f string) *Logger {

	// Get the Base dir
	home := GetBaseDir()
	dir := filepath.Join(home, d)
	filepath := filepath.Join(dir, f)

	// Create the directory if it doesn't exist
	err := os.MkdirAll(d, 0o744)
	if err != nil {
		log.Fatal(err)
	}

	// Open the log file in append mode, create if it doesn't exist
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		log.Fatal(err)
	}

	l := Logger{
		dir:      d,
		filepath: f,
	}

	log.SetOutput(file)

	return &l

}

func LogError(err error) {
	log.Println("[ERROR]", err)
}

func LogInfo(text string) {
	log.Println("[INFO]", text)
}

func LogWarning(text string) {
	log.Println("[WARN]")
}
