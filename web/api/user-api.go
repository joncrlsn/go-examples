package api

import (
	"encoding/json"
	"net/http"
)

// Map method to function
//var userApis map[string](function(w http.ResponseWriter, r *http.Request) error)

// UserHandler handles user requests
func UserHandler(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "text/json; charset=utf-8")

	switch r.Method {
	case "GET":
		return userApiGet(w, r)
	case "PUT":
		return userApiPut(w, r)
	default:
		// TODO: Return an error instead!
		w.WriteHeader(400)
	}

	return nil
}

func userApiGet(w http.ResponseWriter, r *http.Request) error {
	buf, _ := ioutil.ReadAll(r.Body())
	json.Unmarshal(buf, &ryanne)

	res1D := &Response1{
		Page:   101,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	w.Write(res1B)
	return nil
}

func userApiPut(w http.ResponseWriter, r *http.Request) error {
	http.Error(w, "Method not allowed", 405)
	return nil
}

func userApiPost(w http.ResponseWriter, r *http.Request) error {
	http.Error(w, "Method not allowed", 405)
	return nil
}

func userApiDelete(w http.ResponseWriter, r *http.Request) error {
	http.Error(w, "Method not allowed", 405)
	return nil
}

//
//
//
type Response1 struct {
	Page   int
	Fruits []string
}
