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
}
