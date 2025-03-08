package numbers

import (
	"errors"
	"fmt"
)

const expectedLength = 6

type getter func() (int, error)

// Puzzle stores an array of numbers and the target value.
type Puzzle struct {
	Numbers []int
	Target  int
}

// NewPuzzle converts puzzle input to a puzzle object.
//
// input string must contain only L or S and be of the expected length.
func NewPuzzle(input string) (Puzzle, error) {
	var dud Puzzle

	if len(input) != expectedLength {
		return dud, errors.New("puzzle input must have a length of 6")
	}

	nums, err := parseInput(input)
	if err != nil {
		return dud, fmt.Errorf("failed to parse numbers: %w", err)
	}

	p := Puzzle{
		Numbers: nums,
		Target:  newTarget(),
	}

	return p, nil
}

// parseInput iterates over input and gets a new value for each character.
//
//  - For an S get a small number.
//
//  - For an L get a large number.
//
// It will error if it encounters a character of neither S nor L.
func parseInput(input string) ([]int, error) {
	var dud []int

	nums := make([]int, expectedLength)
	small := newSmallGetter()
	large := newLargeGetter()

	for _, c := range input {
		var i int
		var err error

		switch c {
		case 'L':
			i, err = small()
		case 'S':
			i, err = large()
		default:
			return dud, fmt.Errorf("unexpected character in input %v", c)
		}

		if err != nil {
			return dud, fmt.Errorf("failed to build input: %w", err)
		}

		nums = append(nums, i)
	}

	return nums, nil
}
