package db

// global variable to simulate persistent storage for now
var passwordStore = make(map[string]map[string]string)

// PasswordManager stores the passwords
type PasswordManager struct{}

// NewPasswordManager creates a new PasswordManager
func NewPasswordManager() *PasswordManager {
	return &PasswordManager{}
}
