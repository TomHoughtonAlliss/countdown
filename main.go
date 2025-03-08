package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/countdown/numbers"
)

func main() {
	errString := "ERROR: %v\n"
	var done bool = false

	i := "LSSSSS"

	p, err := numbers.NewPuzzle(i)
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
