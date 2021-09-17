/*
 * Copyright (c) 2021. App Nerds LLC. All rights reserved
 */

package identity

/*
JWTResponse is a generic reponse that can be used to communicate a
new JWT token to a caller.
*/
type JWTResponse struct {
	Token    string `json:"token"`
	UserID   string `json:"userID"`
	UserName string `json:"userName"`
}
