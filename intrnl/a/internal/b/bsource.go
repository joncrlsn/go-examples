package b

import (
	"fmt"
	"github.com/joncrlsn/go-examples/intrnl/x"
)

var (
	InternalToA = "b is internal to a"
)

func PrintStuff() {
	fmt.Println("Xvar =", x.Xvar)
}
