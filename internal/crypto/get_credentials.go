package crypto

import (
	"encoding/base64"
	"os"
	"path/filepath"

	"github.com/federicotorres233/gokeys/internal/types"
	"github.com/federicotorres233/gokeys/internal/utils"
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
	salt, err := os.ReadFile(filepath.Join(utils.GetBaseDir(), "tmp/keys", "salt"))
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
	var b []byte
	var salt []byte
	key := os.Getenv("GOKEYS_KEY")

	err := GetSalt(&salt)
	if err != nil {
		utils.LogError(err)
	}

	if key == "" {
		p, err := utils.ReadMasterPassword()
		if err != nil {
			utils.LogError(err)
		}

		k, err := GenerateKey(p, salt)
		if err != nil {
			utils.LogError(err)
		}

		key = k
	}

	b, err = base64.StdEncoding.DecodeString(key)
	if err != nil {
		return err
	}
	*a = b

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
