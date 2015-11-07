//
// Runs an HTTP static file server on the directory that this is executed from
//
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	var port string
	if len(os.Args) > 1 {
		port = os.Args[1]
	} else {
		port = "8080"
	}

	fmt.Printf("Serving files in the current directory on port %s \n", port)
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("err=", err)
		os.Exit(1)
	}
	http.Handle("/", http.FileServer(http.Dir(dir)))
	log.Fatal(http.ListenAndServe(port, nil))

	// Another way: to do it
	// log.Fatal(http.ListenAndServe(":" + port, http.FileServer(http.Dir(dir))))
}
