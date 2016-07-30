package main

import (
	"os"

	jwt "github.com/dgrijalva/jwt-go"
)

const signingString = "SIGN_STRING"

// TokenResponse response object for the payload of a message
type TokenResponse struct {
	Token  string `json:"token"`
	Expiry int64  `json:"expiry"`
}

// Token custom claims for json tokens
type Token struct {
	Body interface{} `json:"body"`
	jwt.StandardClaims
}

// Generate create a token string based on the struct data
func (t Token) Generate() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, t)
	return token.SignedString([]byte(os.Getenv(signingString)))
}
