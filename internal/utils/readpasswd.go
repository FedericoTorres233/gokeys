package utils

import (
	"fmt"

	"golang.org/x/term"
)

func readPassword() ([]byte, error) {
	return term.ReadPassword(0)
}

func ReadMasterPassword() (string, error) {
	fmt.Print("Enter the {Master Password} [does not echo]: ")
	p, err := readPassword()
	if err != nil {
		LogError(err)
		return "", err
	}
	fmt.Print("\n")
	return string(p), nil
}

func ReadPassword() (string, error) {
	fmt.Print("Enter the password [does not echo]: ")
	p, err := readPassword()
	if err != nil {
		LogError(err)
		return "", err
	}
	fmt.Print("\n")
	return string(p), nil
}
