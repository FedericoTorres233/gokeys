package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"os"

	"golang.org/x/crypto/pbkdf2"
)

func GenerateKey(master string) ([]byte, error) {
	salt := make([]byte, 16) // Recommended salt size for pbkdf2
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}

	// Generate key
	key := pbkdf2.Key([]byte(master), salt, 4096, 32, sha256.New)

	// Save salt
	salt_encoded := base64.StdEncoding.EncodeToString(salt)
	err = os.WriteFile("bin/salt", []byte(salt_encoded), 0o644)
	if err != nil {
		return nil, err
	}

	// Save key
	key_encoded := base64.StdEncoding.EncodeToString(key)
	err = os.WriteFile("bin/key", []byte(key_encoded), 0o644)
	if err != nil {
		return nil, err
	}

	return key, nil
}
