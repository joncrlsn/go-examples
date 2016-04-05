//
// run with "go run byte.go"
//
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var byteLf = byte(10) // line feed
var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ\n"

func main() {

	fmt.Println(" == by numeric value")
	// Sequentially go from 32 to 128
	for i := 32; i <= 128; i++ {
		b := byte(i)
		if b == byteLf {
			fmt.Printf("byte: lf %v %b \n", b, b)
		} else {
			fmt.Printf("byte: %v %s %b\n", b, []byte{b}, b)
		}
	}

	fmt.Println(" == by character")
	// Read through the bytes in the string
	// See strings.go for example of reading through the runes in a string
	reader := strings.NewReader(str)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println("error: ", err)
			os.Exit(1)
		}
		if b == byteLf {
			fmt.Printf("byte: lf %b %v\n", b, b)
		} else {
			fmt.Printf("byte: %s %b %v\n", []byte{b}, b, b)
		}
	}

	//
	// See strings.go for example of converting byte array to string and back
	//

}
