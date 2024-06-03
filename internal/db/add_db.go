package db

// AddPassword adds a password for a specific site and username
func (pm *PasswordManager) AddPassword(site, username, password string) {
	if passwordStore[site] == nil {
		passwordStore[site] = make(map[string]string)
	}
	passwordStore[site][username] = password
}
