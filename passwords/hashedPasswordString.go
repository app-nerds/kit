/*
 * Copyright (c) 2021. App Nerds LLC. All rights reserved
 */
package passwords

/*
HashedPasswordString is a type of string that represents a hashed
password. This is often useful when dealing with passwords that must
be stored in a database.
*/
type HashedPasswordString string

func (hp HashedPasswordString) Hash() HashedPasswordString {
	result, _ := HashPassword(string(hp))
	return HashedPasswordString(result)
}

func (hp HashedPasswordString) IsEmpty() bool {
	return len(hp) == 0
}

func (hp HashedPasswordString) IsSameAsPlaintextPassword(plaintextPassword string) bool {
	return IsPasswordValid(string(hp), plaintextPassword)
}
