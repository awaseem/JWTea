package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

const port string = "8000"

func main() {
	r := mux.NewRouter()
	println("Running on localhost:" + port)
	r.HandleFunc("/status", SetHeaders(Status)).Methods("GET")
	r.HandleFunc("/create", SetHeaders(CreateToken)).Methods("POST")
	r.HandleFunc("/decode", SetHeaders(DecodeToken)).Methods("POST")
	println(http.ListenAndServe(":"+port, r))
}
