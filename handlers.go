package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

// Status handler that returns an OK
func Status(w http.ResponseWriter, r *http.Request) {
	message := Message{
		Success: true,
		Message: "OK",
		Payload: struct{}{},
	}
	resBody, err := message.ToBytes()
	if err != nil {
		ErrorResponse(http.StatusBadRequest, "Failed to find status!", w)
	} else {
		Response(http.StatusOK, resBody, w)
	}
}

// CreateToken that creates a new token
func CreateToken(w http.ResponseWriter, r *http.Request) {
	var postBody interface{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ErrorResponse(http.StatusBadRequest, "Failed to read body!", w)
	}
	// create token based on the post body as claims
	json.Unmarshal(body, &postBody)
	tokenHelper := Token{
		postBody,
		jwt.StandardClaims{
			ExpiresAt: 15000,
		},
	}
	token, err := tokenHelper.Generate()
	if err != nil {
		ErrorResponse(http.StatusBadRequest, "Failed to create token!", w)
	}
	// Generate request message payload
	resMessage := Message{
		Success: true,
		Message: "Created Token!",
		Payload: TokenResponse{
			Token:  token,
			Expiry: tokenHelper.ExpiresAt,
		},
	}
	resBody, err := resMessage.ToBytes()
	if err != nil {
		ErrorResponse(http.StatusBadRequest, "Failed to send token!", w)
	} else {
		Response(http.StatusOK, resBody, w)
	}
}
