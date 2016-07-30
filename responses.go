package main

import (
	"encoding/json"
	"net/http"
)

// Message response message
type Message struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

// ToBytes convert Message struct to byte array
func (s Message) ToBytes() ([]byte, error) {
	j, e := json.Marshal(s)
	return j, e
}

// Send convert message to byte array and send the response
func (s Message) Send(w http.ResponseWriter) {
	resBody, err := s.ToBytes()
	if err != nil {
		ErrorResponse(http.StatusBadRequest, "Failed to send!", w)
	} else {
		Response(http.StatusOK, resBody, w)
	}
}

// ErrorResponse send error message and code
func ErrorResponse(errCode int, errMessage string, w http.ResponseWriter) {
	body, _ := Message{
		Success: false,
		Message: errMessage,
		Payload: struct{}{},
	}.ToBytes()
	w.WriteHeader(errCode)
	w.Write(body)
}

// Response set the body and status code
func Response(statusCode int, body []byte, w http.ResponseWriter) {
	w.WriteHeader(statusCode)
	w.Write(body)
}
