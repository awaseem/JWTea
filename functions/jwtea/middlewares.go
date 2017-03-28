package main

import "net/http"

// SetHeaders sets the headers for handlers via a middleware
func SetHeaders(h func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	// this function returns another function which will exectue the handler
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h(w, r)
	}
}
