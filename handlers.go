package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Status handler that returns an OK
func Status(w http.ResponseWriter, r *http.Request) {
	message := Message{
		Success: true,
		Message: "OK",
		Payload: struct{}{},
	}
	message.Send(w)
}

// CreateToken that creates a new token
func CreateToken(w http.ResponseWriter, r *http.Request) {
	// parse post body
	var incomingBody interface{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ErrorResponse(http.StatusBadRequest, "Failed to read body!", w)
		return
	}
	// create token based on the post body as claims
	json.Unmarshal(body, &incomingBody)
	tokenHelper := Token{
		incomingBody,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}
	token, err := tokenHelper.Generate()
	if err != nil {
		ErrorResponse(http.StatusBadRequest, "Failed to create token!", w)
		return
	}
	// generate request message payload
	resMessage := Message{
		Success: true,
		Message: "Created Token!",
		Payload: TokenResponse{
			Token:  token,
			Expiry: tokenHelper.ExpiresAt,
		},
	}
	resMessage.Send(w)
}

// DecodeToken parses token string and sends it back as a response
func DecodeToken(w http.ResponseWriter, r *http.Request) {
	// parse post body
	var incomingToken TokenRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ErrorResponse(http.StatusBadRequest, "Failed to read body!", w)
	}
	// parse the token
	json.Unmarshal(body, &incomingToken)
	var parsedToken Token
	if err := parsedToken.Decode(incomingToken.Token); err != nil {
		ErrorResponse(http.StatusBadRequest, "Failed to parse token!", w)
		return
	}
	// generate request message payload
	resMessage := Message{
		Success: true,
		Message: "Parsed Token!",
		Payload: parsedToken,
	}
	resMessage.Send(w)
}
