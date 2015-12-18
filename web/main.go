package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	flag "github.com/ogier/pflag"
	"log"
	"os"
)

// Configurable variables
var (
	version        = "0.1"
	verbose        = false
	dbType         = "sqlite3"
	dbUrl          = "web.db"
	migrateSqlPath = "./migrations"
	httpPort       = 8080
	httpsPort      = 8443
	certFile       = "https-cert.pem"
	keyFile        = "https-key.pem"
	mailHost       = ""
	mailPort       = 25
	mailUsername   = ""
	mailPassword   = ""
	mailFrom       = ""         // an email address
	mailTo         = []string{} // a slice of email addresses
)

// this variable is accessed in http.go
var db *sqlx.DB

func main() {

	if !processFlags() { // no need to proceed
		return
	}

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

// processFlags returns true if processing should continue, false otherwise
func processFlags() bool {
	var configFileName string
	var versionFlag bool
	var helpFlag bool
	var generateConfig bool
	var testMail bool

	flag.StringVarP(&configFileName, "config", "c", "", "path and name of the config file")
	flag.BoolVarP(&versionFlag, "version", "V", false, "displays version information")
	flag.BoolVarP(&verbose, "verbose", "v", false, "outputs extra information")
	flag.BoolVarP(&helpFlag, "help", "?", false, "displays usage help")
	flag.BoolVarP(&generateConfig, "generate-config", "g", false, "prints a default config file to standard output")
	flag.BoolVarP(&testMail, "test-mail", "m", false, "sends a test email to the configured mail server")
	flag.Parse()

	if versionFlag {
		fmt.Fprintf(os.Stderr, "%s version %s\n", os.Args[0], version)
		fmt.Fprintln(os.Stderr, "\nCopyright (c) 2015 Jon Carlson.  All rights reserved.")
		fmt.Fprintln(os.Stderr, "Use of this source code is governed by the MIT license")
		fmt.Fprintln(os.Stderr, "that can be found here: http://opensource.org/licenses/MIT")
		return false
	}

	if helpFlag {
		usage()
		return false
	}

	if generateConfig {
		generateConfigurationFile()
		return false
	}

	if len(configFileName) > 0 {
		processConfigFile(configFileName)
	}

	if testMail {
		testMailConfig()
		return false
	}

	return true
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s --config <config-file> \n", os.Args[0])
	fmt.Fprintf(os.Stderr, "       %s --generate-config=true > web.config \n", os.Args[0])
	fmt.Fprintln(os.Stderr, `
Program flags are:
  -?, --help            : prints a summary of the arguments 
  -V, --version         : prints the version of this program
  -v, --verbose         : prints additional lines to standard output
  -c, --config          : name and path of config file (required)
  -g, --generate-config : prints an example config file to standard output
  -m, --test-mail       : sends a test alert email using the configured settings
`)
}

// testMailConfig sends a test email to the configured addresses
func testMailConfig() {
	if len(mailHost) == 0 {
		fmt.Fprintln(os.Stderr, "Error, there is no mail host configured to send a test email to.")
		return
	}
	if len(mailTo) == 0 {
		fmt.Fprintln(os.Stderr, "Error, there is no 'mailTo' address to send a test email to.")
		return
	}

	// Send the test email
	err := sendMail("Test email from web-mon", "Receiving this email means your mail configuration is working")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Test email error:", err)
		return
	}

	fmt.Println("Test email sent")
}
