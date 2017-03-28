package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"
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

// CreateUser create new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var incomingBody User
	body, errIO := ioutil.ReadAll(r.Body)
	if errIO != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}
	if errUnmarshal := json.Unmarshal(body, &incomingBody); errUnmarshal != nil {
		http.Error(w, "Failed to prase body into json!", http.StatusBadRequest)
		return
	}
	passwordHash, errGen := bcrypt.GenerateFromPassword([]byte(incomingBody.Password), 10)
	if errGen != nil {
		http.Error(w, "Error failed to generate password", http.StatusInternalServerError)
		return
	}
	errSet := Set(incomingBody.Username, passwordHash)
	if errSet != nil {
		http.Error(w, "Failed to store username and password for user creation!", http.StatusBadRequest)
		return
	}
	resMessage := Message{
		Success: true,
		Message: "Created User!",
	}
	resMessage.Send(w)
}

// Login create token bsaed on username and password
func Login(w http.ResponseWriter, r *http.Request) {
	var incomingBody User
	body, errIO := ioutil.ReadAll(r.Body)
	if errIO != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}
	if errUnmarshal := json.Unmarshal(body, &incomingBody); errUnmarshal != nil {
		http.Error(w, "Failed to prase body into json!", http.StatusBadRequest)
		return
	}
	password, errGet := Get(incomingBody.Username)
	if errGet != nil {
		http.Error(w, "Failed to find user", http.StatusBadRequest)
		return
	}
	errComPass := bcrypt.CompareHashAndPassword(password, []byte(incomingBody.Password))
	if errComPass != nil {
		http.Error(w, "Failed to compare password", http.StatusBadRequest)
		return
	}
	tokenBody := Token{
		incomingBody.Username,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}
	token, errToken := tokenBody.Generate()
	if errToken != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}
	resMessage := Message{
		Success: true,
		Message: "Created Token!",
		Payload: TokenResponse{
			Token:  token,
			Expiry: tokenBody.ExpiresAt,
		},
	}
	resMessage.Send(w)
}

// CheckUser check if token is valid
func CheckUser(w http.ResponseWriter, r *http.Request) {
	// parse post body
	var incomingToken TokenRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ErrorResponse(http.StatusBadRequest, "Failed to read body!", w)
		return
	}
	// parse the token
	if err := json.Unmarshal(body, &incomingToken); err != nil {
		ErrorResponse(http.StatusBadRequest, "Failed to prase body into json!", w)
		return
	}
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
