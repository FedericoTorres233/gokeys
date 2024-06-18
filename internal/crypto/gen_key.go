package crypto

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"os"
	"path/filepath"

	"github.com/federicotorres233/gokeys/pkg/utils"
	"golang.org/x/crypto/pbkdf2"
)

func GenerateKey(master string) ([]byte, error) {
	salt := make([]byte, 16) // Recommended salt size for pbkdf2
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}

	// Create the directory if it doesn't exist
	err = os.MkdirAll(filepath.Join(utils.GetBaseDir(), "tmp/keys"), 0o744)
	if err != nil {
		return nil, err
	}

	// Generate key
	key := pbkdf2.Key([]byte(master), salt, 4096, 32, sha256.New)

	// Save salt
	salt_encoded := base64.StdEncoding.EncodeToString(salt)
	err = os.WriteFile(filepath.Join(utils.GetBaseDir(), "tmp/keys", "salt"), []byte(salt_encoded), 0o644)
	if err != nil {
		return nil, err
	}

	// Save key
	key_encoded := base64.StdEncoding.EncodeToString(key)
	err = os.WriteFile(filepath.Join(utils.GetBaseDir(), "tmp/keys", "key"), []byte(key_encoded), 0o644)
	if err != nil {
		return nil, err
	}

	return key, nil
}
