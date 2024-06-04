package manager

// PasswordManager stores the passwords
type PasswordManager struct{}

// NewPasswordManager creates a new PasswordManager
func NewPasswordManager() *PasswordManager {
 pm := PasswordManager{}
	return &pm
}
