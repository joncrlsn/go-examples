//
// Gives an example of basic usages for the go-flags package:
//    version, help, and verbose flags
//
// go run go-flags.go -?
// go run go-flags.go --version
// go run go-flags.go --verbose
// go run go-flags.go -v -- -hi mom
//
package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

const (
	version = "0.1"
)

// The main options struct
type Options struct {
	Verbose []bool      `short:"v" long:"verbose" description:"verbose output"`
	help    HelpOptions `group:"Help Options"`
	//	Ini     string `short:"f" long:"file" description:"ini file name"`

	// Example of a required flag
	Name string `short:"n" long:"name" description:"a name is required" required:"true"`
}

// This is a sub-struct of the Options struct.  It holds the 2 common help options (help and version)
type HelpOptions struct {
	Help    bool `short:"?" long:"help" description:"help text" default:false`
	Version bool `short:"V" long:"version" description:"show version info" default:false`
}

var (
	options Options
	parser  = flags.NewParser(&options, flags.PrintErrors|flags.PassDoubleDash)
)

// init parses the flags that are passed in
func init() {
	parser.Usage = "[Options] filename"

	if false {
		// Parse an ini file instead of the command-line option flags
		iniParser := flags.NewIniParser(parser)
		err := iniParser.ParseFile("go-flags.ini")
		if err != nil {
			fmt.Println("error parsing ini file: ", err)
			os.Exit(1)
		}
	} else {
		// Parse the command-line options
		args, err := parser.Parse()
		if err != nil {
			// errors are already written to system err
			os.Exit(1)
		}
		fmt.Printf("The remaining args are: %v\n", args)

		// Write out an ini file using the given command-line option flags
		iniParser := flags.NewIniParser(parser)
		iniParser.WriteFile("go-flags.ini", flags.IniIncludeDefaults|flags.IniCommentDefaults|flags.IniIncludeComments)
	}
}

// main function
func main() {

	// Write out the help
	if options.help.Help {
		parser.WriteHelp(os.Stderr)
		os.Exit(0)
	}

	// Write out the version and copyright information
	if options.help.Version {
		fmt.Fprintf(os.Stderr, "%s version %s\n", os.Args[0], version)
		fmt.Fprintln(os.Stderr, "Copyright (c) 2015 Jon Carlson.  All rights reserved.")
		fmt.Fprintln(os.Stderr, "Use of this source code is governed by the MIT license")
		fmt.Fprintln(os.Stderr, "that can be found here: http://opensource.org/licenses/MIT")
		os.Exit(0)
	}

	// Print the values of the flags
	fmt.Printf("Verbose: %v\n", options.Verbose)
	fmt.Printf("Help: %t\n", options.help.Help)
	fmt.Printf("Version: %v\n", options.help.Version)
	fmt.Printf("Name: %s\n", options.Name)

	// Write out an ini file that could be used in lieu of option flags
	iniParser := flags.NewIniParser(parser)
	iniParser.WriteFile("go-flags.ini", flags.IniIncludeDefaults|flags.IniCommentDefaults|flags.IniIncludeComments)
}
