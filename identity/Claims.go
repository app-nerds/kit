/*
 * Copyright (c) 2020. App Nerds LLC. All rights reserved
 */

package identity

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

var ErrInvalidToken error = fmt.Errorf("Invalid token")
var ErrTokenMissingClaims error = fmt.Errorf("Token is missing claims")
var ErrInvalidUser error = fmt.Errorf("Invalid user")
var ErrInvalidIssuer error = fmt.Errorf("Invalid issuer")

type Claims struct {
	jwt.StandardClaims
	UserID         string `json:"userID"`
	UserName       string `json:"userName"`
	AdditionalData map[string]interface{}
}
