//
// Run with "go run channel.go"
//
package main

import "fmt"

type Org struct {
	id   int16
	name string
}

func main() {
	fmt.Println("--- channel ---")
	// Build the channel
	c := make(chan Org)

	// A go routine that sends to the channel
	go fillChannel(c)

	// Read from the channel
	for org := range c {
		fmt.Printf("id:%d, name: %s \n", org.id, org.name)
	}

	x := <-c
	fmt.Printf("id:%d, name: %s \n", x.id, x.name)
	x = <-c
	fmt.Printf("id:%d, name: %s \n", x.id, x.name)
}

func fillChannel(c chan<- Org) {
	for i := 0; i < 8; i++ {
		org := Org{id: int16(i), name: fmt.Sprintf("org%d", i)}
		c <- org
	}
	// Tell the receiver that we're done sending
	close(c)
}

// =================================================
