/*
 * Copyright (c) 2021. App Nerds LLC. All rights reserved
 */

package identity

/*
JWTServiceConfig is a configuration object for initializing the
JWTService struct
*/
type JWTServiceConfig struct {
	AuthSalt         string
	AuthSecret       string
	Issuer           string
	TimeoutInMinutes int
}
