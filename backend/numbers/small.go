package numbers

import "math/rand"

var smallNums []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// newSmallGetter creates a closure that retrieves a random value from the smallNums array.
func newSmallGetter() getter {
	return func() (int, error) {
		i := rand.Intn(len(smallNums))
		return smallNums[i], nil
	}
}
