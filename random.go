package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(RandomString(25))
}

// RandomString returns a random string of letters of the given length
func RandomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		rint := RandomInt(65, 117)
		if rint > 90 {
			rint = rint + 6
		}
		bytes[i] = byte(rint)
	}
	return string(bytes)
}

// RandomInt returns a random integer between the two numbers.
// min is inclusive, max is exclusive
func RandomInt(min int, max int) int {
	return min + rand.Intn(max-min)
}
