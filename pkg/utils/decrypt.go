package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"log"
	"os"
)

func DbDecrypt() error {

	outFile := "bin/passwd.db.decrypted"

	cred, err := GetAllCredentials()
	if err != nil {
		return err
	}

	err = decryptFile(cred.Key, cred.EncryptedDB, outFile)
	if err != nil {
		return err
	}

	log.Println("[INFO] Database decrypted using key: ", cred.Key)
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
