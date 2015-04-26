//
// Sets up the URL routing and runs the web server
//
package main

import (
	"fmt"
	"github.com/joncrlsn/go-examples/web/data"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// startWebServer does as the name implies.  If a certFile and keyFile are given then
// the HTTPS port is also used and HTTP is redirected to HTTPS.
func startWebServer(httpPort int, httpsPort int, certFile string, keyFile string) {

	//
	// Setup business logic handlers
	//
	http.HandleFunc("/info", _infoHandler)
	http.HandleFunc("/secure", _authenticateBasic(_infoHandler))
	fmt.Println("  /info returns information about the request\n    (try https://localhost:8443/info?x=123&y=456 )")

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

// _authenticateBasic wraps a handler with another one that requires BASIC HTTP authentication
func _authenticateBasic(handler func(http.ResponseWriter, *http.Request)) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
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
			return
		}

		//
		// Ensure the given credentials match one of the users in our database
		//
		user, err := data.Authenticate(db, username, password)
		if err != nil {
			log.Println("Error authenticating user", username, err)
		}
		if user.UserId == 0 {
			// User authentication failed
			w.Header().Set("WWW-Authenticate", `Basic realm="go-example-web"`)
			http.Error(w, "Access Denied", http.StatusUnauthorized)
			return
		}

		//
		// The credentials match, so run the wrapped handler
		//
		handler(w, req)
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

// _infoHandler is an example handler that returns info about the request in HTML format
func _infoHandler(w http.ResponseWriter, r *http.Request) {
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
}
