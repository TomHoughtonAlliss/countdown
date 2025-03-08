package numbers

import (
	"errors"
	"math/rand"

	"github.com/countdown/number-puzzle/helpers"
)

var largeNums []int = []int{25, 50, 75, 100}

// newLargeGetter creates a closure that retrieves a random value from the largeNums array, and deletes it.
//
// The closure will error if called more than four times.
func newLargeGetter() getter {
	n := make([]int, len(largeNums))
	copy(n, largeNums)

	var g getter = func() (int, error) {
		if len(n) == 0 {
			return 0, errors.New("too many large numbers used")
		}

		i := rand.Intn(len(n))
		num := n[i]
		n = helpers.DeleteElt(n, i)

		return num, nil
	}

	return g
}
