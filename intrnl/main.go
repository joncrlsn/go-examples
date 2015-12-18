package main

//
// Example of using an internal package.
// This package *can* access the stuff in "a", but not the stuff in "a/internal"
//

import (
	"fmt"
	"github.com/joncrlsn/go-examples/intrnl/a"
	//	"github.com/joncrlsn/go-examples/intrnl/a/internal/b" // (This is NOT allowed)
)

func main() {

	// we can access variables and functions in a
	fmt.Println("a.Avar =", a.Avar)
	a.PrintStuff()

	// but we cannot access varibles and functions in b
	// fmt.Println("b.InternalToA =", b.InternalToA)
	// b.PrintStuff()
}
