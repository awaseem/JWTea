package main

import (
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

const signingString = "SIGN_STRING"

// TokenRequest token request object for the incoming payload
type TokenRequest struct {
	Token string `json:"token"`
}

// TokenResponse token response object for the payload of a Message
type TokenResponse struct {
	Token  string `json:"token"`
	Expiry int64  `json:"expiry"`
}

// Token custom claims for json tokens
type Token struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// User struct for username creation
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Generate create a token string based on the struct data
func (t *Token) Generate() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, t)
	return token.SignedString([]byte(os.Getenv(signingString)))
}

// Decode parse the token string
func (t *Token) Decode(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &Token{}, func(token *jwt.Token) (interface{}, error) {
		// since we only use the one private key to sign the tokens,
		// we also only use its public counter part to verify
		return []byte(os.Getenv(signingString)), nil
	})
	if err != nil {
		return err
	}
	*t = *token.Claims.(*Token)
	return nil
}
