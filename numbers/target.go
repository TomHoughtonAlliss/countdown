package numbers

import (
	"math/rand"
)

const (
	lowerBound = 100
	upperBound = 1000
)

// Generates a random integer in [l, u).
func newTarget() int {
	clippedUpper := upperBound - lowerBound
	return rand.Intn(clippedUpper) + lowerBound
}
