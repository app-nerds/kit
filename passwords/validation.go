package passwords

import (
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

/*
IsPasswordValid takes a hashed password and a plaintext version and returns
*/
func IsPasswordValid(hashedPassword, plaintextPassword string) bool {
	passwordBytes, _ := hex.DecodeString(hashedPassword)
	err := bcrypt.CompareHashAndPassword(passwordBytes, []byte(plaintextPassword))

	if err == nil {
		return true
	}

	return false
}
