package svc

import (
	"math/rand"
	"time"
)

const (
	// max24Bits is the max number that can fit in 24 bits (2^24)
	max24Bits = 16777216     // 16 million
	max39Bits = 549755813888 // 549 billion
)

var (
	//  janFirst2015 is the 100ths of a second between Jan 1, 1970 and Jan 1, 2015
	janFirst2015 = time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC).UnixNano() / 1000 / 1000 / 10
)

// Guid provides pseudo globally unique identifiers as an int64
// The first 5 bytes are the time portion (100ths of a second since Jan 1, 2015)
// The last 3 bytes are a random number
type Guids struct {
	initialized bool
}

// NewGuidMaker instantiates a GuidMaker
func NewGuids() *Guids {
	return &Guids{false}
}

func (guids *Guids) next() int64 {
	if !guids.initialized {
		rand.Seed(time.Now().UTC().UnixNano())
		guids.initialized = true
	}

	// now is 100ths of a second since Jan 1, 1970
	now := time.Now().UTC().UnixNano() / 1000 / 1000 / 10

	// subtract now from Jan 1,2015
	timePortion := now - janFirst2015

	// bitshift time portion by 24 bits to the left (we are only losing zero bits)
	timePortion = timePortion << 24

	random24 := int64(rand.Intn(max24Bits))

	// bit or the time portion and the random24 portion
	guid := timePortion | random24

	return guid
}
