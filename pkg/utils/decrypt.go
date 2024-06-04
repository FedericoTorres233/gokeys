package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
	"os"
)

func DbDecrypt(tmpdir string) error {

	outDir := tmpdir

	cred, err := GetAllCredentials()
	if err != nil {
		return err
	}

	err = decryptFile(cred.Key, cred.EncryptedDB, outDir)
	if err != nil {
		return err
	}

	log.Println("[INFO] Database decrypted and stored at ", tmpdir)
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

	err = os.WriteFile(outFile, decrypted, 0644)
	return err
}
