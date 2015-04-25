//
// Provides examples of converting between data types
//
package main

import (
	"fmt"
	"strconv"
)

func main() {

	// convert string to int
	i, _ := strconv.Atoi("100")

	// convert int to string
	var str string = strconv.Itoa(i)
	fmt.Println("str is back to a string", str)
}
