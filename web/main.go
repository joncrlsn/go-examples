package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

// Configurable variables
var (
	dbType         = "sqlite3"
	dbUrl          = "web.db"
	migrateSqlPath = "./migrations"
	httpPort       = 8080
	httpsPort      = 8443
	certFile       = "https-cert.pem"
	keyFile        = "https-key.pem"
	//certFile = ""
	//keyFile  = ""
)

// this variable is accessed in the http.go
var db *sqlx.DB

func main() {
	fmt.Println("==================================\n== Jon's Example Web App\n==================================")

	//
	// Setting up the database
	//
	var err error
	db, err = setupDb(dbType, dbUrl, migrateSqlPath)
	if err != nil {
		log.Println("Error setting up the db", err)
		return
	}

	//
	// Starting Web Server
	//
	startWebServer(httpPort, httpsPort, certFile, keyFile)
}
