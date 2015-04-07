package main

import "fmt"
import "os"

//import "strings"
import "path/filepath"

func main() {
	fmt.Println("--- env ---")
	// To set a key/value pair, use `os.Setenv`. To get a
	// value for a key, use `os.Getenv`. This will return
	// an empty string if the key isn't present in the
	// environment.
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("DBHOST:", os.Getenv("DBHOST"))
	fmt.Println("HOME:", os.Getenv("HOME"))

	cwd, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	fmt.Println("Current Working Directory:", cwd)

	// Use `os.Environ` to list all key/value pairs in the
	// environment. This returns a slice of strings in the
	// form `KEY=value`. You can `strings.Split` them to
	// get the key and value. Here we print all the keys.
	//	fmt.Println()
	//	for _, e := range os.Environ() {
	//		pair := strings.Split(e, "=")
	//		fmt.Printf("%s=%s    ", pair[0], pair[1])
	//	}
	fmt.Println()
}
