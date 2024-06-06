package crypto

import (
	"encoding/base64"
	"os"
	"path/filepath"

	"github.com/federicotorres233/gokeys/internal/types"
	"github.com/federicotorres233/gokeys/pkg/utils"
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
	salt, err := os.ReadFile(filepath.Join(utils.GetBaseDir(), "keys", "salt"))
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
	key, err := os.ReadFile(filepath.Join(utils.GetBaseDir(), "keys", "key"))
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

	// Get default filepaths for databases
	def := utils.LoadConfigs()
	enc_db := filepath.Join(utils.GetBaseDir(), def["encrypted_db"])
	db_encrypted, err := os.ReadFile(enc_db)
	if err != nil {
		return err
	}

	*a, err = base64.StdEncoding.DecodeString(string(db_encrypted))
	if err != nil {
		return err
	}

	return nil
}
