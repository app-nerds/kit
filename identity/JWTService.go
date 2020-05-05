/*
 * Copyright (c) 2020. App Nerds LLC. All rights reserved
 */

package identity

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"io"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"golang.org/x/crypto/pbkdf2"
)

type IJWTService interface {
	CreateToken(createRequest *CreateTokenRequest) (string, error)
	GetAdditionalDataFromToken(token *jwt.Token) map[string]interface{}
	GetUserFromToken(token *jwt.Token) (string, string)
	ParseToken(tokenFromHeader string) (*jwt.Token, error)
	IsTokenValid(token *jwt.Token) error
}

/*
JWTService provides methods for working with
JWTs in MailSlurper
*/
type JWTService struct {
	authSalt         string
	authSecret       string
	issuer           string
	timeoutInMinutes int
}

/*
CreateToken creates a new JWT token and encrypts it
*/
func (s *JWTService) CreateToken(createRequest *CreateTokenRequest) (string, error) {
	var err error
	var signedToken string
	var encryptedBase64Token string

	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(s.timeoutInMinutes)).Unix(),
			Issuer:    s.issuer,
		},
		UserID:   createRequest.UserID,
		UserName: createRequest.UserName,
	}

	if createRequest.AdditionalData != nil {
		claims.AdditionalData = createRequest.AdditionalData
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	if signedToken, err = token.SignedString([]byte(s.authSecret)); err != nil {
		return "", errors.Wrapf(err, "Error signing JWT token")
	}

	if encryptedBase64Token, err = s.encryptToken(signedToken); err != nil {
		return "", errors.Wrapf(err, "Error encrypting and encoding token")
	}

	return encryptedBase64Token, nil
}

/*
DecryptToken takes a Base64 encoded token which has been encrypted
using AES-256 encryption. This returns the unencoded, unencrypted
token
*/
func (s *JWTService) decryptToken(token string) (string, error) {
	var err error
	var aesBlock cipher.Block
	var unencodedToken []byte
	var gcm cipher.AEAD
	var nonce []byte
	var resultBytes []byte

	key := s.generateAESKey()

	if unencodedToken, err = base64.RawStdEncoding.DecodeString(token); err != nil {
		return "", errors.Wrapf(err, "Unable to base64 decode JWT token")
	}

	if aesBlock, err = aes.NewCipher(key); err != nil {
		return "", errors.Wrapf(err, "Unable to create AES cipher block")
	}

	if gcm, err = cipher.NewGCM(aesBlock); err != nil {
		return "", errors.Wrapf(err, "Problem creating GCM")
	}

	nonceSize := gcm.NonceSize()
	if len(unencodedToken) < nonceSize {
		return "", errors.Wrapf(err, "Ciphertext too short")
	}

	nonce, cipherText := unencodedToken[:nonceSize], unencodedToken[nonceSize:]

	if resultBytes, err = gcm.Open(nil, nonce, cipherText, nil); err != nil {
		return "", errors.Wrapf(err, "Problem decrypting token")
	}

	return string(resultBytes), nil
}

/*
EncryptToken takes a token string, encrypts it using AES-256,
then encodes it in Base64.
*/
func (s *JWTService) encryptToken(token string) (string, error) {
	var err error
	var aesBlock cipher.Block
	var gcm cipher.AEAD
	var nonce []byte
	var encryptedResult []byte

	key := s.generateAESKey()

	if aesBlock, err = aes.NewCipher(key); err != nil {
		return "", errors.Wrapf(err, "Unable to create AES cipher block")
	}

	if gcm, err = cipher.NewGCM(aesBlock); err != nil {
		return "", errors.Wrapf(err, "Problem creating GCM")
	}

	nonce = make([]byte, gcm.NonceSize())
	io.ReadFull(rand.Reader, nonce)

	encryptedResult = gcm.Seal(nonce, nonce, []byte(token), nil)
	encodedResult := base64.RawStdEncoding.EncodeToString(encryptedResult)

	return encodedResult, nil
}

/*
GetAdditionalDataFromToken retrieves the additional data from the claims object
*/
func (s *JWTService) GetAdditionalDataFromToken(token *jwt.Token) map[string]interface{} {
	var claims *Claims

	claims, _ = token.Claims.(*Claims)
	return claims.AdditionalData
}

/*
GetUserFromToken retrieves the user ID and name from the claims in
a JWT token
*/
func (s *JWTService) GetUserFromToken(token *jwt.Token) (string, string) {
	var claims *Claims

	claims, _ = token.Claims.(*Claims)
	return claims.UserID, claims.UserName
}

/*
NewJWTService creates a new instance of the JWTService struct
*/
func NewJWTService(config *JWTServiceConfig) *JWTService {
	return &JWTService{
		authSalt:         config.AuthSalt,
		authSecret:       config.AuthSecret,
		issuer:           config.Issuer,
		timeoutInMinutes: config.TimeoutInMinutes,
	}
}

/*
ParseToken decrypts the provided token and returns a JWT token object
*/
func (s *JWTService) ParseToken(tokenFromHeader string) (*jwt.Token, error) {
	var result *jwt.Token
	var decryptedToken string
	var err error

	/*
	 * Decrypt token first
	 */
	if decryptedToken, err = s.decryptToken(tokenFromHeader); err != nil {
		return result, errors.Wrapf(err, "Problem decrypting JWT token in Parse")
	}

	if result, err = jwt.ParseWithClaims(decryptedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		var ok bool

		if _, ok = token.Method.(*jwt.SigningMethodHMAC); !ok {
			return result, ErrInvalidToken
		}

		return []byte(s.authSecret), nil
	}); err != nil {
		return result, errors.Wrapf(err, "Problem parsing JWT token")
	}

	if err = s.IsTokenValid(result); err != nil {
		return result, err
	}

	return result, nil
}

/*
IsTokenValid returns an error if there are any issues with the
provided JWT token. Possible issues include:
	* Missing claims
	* Invalid token format
	* Invalid issuer
	* User doesn't have a corresponding entry in the credentials table
*/
func (s *JWTService) IsTokenValid(token *jwt.Token) error {
	var claims *Claims
	var ok bool

	claims, ok = token.Claims.(*Claims)

	if !ok {
		return ErrTokenMissingClaims
	}

	if !token.Valid {
		return ErrInvalidToken
	}

	if claims.Issuer != s.issuer {
		return ErrInvalidIssuer
	}

	return nil
}

func (s *JWTService) generateAESKey() []byte {
	return pbkdf2.Key([]byte(s.authSecret), []byte(s.authSalt), 4096, 32, sha1.New)
}

// func (s *JWTService) pkcs5Padding(content []byte) []byte {
// 	padding := aes.BlockSize - len(content)%aes.BlockSize
// 	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
// 	return append(content, padtext...)
// }
