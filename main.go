package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/countdown/numbers"
)

const (
	// The arrangement of Large and Small numbers in puzzle input.
	//
	// Can have at most 4 Large.
	//
	// Default is 6 characters long.
	puzzleInput = "LSSSSS"

	// Interval of type [l, u).
	targetLowerBound = 100
	targetUpperBound = 1000
)

func main() {
	errString := "ERROR: %v\n"
	var done bool = false

	c := numbers.NewConfig(
		puzzleInput,
		targetLowerBound,
		targetUpperBound,
	)

	p, err := numbers.NewPuzzle(c)
	if err != nil {
		done = true
		fmt.Printf(errString, fmt.Errorf("failed to create puzzle: %w", err))
	}

	scanner := bufio.NewScanner(os.Stdin)

	for !done {
		err := loop(&p)
		if err != nil {
			fmt.Printf(errString, fmt.Errorf("failed to generate next puzzle: %w", err))
		}

		fmt.Println("")
		scanner.Scan()
	}
}

// loop prints the current instance of p, and then refreshes it for the next round.
func loop(p *numbers.Puzzle) error {
	fmt.Println("Countdown Numbers")

	p.Print()

	err := p.Refresh()
	if err != nil {
		return err
	}

	return nil
}
