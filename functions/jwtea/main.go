package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	apex "github.com/apex/go-apex"
	"github.com/gorilla/mux"
)

const (
	port string = "8000"
)

// func main() {
// 	Initialize()
// 	println("Running on localhost:" + port)
// 	fmt.Println(http.ListenAndServe(":"+port, Handlers()))
// }

func main() {
	Initialize()
	// set up the HTTP routing
	handler := Handlers()

	// register the Lambda event handler
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		req, err := ParseRequest(event)
		if err != nil {
			return FormatError(http.StatusBadRequest, err), nil
		}

		res := httptest.NewRecorder()

		// handle the HTTP request
		handler.ServeHTTP(res, req)

		return FormatResponse(res), nil
	})
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
