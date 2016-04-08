//
// password is an example of reading a password from standard in
// without echoing it back to the user.
//
package main

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

func main() {
	stdin := 1
	fmt.Fprintf(os.Stdin, "Enter a test password: ")

	// Get current state of terminal
	//    s, err := terminal.MakeRaw(stdin)
	//    check(err, "making raw terminal, Saving old terminal state")
	//    defer terminal.Restore(stdin, s)

	// Read password from stdin
	b, err := terminal.ReadPassword(stdin)
	check(err, "reading from terminal")

	fmt.Println("Your test password is:", string(b))
}

func check(err error, action string) {
	if err != nil {
		fmt.Printf("Error %s: %v\n", action, err)
	}
}
