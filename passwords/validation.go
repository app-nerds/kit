/*
 * Copyright (c) 2021. App Nerds LLC. All rights reserved
 */
package passwords

import (
	"golang.org/x/crypto/bcrypt"
)

/*
IsPasswordValid takes a hashed password and a plaintext version and returns
*/
func IsPasswordValid(hashedPassword, plaintextPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plaintextPassword))

	if err == nil {
		return true
	}

	return false
}
