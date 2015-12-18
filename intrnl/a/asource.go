package a

import (
	"fmt"
	"github.com/joncrlsn/go-examples/intrnl/a/internal/b"
)

var (
	Avar = "avar"
)

func PrintStuff() {
	fmt.Println("Stuff internal to A =", b.InternalToA)
	b.PrintStuff()
}
