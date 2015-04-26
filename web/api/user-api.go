package api

import (
	"net/http"
)

// infoHandler returns 200 and info about the request
func UserApiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		userApiGet(w, r)
	} else if r.Method == "PUT" {
		userApiPut(w, r)
	} else if r.Method == "POST" {
		userApiPost(w, r)
	} else if r.Method == "DELETE" {
		userApiDelete(w, r)
	}
}

func userApiGet(w http.ResponseWriter, r *http.Request) {
}
func userApiPut(w http.ResponseWriter, r *http.Request) {
}
func userApiPost(w http.ResponseWriter, r *http.Request) {
}
func userApiDelete(w http.ResponseWriter, r *http.Request) {
}
