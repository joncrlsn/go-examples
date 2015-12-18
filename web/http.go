//
// Sets up the URL routing and starts the web server
//
// Copyright (c) 2015 Jon Carlson.  All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.
//
package main

import (
	"fmt"
	"github.com/joncrlsn/go-examples/web/api"
	"github.com/joncrlsn/go-examples/web/data"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// Taken from http://blog.golang.org/error-handling-and-go
// errorHandler adds a ServeHttp method to every errorHandler function
type errorHandler func(http.ResponseWriter, *http.Request) error

// Adds a ServeHttp method to every errorHandler function
func (fn errorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := fn(w, r); err != nil {
		log.Println("Error handling request", err)
		http.Error(w, "Internal Server Error.  Check logs for actual error", 500)
	}
}

// startWebServer does as the name implies.  If a certFile and keyFile are given then
// the HTTPS port is also used and HTTP is redirected to HTTPS.
func startWebServer(httpPort int, httpsPort int, certFile string, keyFile string) {

	// Any static files in the www dir will be served as-is
	http.Handle("/", http.FileServer(http.Dir("./www")))
	fmt.Println("  /login.html is self-explanatory\n    (try https://localhost:8080)")

	// REST/HTTP API handlers
	http.Handle("/api/user", errorHandler(_authBasic(api.UserHandler)))
	fmt.Println("  /api/user returns JSON about users in the system. requires authentication\n    (use joe@example.com/supersecret for credentials)")

	// Setup business logic handlers
	http.Handle("/info", errorHandler(_viewInfo))
	http.Handle("/auth", errorHandler(_authBasic(_viewInfo)))
	fmt.Println("  /info returns information about the request\n    (try https://localhost:8443/info?x=123&y=456 )")
	fmt.Println("  /auth same as above, but requires authentication\n    (use joe@example.com/supersecret for credentials)")

	//
	// Start listening
	//
	log.Println("=== Starting web server ===")

	// If we were given a certificate file, listen for HTTP but redirect to HTTPS
	if len(certFile) > 0 && httpsPort > 0 {
		// Handle everything on HTTPS
		go func() {
			log.Printf("Redirecting HTTP to HTTPS (port %v redirects to %v)", httpPort, httpsPort)
			log.Fatal(http.ListenAndServe(":"+strconv.Itoa(httpPort), http.HandlerFunc(_redirectToHttps(httpPort, httpsPort))))
		}()

		log.Println("Listening for HTTPS on port", httpsPort)
		log.Fatal(http.ListenAndServeTLS(":"+strconv.Itoa(httpsPort), certFile, keyFile, nil))
		return
	}

	// In this case, we'll go insecure (no HTTPS certificate was given)
	log.Println("Listening for HTTP on port", httpPort)
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(httpPort), nil))
}

// _authBasic wraps a request handler with another one that requires BASIC HTTP authentication
func _authBasic(handler func(http.ResponseWriter, *http.Request) error) func(w http.ResponseWriter, req *http.Request) error {
	return func(w http.ResponseWriter, req *http.Request) error {
		//
		// Ensure request has an "Authorization" header (needed for "Basic" authentication)
		// (That header is misnamed. It should be Authentication, but I don't make the specs)
		// (Authorization makes sure the authenticated user is permitted to do the action)
		//
		username, password, ok := req.BasicAuth()
		if !ok {
			// Request the "Authorization" header
			w.Header().Set("WWW-Authenticate", `Basic realm="go-example-web"`)
			http.Error(w, "Access Denied", http.StatusUnauthorized)
			return nil
		}

		//
		// Ensure the given credentials match one of the users in our database
		//
		user, err := data.Authenticate(db, username, password)
		if err != nil {
			log.Println("Error authenticating user", username, err)
			return err
		}
		if user.UserId == 0 {
			// User authentication failed
			w.Header().Set("WWW-Authenticate", `Basic realm="go-example-web"`)
			http.Error(w, "Access Denied", http.StatusUnauthorized)
			return nil
		}

		//
		// The credentials match, so run the wrapped handler
		//
		return handler(w, req)
	}
}

// _redirectToHttps returns a function that redirects anything on the http port over to the https port.
// We have to wrap the function in a function so we can dynamically provide the http and https ports.
func _redirectToHttps(httpPort int, httpsPort int) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		newHost := strings.Replace(req.Host, strconv.Itoa(httpPort), strconv.Itoa(httpsPort), 1)
		newUrl := fmt.Sprintf("https://%s/%s", newHost, req.RequestURI)
		http.Redirect(w, req, newUrl, http.StatusMovedPermanently)
	}
}

// _viewInfo is an example handler that returns info about the request in HTML format
func _viewInfo(w http.ResponseWriter, r *http.Request) error {
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
		<br>
		r.Header.Get("Authorization") = %s<br>
		</html>`, r.Method, r.RequestURI, r.Host, r.Header, r.RemoteAddr, r.MultipartForm, r.PostForm, r.Form, r.URL, r.URL.Host, r.URL.Path, r.URL.RawQuery, r.URL.Scheme, r.URL.Opaque, r.URL.User, r.URL.Fragment, r.Header.Get("Authorization"))
	}
	return nil
}
