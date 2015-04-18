//
// iface gives examples of an interface and a struct that implements it
//
package main

import "fmt"

func main() {
	fmt.Println("--- interface ---")
	var imp Iface = &IfaceImpl{x: 1}
	imp.SayHi()
	imp.SayBye()
	callit(imp)
}

type Iface interface {
	SayHi()
	SayBye()
}

type IfaceImpl struct {
	x int
}

func (i *IfaceImpl) SayHi() {
	fmt.Println("Saying Hi before assignment. x=", i.x)
	i.x = 11
	fmt.Println("Saying Hi after assignment. x=", i.x)
}

func (i *IfaceImpl) SayBye() {
	fmt.Println("Saying Bye. x=", i.x)
}

func callit(iface Iface) {
	iface.SayHi()
	iface.SayBye()
}
