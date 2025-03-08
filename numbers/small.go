package numbers

import "math/rand"

var smallNums []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func getSmall() int {
	i := rand.Intn(len(smallNums))
	return smallNums[i]
}
