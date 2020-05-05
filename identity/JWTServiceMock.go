/*
 * Copyright (c) 2020. App Nerds LLC. All rights reserved
 */

package identity

import (
	"github.com/dgrijalva/jwt-go"
)

type JWTServiceMock struct {
	CreateTokenFunc                func(createRequest *CreateTokenRequest) (string, error)
	GetAdditionalDataFromTokenFunc func(token *jwt.Token) map[string]interface{}
	GetUserFromTokenFunc           func(token *jwt.Token) (string, string)
	ParseTokenFunc                 func(tokenFromHeader string) (*jwt.Token, error)
	IsTokenValidFunc               func(token *jwt.Token) error
}

func (m *JWTServiceMock) CreateToken(createRequest *CreateTokenRequest) (string, error) {
	return m.CreateTokenFunc(createRequest)
}

func (m *JWTServiceMock) GetAdditionalDataFromToken(token *jwt.Token) map[string]interface{} {
	return m.GetAdditionalDataFromTokenFunc(token)
}

func (m *JWTServiceMock) GetUserFromToken(token *jwt.Token) (string, string) {
	return m.GetUserFromTokenFunc(token)
}

func (m *JWTServiceMock) ParseToken(tokenFromHeader string) (*jwt.Token, error) {
	return m.ParseTokenFunc(tokenFromHeader)
}

func (m *JWTServiceMock) IsTokenValid(token *jwt.Token) error {
	return m.IsTokenValidFunc(token)
}
