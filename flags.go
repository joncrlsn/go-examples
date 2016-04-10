package main

import "flag"
import "fmt"
import "os"

func main() {
	fmt.Println("--- flag ---")
	var port int
	var hostPtr = flag.String("host", "localhost", "database host")
	flag.IntVar(&port, "port", 5432, "database port")
	flag.Parse()
	fmt.Printf("host = %s\n", *hostPtr)
	fmt.Printf("port = %d\n", port)
	fmt.Println("all args:", os.Args)
	fmt.Println("remaining args:", flag.Args())
	fmt.Printf("remaining args: %v\n", flag.Args())
	fmt.Println()
}

func usage() {
	fmt.Println("--- usage ---")
	fmt.Fprintf(os.Stderr, "usage: %s [-bloviate <string>] [-port <int>]\n", os.Args[0])
	flag.PrintDefaults()
	os.Exit(2)
}
