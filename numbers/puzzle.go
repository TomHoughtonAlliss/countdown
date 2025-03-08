package numbers

import (
	"fmt"
	"strconv"
	"strings"
)

// Puzzle stores an array of numbers and the target value.
type Puzzle struct {
	input   string
	Numbers []int
	Target  int
}

// Puzzle.Refresh uses the original input to create and set a new puzzle.
func (p *Puzzle) Refresh() error {
	nums, err := parseInput(p.input)
	if err != nil {
		return fmt.Errorf("failed to parse numbers: %w", err)
	}

	p.Numbers = nums
	p.Target = newTarget()

	return nil
}

// Puzzle.Print gives a nice output of the current puzzle.
func (p *Puzzle) Print() {
	strNums := make([]string, len(p.Numbers))
	for i, n := range p.Numbers {
		s := strconv.Itoa(n)

		strNums[i] = s
	}

	fmt.Printf("%v -> %v\n", strings.Join(strNums, ", "), p.Target)
}
