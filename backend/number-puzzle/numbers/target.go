package numbers

import (
	"math/rand"
)

// Generates a random integer in [l, u).
func newTarget(c Config) int {
	clippedUpper := c.Upper - c.Lower
	return rand.Intn(clippedUpper) + c.Lower
}
