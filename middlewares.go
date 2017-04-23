package main

import (
	"net/http"

	"github.com/bnkamalesh/webgo"
)

//authCheck is a middleware to check if a user is authorized based on a header token
func authCheck(w http.ResponseWriter, r *http.Request) {
	// User is authorized if the request header has the key `Authorization`, and is of length greater than 3
	if len(r.Header.Get("Authorization")) < 3 {
		// Unauthorized
		webgo.R403(w, "Sorry, you're not authorized to access this page")
		return
	}
}

//afterResponse is a middleware which is prints `after response` to stdout
func afterResponse(w http.ResponseWriter, r *http.Request) {
	println("I was executed after writing a response!")
}
