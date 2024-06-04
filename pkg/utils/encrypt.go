package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"log"
	"os"
)

func encryptFile(key []byte, inFile string, outFile string) error {
	plaintext, err := os.ReadFile(inFile)
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	dst := make([]byte, gcm.NonceSize())
	nonce := make([]byte, gcm.NonceSize())
	if _, err = rand.Read(nonce); err != nil {
		return err
	}

	ciphertext := gcm.Seal(dst[:0], nonce, plaintext, nil)

	combined := append(nonce, ciphertext...)
	encoded := base64.StdEncoding.EncodeToString(combined)

	err = os.WriteFile(outFile, []byte(encoded), 0o644)
	if err != nil {
		return err
	}

	return nil
}

func DbEncrypt(key []byte, encdb string, tmpdb string) error {
	err := encryptFile(key, tmpdb, encdb)
	if err != nil {
		return err
	}

	log.Println("[INFO] Database encrypted using key: ", key)
	return nil
}
