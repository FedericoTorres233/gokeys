package utils

import "os/user"

func GetUserHome() (string, error) {
	// Get the current user
	currentUser, err := user.Current()
	if err != nil {

		return "", err
	}

	// Get the home directory of the current user
	homeDir := currentUser.HomeDir
	return homeDir, nil
}
