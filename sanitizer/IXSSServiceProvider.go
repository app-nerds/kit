/*
 * Copyright (c) 2021. App Nerds LLC. All rights reserved
 */
package sanitizer

/*
IXSSServiceProvider is an interface for providing cross-site scripting
and sanitization services.
*/
type IXSSServiceProvider interface {
	SanitizeString(input string) string
}
