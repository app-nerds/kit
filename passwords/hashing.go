package passwords

import (
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

/*
HashPassword takes a password as a string and returns a hex encoded bcrypted
representation.
*/
func HashPassword(password string) (string, error) {
	result := ""
	passwordBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return result, err
	}

	result = hex.EncodeToString(passwordBytes)
	return result, nil
}
