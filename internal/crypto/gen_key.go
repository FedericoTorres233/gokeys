package crypto

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"os"
	"path/filepath"

	"github.com/federicotorres233/gokeys/internal/utils"
	"golang.org/x/crypto/pbkdf2"
)

func GenerateNewKey(master string) ([]byte, error) {
	salt := make([]byte, 16) // Recommended salt size for pbkdf2
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}

	// Create the directory if it doesn't exist
	err = os.MkdirAll(filepath.Join(utils.GetBaseDir(), "tmp", "keys"), 0o744)
	if err != nil {
		return nil, err
	}

	// Generate key
	key := pbkdf2.Key([]byte(master), salt, 4096, 32, sha256.New)

	// Save salt
	salt_encoded := base64.StdEncoding.EncodeToString(salt)
	err = os.WriteFile(filepath.Join(utils.GetBaseDir(), "tmp", "keys", "salt"), []byte(salt_encoded), 0o644)
	if err != nil {
		return nil, err
	}

	// Save key
	key_encoded := base64.StdEncoding.EncodeToString(key)
	err = os.Setenv("GOKEYS_KEY", string(key_encoded))
	if err != nil {
		utils.LogError(err)
		return nil, err
	}

	return key, nil
}

func GenerateKey(master string, salt []byte) (string, error) {

	// Generate key
	key := pbkdf2.Key([]byte(master), salt, 4096, 32, sha256.New)

	// Save key
	key_encoded := base64.StdEncoding.EncodeToString(key)
	err := os.Setenv("GOKEYS_KEY", key_encoded)
	if err != nil {
		utils.LogError(err)
		return "", err
	}

	return key_encoded, nil
}
