package numbers

import (
	"errors"
	"fmt"
)

const expectedLength = 6

type getter func() (int, error)

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
		input:   input,
		Numbers: nums,
		Target:  newTarget(),
	}

	return p, nil
}

// parseInput iterates over input and gets a new value for each character.
//
//   - For an S get a small number.
//
//   - For an L get a large number.
//
// It will error if it encounters a character of neither S nor L.
func parseInput(input string) ([]int, error) {
	var dud []int

	nums := make([]int, expectedLength)
	small := newSmallGetter()
	large := newLargeGetter()

	for i, c := range input {
		var n int
		var err error

		switch c {
		case 'L':
			n, err = large()
		case 'S':
			n, err = small()
		default:
			return dud, fmt.Errorf("unexpected character in input %v", c)
		}

		if err != nil {
			return dud, fmt.Errorf("failed to build input: %w", err)
		}

		nums[i] = n
	}

	return nums, nil
}
