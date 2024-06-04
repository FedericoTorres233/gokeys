package types

type Credentials struct {
	EncryptedDB []byte
	Key         []byte
	Salt        []byte
}
