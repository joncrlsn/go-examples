package main

import (
	"bufio"
	"fmt"
	"os"
)

// Prompt user for input and repeat it back
func main() {
	// Here is one way to read
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println("You entered:", text)

	// Read some more with the same reader
	fmt.Print("Enter more text: ")
	textmore, _ := reader.ReadString('\n')
	fmt.Println("You entered:", textmore)

	// Here is another way
	fmt.Print("Enter even more text: ")
	text2 := ""
	fmt.Scanln(&text2)
	fmt.Println("You entered:", text2)

	// Not sure about this one
	//ln := ""
	//fmt.Sscanln("%v", ln)
	//fmt.Println(ln)
}
