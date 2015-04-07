//
// Run with "go run chooseone.go"
//
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	value := ChooseOne("Continue, Quit, or Redo? [cqr]:", "c", "q", "r")
	fmt.Println("You chose:", value)
}

// ChooseOne prompts the user to enter one of n values
// Example: answer := misc.PromptOne("Continue, Quit, or Redo? [cqr]: ", ["c", "q", "r"])
func ChooseOne(prompt string, values ...string) string {
	for {
		fmt.Print(prompt)

		// Read input
		inReader := bufio.NewReader(os.Stdin)
		text, _ := inReader.ReadString('\n')

		// Trim leading and trailing spaces, and convert to lower case
		text = strings.TrimSpace(text)
		if len(text) == 0 {
			return values[0]
		}

		for _, v := range values {
			if text == v {
				return v
			}
		}

		fmt.Println("Invalid entry, try again.")
	}
}
