package api

import (
	"fmt"
)

func TestApiPackage() {
	switch req.Method {
	case "GET":
		buf, _ := json.Marshal(&ryanne)
		w.Write(buf)
	case "PUT":
		buf, _ := ioutil.ReadAll(req.Body())
		json.Unmarshal(buf, &ryanne)
	default:
		w.WriteHeader(400)
	}
	fmt.Println("api package")
}
