package checker

import (
	"fmt"

	"github.com/countdown/helpers"
)

// Checker is used to check a user's mathematical steps to reach the final answer.
type Checker struct {

	// original stores the original input values, so we can reset the Checker at any time.
	original []int

	// current stores the current usable values, so the progress at any point is stored here.
	current []int

	// expressions store the strings used to reach each value.
	expressions map[int]string

	// target is the value we're aiming for.
	target int
}

func NewChecker(n []int, t int) Checker {
	return Checker{
		original:    n,
		current:     n,
		expressions: map[int]string{},
		target:      t,
	}
}

func (c *Checker) Print() {
	fmt.Printf("Target: %v\n", c.target)
	fmt.Printf("Available: %v\n", helpers.CommaSeparate(c.current))
}

func (c *Checker) Expression(expr string) (bool, error) {
	e, err := NewExpression(expr)
	if err != nil {
		return false, fmt.Errorf("failed to create expression: %w", err)
	}

	res, err := e.compute()
	if err != nil {
		return false, fmt.Errorf("failed to compute: %w", err)
	}

	if res == c.target {
		return true, nil
	}

	c.current, err = helpers.Remove(c.current, e.first)
	if err != nil {
		return false, fmt.Errorf("failed to remove element: %w", err)
	}

	c.current, err = helpers.Remove(c.current, e.second)
	if err != nil {
		return false, fmt.Errorf("failed to remove element: %w", err)
	}

	c.current = append(c.current, res)

	return false, nil
}
