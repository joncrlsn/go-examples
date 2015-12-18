package api

import (
	"encoding/json"
	"fmt"
	//	"io/ioutil"
	"net/http"
)

func TestApiPackage(w http.ResponseWriter, r *http.Request) {
	var ryanne map[string]string
	switch r.Method {
	case "GET":
		buf, _ := json.Marshal(&ryanne)
		w.Write(buf)
	case "PUT":
		//		buf, _ := ioutil.ReadAll(r.Body)
		//		json.Unmarshal(buf, &ryanne)
	default:
		w.WriteHeader(400)
	}
	fmt.Println("api package")
}
