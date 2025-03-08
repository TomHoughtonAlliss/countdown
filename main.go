package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/countdown/checker"
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

	for !done {
		p.Print()

		err = check(&p)
		if err != nil {
			done = true
			fmt.Printf(errString, fmt.Errorf("failed to generate next puzzle: %w", err))
		}

		err := p.Refresh()
		if err != nil {
			done = true
			fmt.Printf(errString, fmt.Errorf("failed to generate next puzzle: %w", err))
		}
	}
}

// check begins the while loop to allow a user to attempt the puzzle.
func check(p *numbers.Puzzle) error {
	scanner := bufio.NewScanner(os.Stdin)

	c := checker.NewChecker(
		p.Numbers,
		p.Target,
	)

	done := false
	for !done {

		c.Print()
		fmt.Print("> ")

		if scanner.Scan() {
			expr := scanner.Text()

			switch expr {
			case "reset":
				// reset the checker
			default:
				solved, err := c.Expression(expr)
				if err != nil {
					return fmt.Errorf("failed to parse expression: %w", err)
				}
	
				if solved {
					done = true
					// Print the expressions here
				}
			}
		}
	}

	return nil
}
