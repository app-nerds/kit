// Copyright 2018 AppNinjas. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

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
	UserID   string `json:"userID"`
	UserName string `json:"userName"`
}
