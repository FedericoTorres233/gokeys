package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"log"
	"os"

	"golang.org/x/crypto/pbkdf2"
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

	err = os.WriteFile(outFile, []byte(encoded), 0644)
	return err
}

func deriveKey(password, salt []byte) []byte {
	return pbkdf2.Key(password, salt, 4096, 32, sha256.New)
}

func DbEncrypt(master string, dbdir string) {

	password := []byte(master)
	salt := make([]byte, 16) // Recommended salt size for pbkdf2
	_, err := rand.Read(salt)
	if err != nil {
		log.Println(err)
	}

	key := deriveKey(password, salt)
	if err != nil {
		log.Println(err)
	}

	inFile := dbdir
	outFile := dbdir + ".encrypted" //(This file won't be usable by the database)

	err = encryptFile(key, inFile, outFile)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("[INFO] Database encrypted using key: ", key)
}
