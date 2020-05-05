/*
 * Copyright (c) 2020. App Nerds LLC. All rights reserved
 */

package identity

type JWTResponse struct {
	Token    string `json:"token"`
	UserID   string `json:"userID"`
	UserName string `json:"userName"`
}
