package main

import (
	"fmt"
)

type Thing struct {
	value string
}

func main() {
	pointerTest1()
}

// Test passing structs to functions that set values on the struct
func pointerTest1() {
	thing := Thing{}
	setValue(&thing, "my new value")
	fmt.Println("thing.value: ", thing.value)
}

func setValue(thing *Thing, myVal string) {
	thing.value = myVal
}
