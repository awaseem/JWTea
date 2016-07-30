package main

import "net/http"

// Status handler that returns an OK
func Status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	message := Success{
		Success: true,
		Message: "OK",
		Payload: struct{}{},
	}
	res, _ := message.toBytes()
	w.Write(res)
}
