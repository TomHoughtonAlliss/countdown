package main

import (
	"fmt"

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

	for !done {
		p.Print()
		if err != nil {
			done = true
			fmt.Printf(errString, err.Error())
		}
	}
}
