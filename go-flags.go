//
// Gives an example of basic usages for the go-flags package:
//    version, help, verbose flags
//
// go run go-flags.go -?
// go run go-flags.go -V
// go run go-flags.go --verbose
// go run go-flags.go -v -- -hi mom
//

package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

type Options struct {
	Verbose []bool `short:"v" long:"verbose" description:"verbose output"`
	Help    bool   `short:"?" long:"help" description:"help text"`
	Version bool   `short:"V" long:"version" description:"show version info"`
}

const (
	version = "0.1"
)

var (
	options Options
	parser  = flags.NewParser(&options, flags.PrintErrors|flags.PassDoubleDash)
)

func main() {
	parser.Usage = "[Options] filename"

	args, err := parser.Parse()
	if err != nil {
		fmt.Println("Error parsing flags", err)
		os.Exit(1)
	}
	fmt.Printf("The remaining args are: %v\n", args)

	if options.Help {
		parser.WriteHelp(os.Stderr)
		os.Exit(0)
	}

	if len(options.Verbose) > 1 {
		fmt.Println("verbose+")
	} else if len(options.Verbose) > 0 {
		fmt.Println("verbose")
	}

	if options.Version {
		fmt.Fprintf(os.Stderr, "%s version %s\n", os.Args[0], version)
		fmt.Fprintln(os.Stderr, "Copyright (c) 2015 Jon Carlson.  All rights reserved.")
		fmt.Fprintln(os.Stderr, "Use of this source code is governed by the MIT license")
		fmt.Fprintln(os.Stderr, "that can be found here: http://opensource.org/licenses/MIT")
		os.Exit(0)
	}
}
