package numbers

import (
	"fmt"

	"github.com/countdown/number-puzzle/helpers"
)

// Puzzle stores an array of usable numbers, the target value, and the configuration used to build this instance.
type Puzzle struct {
	config  Config
	Numbers []int
	Target  int
}

// Puzzle.Refresh uses the original input to create and set a new puzzle.
func (p *Puzzle) Refresh() error {
	nums, err := parseInput(p.config.Input)
	if err != nil {
		return fmt.Errorf("failed to parse numbers: %w", err)
	}

	p.Numbers = nums
	p.Target = newTarget(p.config)

	return nil
}

// Puzzle.Print gives a nice output of the current puzzle.
func (p *Puzzle) Print() {
	fmt.Println("Numbers Puzzle")
	fmt.Printf("%v -> %v\n", helpers.CommaSeparate(p.Numbers), p.Target)
}
