package utils

import "path/filepath"

func GetBaseDir() string {
	// Directory for storing credentials
	var dir string
	if IsProduction() {
		s, _ := GetUserHome()
		dir = filepath.Join(s, ".gokeys")
	} else {
		dir = "."
	}
	return dir
}
