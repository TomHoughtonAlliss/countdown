package numbers

import "math/rand"

type getter func() int

var largeNums []int = []int{25, 50, 75, 100}

func newLargeGetter() getter {
	n := make([]int, len(largeNums))
	copy(n, largeNums)

	return func() int {
		if len(n) == 0 {
			panic("large number set empty")
		}

		i := rand.Intn(len(n))
		num := n[i]
		n = append(n[:i], n[i+1:]...)
		return num
	}

}

func getLarge() int {
	return 0
}
