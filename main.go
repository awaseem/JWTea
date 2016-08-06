package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const port string = "8000"

func main() {
	Initialize()
	println("Running on localhost:" + port)
	fmt.Println(http.ListenAndServe(":"+port, Handlers()))
}

// Handlers all route handles
func Handlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/status", SetHeaders(Status)).Methods("GET")
	r.HandleFunc("/create", SetHeaders(CreateUser)).Methods("POST")
	r.HandleFunc("/login", SetHeaders(Login)).Methods("POST")
	r.HandleFunc("/check", SetHeaders(CheckUser)).Methods("POST")
	return r
}
