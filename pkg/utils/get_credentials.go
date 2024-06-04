package utils

import (
	"encoding/base64"
	"os"

	"github.com/federicotorres233/gokeys/internal/types"
)

func GetAllCredentials() (types.Credentials, error) {
	var c types.Credentials

	err := GetSalt(&c.Salt)
	if err != nil {
		return c, err
	}

	err = GetKey(&c.Key)
	if err != nil {
		return c, err
	}

	err = GetEncryptedDB(&c.EncryptedDB)
	if err != nil {
		return c, err
	}

	return c, nil
}

func GetSalt(a *[]byte) error {
	salt, err := os.ReadFile("bin/salt")
	if err != nil {
		return err
	}

	*a, err = base64.StdEncoding.DecodeString(string(salt))
	if err != nil {
		return err
	}

	return nil
}

func GetKey(a *[]byte) error {
	key, err := os.ReadFile("bin/key")
	if err != nil {
		return err
	}

	*a, err = base64.StdEncoding.DecodeString(string(key))
	if err != nil {
		return err
	}

	return nil
}

func GetEncryptedDB(a *[]byte) error {
	db_encrypted, err := os.ReadFile("bin/passwd.db.encrypted")
	if err != nil {
		return err
	}

	*a, err = base64.StdEncoding.DecodeString(string(db_encrypted))
	if err != nil {
		return err
	}

	return nil
}
