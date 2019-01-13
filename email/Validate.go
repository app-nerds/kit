package email

import "net/mail"

/*
IsValidEmailAddress returns true/false if a provided email address is valid
*/
func IsValidEmailAddress(email string) bool {
	_, err := mail.ParseAddress(email)

	if err != nil {
		return false
	}

	return true
}
