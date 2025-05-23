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

	// A string for formatting errors in main.
	errString = "ERROR: %v\n"
)

func main() {
	c := numbers.NewConfig(
		puzzleInput,
		targetLowerBound,
		targetUpperBound,
	)

	p, err := numbers.NewPuzzle(c)
	if err != nil {
		fmt.Printf(errString, fmt.Errorf("failed to create puzzle: %w", err))
		return
	}

	puzzleLoop(&p)
}

func puzzleLoop(p *numbers.Puzzle) {
	var done bool
	var err error

	for !done {
		p.Print()

		err = check(p)
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

	c := checker.NewChecker(
		p.Numbers,
		p.Target,
	)

	err := solvingLoop(&c)
	if err != nil {
		return err
	}

	return nil
}

func solvingLoop(c *checker.Checker) error {
	var done bool
	var err error
	scanner := bufio.NewScanner(os.Stdin)

	for !done {
		fmt.Println()
		c.Print()
		fmt.Print("> ")

		if scanner.Scan() {
			done, err = c.HandleInput(scanner.Text())
			if err != nil {
				return fmt.Errorf("error occured on input: %w", err)
			}
		}
	}

	return nil
}
