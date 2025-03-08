package numbers

import (
	"errors"
	"fmt"
	"math/rand"
	"slices"
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
			return 0, errors.New("large number set empty")
		}

		i := rand.Intn(len(n))
		num := n[i]
		n = deleteElt(n, i)
		fmt.Printf("length: %v\tcontents: %v\n", len(n), n)
		return num, nil
	}

	return g
}

func deleteElt[T any](s []T, i int) []T {
	j := i + 1
	return slices.Delete(s, i, j)
}
