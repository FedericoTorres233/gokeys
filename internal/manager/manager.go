package manager

// PasswordManager stores the passwords
type PasswordManager struct {
	encdb string
	tmpdb string
}

// NewPasswordManager creates a new PasswordManager
func NewPasswordManager(tmpdb string, encdb string) *PasswordManager {
	pm := PasswordManager{
		tmpdb: tmpdb,
		encdb: encdb,
	}
	return &pm
}
