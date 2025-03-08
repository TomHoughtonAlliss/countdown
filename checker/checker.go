package checker

import (
	"fmt"
	"strconv"

	"github.com/countdown/helpers"
)

// Checker is used to check a user's mathematical steps to reach the final answer.
type Checker struct {

	// original stores the original input values, so we can reset the Checker at any time.
	original []int

	// current stores the current usable values, so the progress at any point is stored here.
	current []int

	// expressions store the strings used to reach each value.
	expressions []expression

	// target is the value we're aiming for.
	target int
}

func NewChecker(n []int, t int) Checker {
	new := make([]int, len(n))
	copy(new, n)
	return Checker{
		original:    n,
		current:     new,
		expressions: []expression{},
		target:      t,
	}
}

func (c *Checker) Print() {
	fmt.Printf("Target:\t\t%v\n", c.target)
	fmt.Printf("Available:\t%v\n", helpers.CommaSeparate(c.current))
}

func (c *Checker) Expression(expr string) (solved bool, err error) {
	e, err := NewExpression(expr)
	if err != nil {
		return false, fmt.Errorf("failed to create expression: %w", err)
	}

	res, err := e.compute()
	if err != nil {
		return false, fmt.Errorf("failed to compute: %w", err)
	}

	e.result = res

	c.expressions = append(c.expressions, e)

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

// Checker.Reset sets all values to their inital configuration.
func (c *Checker) Reset() {
	new := make([]int, len(c.original))

	copy(new, c.original)

	c.current = new
	c.expressions = []expression{}
}

// ToString recurs on our expressions to fully expand our target into the initial values, in the structure given by the user.
func (c *Checker) ToString() string {
	var recur func(int) string
	recur = func(n int) string {
		i, found := Find(c.expressions, n)

		if !found {
			return strconv.Itoa(n)
		}

		expr := c.expressions[i]

		c.expressions = helpers.DeleteElt(c.expressions, i)

		f := recur(expr.first)
		s := recur(expr.second)
		o := expr.operator

		return fmt.Sprintf("(%v %v %v)", f, o, s)
	}

	equation := recur(c.target)

	return equation[1 : len(equation)-1]
}

func (c *Checker) HandleInput(input string) (done bool, err error) {
	// TODO break this down more
	switch input {
	case "reset":
		c.Reset()
	case "next":
		fmt.Print("\nSkipping\n\n")
		done = true
	default:
		solved, err := c.Expression(input)
		if err != nil {
			return true, fmt.Errorf("failed to parse expression: %w", err)
		}

		done = c.HandleSolve(solved)
	}

	return done, err
}

func (c *Checker) HandleSolve(solved bool) bool {
	if !solved {
		return false
	}

	fmt.Println()
	fmt.Printf("%v = %v\n", c.target, c.ToString())
	fmt.Println()

	return solved
}
