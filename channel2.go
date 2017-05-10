//
// Demonstrates an untyped channel that receives multiple types
//
// Run with "go run channel2.go"
//
package main

import (
	"fmt"
	"math"
)

type Org struct {
	id   int16
	name string
}

func main() {
	fmt.Println("--- channel ---")
	// Build the channel
	c := make(chan interface{})

	// A go routine that sends to the channel
	go fillChannel(c)

	// Read from the channel
	for value := range c {
		switch str := value.(type) {
		case string:
			fmt.Printf("string: %s\n", str)
		case Org:
			fmt.Printf("org: %s\n", str.name)
		default:
			fmt.Println("Oops, not a string or an Org")
		}
	}

	//x := <-c
	//fmt.Printf("id:%d, name: %s \n", x.id, x.name)
	//x = <-c
	//fmt.Printf("id:%d, name: %s \n", x.id, x.name)
}

func fillChannel(c chan<- interface{}) {
	for i := 0; i < 8; i++ {
		org := Org{id: int16(i), name: fmt.Sprintf("org%d", i)}
		if math.Mod(float64(i), float64(2)) == 0 {
			c <- org
		} else {
			c <- org.name
		}
	}
	// Tell the receiver that we're done sending
	close(c)
}

// =================================================
