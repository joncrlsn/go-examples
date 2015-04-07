package main

import (
	"fmt"
	"time"
)

// layout shows by example how the reference time should be represented.
const layout = "Jan 2, 2006 at 3:04pm -0700 (MST)"
const layout2 = "2006-01-02 15:04:05"

func main() {
	//t := time.Date(2009, time.November, 10, 15, 0, 0, 0, time.Local)
	t := time.Now()
	fmt.Println(t.Format(layout2))
	fmt.Println(t.UTC().Format(layout2))
}
