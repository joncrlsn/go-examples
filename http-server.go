package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// randomHandler responds in a random amount of time
func randomHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		duration := time.Duration(rand.Float64()*30) * time.Second
		time.Sleep(duration)
		fmt.Fprintf(w, "<html>Responding in %s</html>", duration)
	}
}

// errorHandler responds with a 500 Internal Server Error
func errorHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Error(w, "<html>500 Internal Server Error</html>", http.StatusInternalServerError)
	}
}

// notFoundHandler responds with 404 Not Found
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Error(w, "<html>404 Not Found</html>", http.StatusNotFound)
	}
}

// infoHandler returns 200 and info about the request
func infoHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		fmt.Fprintf(w, `<html>
		r.Method = %s <br>
		r.RequestURI = %s <br>
		r.Host = %s <br>
		r.Header = %v <br>
		r.RemoteAddr = %s <br>
		r.MultipartForm = %v <br>
		r.PostForm = %v <br>
		r.Form = %v <br>
		<br>
		r.URL = %s <br>
		r.URL.Host = %s <br>
		r.URL.Path = %s <br>
		r.URL.RawQuery = %s <br>
		r.URL.Scheme = %s <br>
		r.URL.Opaque = %s <br>
		r.URL.User = %s <br>
		r.URL.Fragment = %s <br>
		</html>`, r.Method, r.RequestURI, r.Host, r.Header, r.RemoteAddr, r.MultipartForm, r.PostForm, r.Form, r.URL, r.URL.Host, r.URL.Path, r.URL.RawQuery, r.URL.Scheme, r.URL.Opaque, r.URL.User, r.URL.Fragment)
	}
}

// Redirects port 8080 to port 8443 and changes the protocol to https
func redirectToHttps(w http.ResponseWriter, req *http.Request) {
	newHost := strings.Replace(req.Host, "8080", "8443", 1)
	newUrl := fmt.Sprintf("https://%s/%s", newHost, req.RequestURI)
	http.Redirect(w, req, newUrl, http.StatusMovedPermanently)
}

func main() {
	http.HandleFunc("/random", randomHandler)
	http.HandleFunc("/error", errorHandler)
	http.HandleFunc("/notfound", notFoundHandler)
	http.HandleFunc("/info", infoHandler)

	fmt.Println("  /random returns in a random amount of time")
	fmt.Println("  /error returns an internal-server-error code")
	fmt.Println("  /notfound returns a not-found error code")
	fmt.Println("  /info returns information about the request\n    (try https://localhost:8443/info?x=123&y=456 )")

	var err error

	// Listen for HTTP but redirect to HTTPS
	go func() {
		fmt.Println("Listening for HTTP on port 8080")
		//err = http.ListenAndServe(":8080", nil)
		err = http.ListenAndServe(":8080", http.HandlerFunc(redirectToHttps))
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Handle everything on HTTPS
	fmt.Println("Listening for HTTPS on port 8443")
	err = http.ListenAndServeTLS(":8443", "http-cert.pem", "http-key.pem", nil)
	if err != nil {
		log.Fatal(err)
	}
}
