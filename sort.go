package main

import "fmt"

// http://nerdyworm.com/blog/2013/05/15/sorting-a-slice-of-structs-in-go/

type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}

func main() {
	fmt.Println("sort not implemented yet")
}
