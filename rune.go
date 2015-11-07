package main

import (
	"fmt"
)

func main() {

	// Iterate over the runes in a string
	str := "Hello äåéö€"
	fmt.Println()
	for i, roon := range str {
		fmt.Println(i, roon, string(roon))
	}
}
