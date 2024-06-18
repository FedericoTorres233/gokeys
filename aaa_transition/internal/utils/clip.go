package utils

import (
	"github.com/atotto/clipboard"
)

func SetClip(text string) error {
	// Copy text to the clipboard
	err := clipboard.WriteAll(text)
	if err != nil {
		return err
	}

	return nil
}
