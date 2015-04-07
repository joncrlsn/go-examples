package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(RandomString(10))
}

// RandomString returns a random string of capital letters of the given length
func RandomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = byte(RandomInt(65, 90))
	}
	return string(bytes)
}

// RandomInt returns a random integer between the two numbers.
// min is inclusive, max is exclusive
func RandomInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
