package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const port string = "8000"

func main() {
	println("Running on localhost:" + port)
	fmt.Println(http.ListenAndServe(":"+port, Handlers()))
}

// Handlers all route handles
func Handlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/status", SetHeaders(Status)).Methods("GET")
	r.HandleFunc("/create", SetHeaders(CreateToken)).Methods("POST")
	r.HandleFunc("/decode", SetHeaders(DecodeToken)).Methods("POST")
	return r
}
