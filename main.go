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
	println(http.ListenAndServe(":"+port, r))
}
