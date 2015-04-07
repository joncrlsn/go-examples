package main

import (
	"fmt"
	//"log"
	"math/rand"
	"time" // or "runtime"
)

// max24Bits is the max number that can fit in 24 bits (2^24)
var max24Bits = 16777216     // 16 million
var max39Bits = 549755813888 // 549 billion

//  janFirst2015 is the 100ths of a second between Jan 1, 1970 and Jan 1, 2015
var janFirst2015 = time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC).UnixNano() / 1000 / 1000 / 10

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("=== guid ===")

	now := time.Now().UTC().UnixNano()
	fmt.Printf("time.Now.UnixNano():\n%v\n%b\n", now, now)

	Guid()
}

// Guid returns an pseudo globally unique identifier as an int64
// The first 5 bytes are the time portion (100ths of a second since Jan 1, 2015)
// The last 3 bytes are a random number
func Guid() int64 {
	// now is 100ths of a second since Jan 1, 1970
	now := time.Now().UTC().UnixNano() / 1000 / 1000 / 10
	fmt.Printf("== now 100ths:\n%v\n%b\n", now, now)

	// subtract now from Jan 1,2015
	timePortion := now - janFirst2015
	fmt.Printf("== now - jan2015 100ths:\n%v\n%b\n", timePortion, timePortion)

	// bitshift time portion by 24 bits to the left
	timePortion = timePortion << 24
	fmt.Printf("== time shifted 24 bits:\n%v\n%b\n", timePortion, timePortion)

	random24 := int64(rand.Intn(max24Bits))
	fmt.Printf("== random 24 bits:\n%v\n%b\n", random24, random24)

	// bit or the time portion and the random24 portion
	guid := timePortion | random24

	fmt.Printf("== guid:\n%v\n%b\n", guid, guid)
	return guid
}
