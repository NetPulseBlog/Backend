package password

import (
	"encoding/base64"
	"math/rand"
)

const DefaultSaltSize = 16

func GenerateRandomSalt(saltSize int) (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(salt), nil
}
