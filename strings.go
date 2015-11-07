// The standard library's `strings` package provides many
// useful string-related functions. Here are some examples
// to give you a sense of the package.
package main

import "fmt"
import "strings"

// alias `fmt.Println` to a shorter name
var p = fmt.Println

func main() {
	fmt.Println("--- strings ---")

	p("Len: ", len("hello"))
	p("Char:", "hello"[1])

	spacey := " x "

	fmt.Println(strings.Fields("Hello GoLang, my friend"))
	fmt.Println("Hello " + "GoLang, my friend")
	fmt.Print("Before TrimSpace: '" + spacey + "'")
	fmt.Println("  After TrimSpace: '" + strings.TrimSpace(spacey) + "'")

	// Here's a sample of the functions available in
	// `strings`. Note that these are all functions from
	// package, not methods on the string object itself.
	// This means that we need pass the string in question
	// as the first argument to the function.
	fmt.Println("Contains:  ", strings.Contains("test", "es"))
	fmt.Println("Count:     ", strings.Count("test", "t"))
	fmt.Println("HasPrefix: ", strings.HasPrefix("test", "te"))
	fmt.Println("HasSuffix: ", strings.HasSuffix("test", "st"))
	fmt.Println("Index:     ", strings.Index("test", "e"))
	fmt.Println("Join:      ", strings.Join([]string{"a", "b"}, "-"))
	fmt.Println("Repeat:    ", strings.Repeat("a", 5))
	fmt.Println("Replace:   ", strings.Replace("foo", "o", "0", -1))
	fmt.Println("Replace:   ", strings.Replace("foo", "o", "0", 1))
	fmt.Println("Split:     ", strings.Split("a-b-c-d-e", "-"))
	fmt.Println("ToLower:   ", strings.ToLower("TEST"))
	fmt.Println("ToUpper:   ", strings.ToUpper("test"))
	fmt.Println()

	// You can find more functions in the [`strings`](http://golang.org/pkg/strings/)
	// package docs.

	convertByteArrayToString()
}

func convertByteArrayToString() {
	src := []byte{'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd'}

	str := string(src[:])
	fmt.Println(str)

	str = string(src)
	fmt.Println(str)
}

// iterate over the runes in a string
func iterateOverString() {
	str := "Hello"
	fmt.Println()
	for i, roon := range str {
		fmt.Println(i, roon, string(roon))
	}
}
