package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"os"

	"github.com/federicotorres233/gokeys/internal/utils"
)

func DbDecrypt(tmpdb string) error {
	outDir := tmpdb

	cred, err := GetAllCredentials()
	if err != nil {
		return err
	}

	err = decryptFile(cred.Key, cred.EncryptedDB, outDir)
	if err != nil {
		return err
	}

	utils.LogInfo("Database decrypted and stored at" + tmpdb)
	return nil
}

func decryptFile(key []byte, encryptedDB []byte, outFile string) error {
	nonceSize := 12 // This is the GCM nonce size
	nonce, file := encryptedDB[:nonceSize], encryptedDB[nonceSize:]

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	decrypted, err := gcm.Open(nil, nonce, file, nil)
	if err != nil {
		return err
	}

	err = os.WriteFile(outFile, decrypted, 0o644)
	return err
}
