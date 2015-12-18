//
// Copyright (c) 2015 Jon Carlson.  All rights reserved.
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.
//
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var propertySplittingRegex = regexp.MustCompile(`\s*=\s*`)
var commaSplittingRegex = regexp.MustCompile(`\s*,\s*`)

// processConfigFile reads the properties in the given file and assigns them to global variables
func processConfigFile(fileName string) {
	if verbose {
		fmt.Println("Processing config file:", fileName)
	}
	props, err := _readConfigFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading config file "+fileName+":", err)
		os.Exit(1)
	}
	_processConfig(props)
}

func _processConfig(props map[string]string) {

	var intVal int
	var strVal string
	var boolVal bool
	var ok bool

	if boolVal, ok = _boolValue(props, "verbose"); ok {
		verbose = boolVal
		if verbose {
			fmt.Println("verbose:", verbose)
		}
	}

	// Configurable variables
	//var (
	//    version        = "0.1"
	//    verbose        = false
	//    dbType         = "sqlite3"
	//    dbUrl          = "web.db"
	//    migrateSqlPath = "./migrations"
	//    httpPort       = 8080
	//    httpsPort      = 8443
	//    certFile       = "https-cert.pem"
	//    keyFile        = "https-key.pem"
	//    //certFile = ""
	//    //keyFile  = ""
	//)
	if strVal, ok = props["dbType"]; ok {
		dbType = strVal
		fmt.Println("dbType:", dbType)
	}

	if strVal, ok = props["dbUrl"]; ok {
		dbUrl = strVal
		fmt.Println("dbUrl:", dbUrl)
	}

	if strVal, ok = props["migrateSqlPath"]; ok {
		migrateSqlPath = strVal
		fmt.Println("migrateSqlPath:", migrateSqlPath)
	}

	if intVal, ok = _intValue(props, "httpPort"); ok {
		httpPort = intVal
		fmt.Println("httpPort:", httpPort)
	}

	if intVal, ok = _intValue(props, "httpsPort"); ok {
		httpsPort = intVal
		fmt.Println("httpsPort:", httpsPort)
	}

	if strVal, ok = props["certFile"]; ok {
		certFile = strVal
		fmt.Println("certFile:", certFile)
	}

	if strVal, ok = props["keyFile"]; ok {
		keyFile = strVal
		fmt.Println("keyFile:", keyFile)
	}

	if strVal, ok = props["mailHost"]; ok {
		mailHost = strVal
		fmt.Println("mailHost:", mailHost)
	}

	if intVal, ok = _intValue(props, "mailPort"); ok {
		mailPort = intVal
		fmt.Println("mailPort:", mailPort)
	}

	if strVal, ok = props["mailUsername"]; ok {
		mailUsername = strVal
		fmt.Println("mailUsername:", mailUsername)
	}
	if strVal, ok = props["mailPassword"]; ok {
		mailPassword = strVal
		fmt.Println("mailPassword: *******")
	}
	if strVal, ok = props["mailFrom"]; ok {
		mailFrom = strVal
		fmt.Println("mailFrom:", mailFrom)
	}
	if strVal, ok = props["mailTo"]; ok {
		mailTo = commaSplittingRegex.Split(strVal, -1)
		fmt.Println("mailTo:", mailTo)
	}

}

// generateConfigurationFile prints an example configuration file to standard output
func generateConfigurationFile() {
	fmt.Println(`# web configuration file.  Uncomment the values you change:
# ======================
# Web app configuration
# ======================

# Just sqlite3 is tested, but others could be supported too
dbType         = sqlite3

# The name of the sqlite3 file
dbUrl          = myapp.db

# Directory that holds the database migration SQL
migrateSqlPath = ./migrations

# Ports
httpPort       = 8080
httpsPort      = 8443

# SSL files
certFile       = https-cert.pem
keyFile        = https-key.pem

# ===================
# Mail configuration
# ===================

# mailHost     = localhost
# mailPort     = 25
# mailUsername =
# mailPassword =

# An email address to be used as the "from" address in emails
# mailFrom =

# A comma-separated list of email addresses that will receive alert emails
# mailTo =
`)
}

// readLinesChannel reads a text file line by line into a channel.
func _readLinesChannel(filePath string) (<-chan string, error) {
	c := make(chan string)
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	go func() {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			c <- scanner.Text()
		}
		close(c)
	}()
	return c, nil
}

// readConfigFile reads name-value pairs from a properties file
func _readConfigFile(fileName string) (map[string]string, error) {
	c, err := _readLinesChannel(fileName)
	if err != nil {
		return nil, err
	}

	properties := make(map[string]string)
	for line := range c {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") {
			// Ignore this line
		} else if len(line) == 0 {
			// Ignore this line
		} else {
			parts := propertySplittingRegex.Split(line, 2)
			properties[parts[0]] = parts[1]
		}
	}

	return properties, nil
}

func _intValue(props map[string]string, name string) (int, bool) {
	if value, ok := props[name]; ok {
		_intValue, err := strconv.Atoi(value)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid integer config value for %s: %s \n", name, value)
			return 0, false
		}
		return _intValue, true
	}

	return 0, false
}

func _boolValue(props map[string]string, name string) (bool, bool) {
	if value, ok := props[name]; ok {
		_boolValue, err := strconv.ParseBool(value)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid bool config value for %s: %s \n", name, value)
			return false, false
		}
		return _boolValue, true
	}
	return false, false
}
