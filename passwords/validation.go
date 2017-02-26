/*
Copyright 2017 AppNinjas LLC

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
*/
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
