package main

import "encoding/json"

// Success response success message
type Success struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Payload interface{} `json:"payload"`
}

func (s Success) toBytes() ([]byte, error) {
	j, e := json.Marshal(s)
	return j, e
}
