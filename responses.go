package main

import (
	"encoding/json"
	"net/http"
)

// Message response success message
type Message struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

func (s Message) ToBytes() ([]byte, error) {
	j, e := json.Marshal(s)
	return j, e
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
