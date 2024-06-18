package types

type Record struct {
	Password string
	Website  string
	Username string
	Notes    string
	Tag      string
	Url      string
	Favorite bool
	Status   bool
}

type Credentials struct {
	EncryptedDB []byte
	Key         []byte
	Salt        []byte
}
