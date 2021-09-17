/*
 * Copyright (c) 2021. App Nerds LLC. All rights reserved
 */

package email

/*
Mail represents an email. Who's sending, recipients, subject, and message
*/
type Mail struct {
	Body    string
	From    Person
	Subject string
	To      []Person
}
