package manager

import "errors"

// GetPassword retrieves a password for a specific site and username
func (pm *PasswordManager) GetPassword(site, username string) (string, error) {
	if users, found := passwordStore[site]; found {
		if password, found := users[username]; found {
			return password, nil
		}
	}
	return "", errors.New("Password not found")
}
