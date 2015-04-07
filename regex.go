package main

import (
	"bytes"
	"fmt"
	"log"
	"regexp"
	"strings"
)

func main() {
	regexp1()
	regexp2()
}

func regexp1() {
	fmt.Println("--- regexp1 ---")

	includeRegex, err := regexp.Compile(`^\s*include\("(\\\"|[^"])+"\);`)
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range strings.Split(`
foo
include "abc.def"
include("file.js");
            include "me\"to\""
            include("please\"!\"");
        nothing here          
`, "\n") {
		if includeRegex.Match([]byte(line)) {
			includeFile := includeRegex.FindString(line)
			fmt.Println("INCLUDE", includeFile)
		} else {
			fmt.Printf("no match for \"%s\"\n", line)
		}
	}
}

func regexp2() {
	fmt.Println("--- regexp2 ---")

	// This tests whether a pattern matches a string.
	match1, _ := regexp.MatchString(`(?i)(^|\s)FROM\s`, " from t_user")
	fmt.Println(match1)

	// This tests whether a pattern matches a string.
	match, _ := regexp.MatchString(`p([a-z]+)ch`, "peach")
	fmt.Println(match)

	// Above we used a string pattern directly, but for
	// other regexp tasks you'll need to `Compile` an
	// optimized `Regexp` struct.
	r, _ := regexp.Compile("p([a-z]+)ch")

	// Many methods are available on these structs. Here's
	// a match test like we saw earlier.
	fmt.Println(r.MatchString("peach"))

	// This finds the match for the regexp.
	fmt.Println(r.FindString("peach punch"))

	// The also finds the first match but returns the
	// start and end indexes for the match instead of the
	// matching text.
	fmt.Println(r.FindStringIndex("peach punch"))

	// The `Submatch` variants include information about
	// both the whole-pattern matches and the submatches
	// within those matches. For example this will return
	// information for both `p([a-z]+)ch` and `([a-z]+)`.
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// Similarly this will return information about the
	// indexes of matches and submatches.
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// The `All` variants of these functions apply to all
	// matches in the input, not just the first. For
	// example to find all matches for a regexp.
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// These `All` variants are available for the other
	// functions we saw above as well.
	fmt.Println(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	// Providing a non-negative integer as the second
	// argument to these functions will limit the number
	// of matches.
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// Our examples above had string arguments and used
	// names like `MatchString`. We can also provide
	// `[]byte` arguments and drop `String` from the
	// function name.
	fmt.Println(r.Match([]byte("peach")))

	// When creating constants with regular expressions
	// you can use the `MustCompile` variation of
	// `Compile`. A plain `Compile` won't work for
	// constants because it has 2 return values.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// The `regexp` package can also be used to replace
	// subsets of strings with other values.
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// The `Func` variant allows you to transform matched
	// text with a given function.
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
