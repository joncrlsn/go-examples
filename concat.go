package main

import "fmt"
import "bytes"

func main() {
	fmt.Println("--- concat ---")
	_doSprint()
	_doBuffer()
}

func _doSprint() {
	a := fmt.Sprint(1, " + ", 2, " == ", 3, " evaluates ", true, ".")
	fmt.Println(a)
}

func _doBuffer() {
	a := []interface{}{"Apple", "Banana", "Coffee", "Donut"}
	var buffer bytes.Buffer
	for i, v := range a {
		buffer.WriteString(fmt.Sprint(i, ":", v, ", "))
	}
	s := buffer.String()
	fmt.Println(s)
}
