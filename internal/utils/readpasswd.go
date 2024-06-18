package utils

import (
	"fmt"

	"golang.org/x/term"
)

func ReadPassword() (string, error) {
	fmt.Print("Enter the password [does not echo]: ")
	p, err := term.ReadPassword(0)
	if err != nil {
		LogError(err)
		return "", err
	}
	fmt.Print("\n")
	return string(p), nil
}
