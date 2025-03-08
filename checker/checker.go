package checker

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/countdown/helpers"
)

// Checker is used to check a user's mathematical steps to reach the final answer.
type Checker struct {

	// original stores the original input values, so we can reset the Checker at any time.
	original []int

	// usable stores the usable usable values, so the progress at any point is stored here.
	usable []int

	// expressions store the strings used to reach each value.
	expressions []expression

	// target is the value we're aiming for.
	target int
}

// NewChecker creates and returns a new Checker object.
func NewChecker(n []int, t int) Checker {
	new := make([]int, len(n))
	copy(new, n)
	return Checker{
		original:    n,
		usable:      new,
		expressions: []expression{},
		target:      t,
	}
}

// Checker.Print gives the current puzzle instance, with up-to-date usable numbers.
func (c *Checker) Print() {
	fmt.Printf("Target:\t\t%v\n", c.target)
	fmt.Printf("Available:\t%v\n", helpers.CommaSeparate(c.usable))
}

// Checker.Expression takes a string (e.g. 5 + 6) and handles it.
//
// The result is calculated and checked against the target.
//
// It handles removing used values from our usable numbers array, as well as adding the result.
func (c *Checker) Expression(expr string) (solved bool, exprErr error, err error) {
	e, err := NewExpression(expr)
	if err != nil {
		return false, fmt.Errorf("failed to create expression: %w", err), nil
	}

	res, err := e.compute()
	if err != nil {
		return false, nil, fmt.Errorf("failed to compute: %w", err)
	}

	e.result = res

	c.expressions = append(c.expressions, e)

	if res == c.target {
		return true, nil, nil
	}

	c.usable, err = helpers.Remove(c.usable, e.first)
	if err != nil {
		return false, nil, fmt.Errorf("failed to remove element: %w", err)
	}

	c.usable, err = helpers.Remove(c.usable, e.second)
	if err != nil {
		return false, nil, fmt.Errorf("failed to remove element: %w", err)
	}

	c.usable = append(c.usable, res)

	return false, nil, nil
}

// Checker.Reset sets all values to their inital configuration.
func (c *Checker) Reset() {
	new := make([]int, len(c.original))

	copy(new, c.original)

	c.usable = new
	c.expressions = []expression{}
}

// Checker.ToString recurs on our expressions to fully expand our target into the initial values, in the structure given by the user.
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

// Checker.HandleInput takes a user-inputted string and handles it accordingly.
func (c *Checker) HandleInput(input string) (done bool, err error) {
	// TODO break this down more
	switch input {
	case "reset", "restart":
		c.Reset()

	case "next", "skip", "pass":
		fmt.Print("\nSkipping\n\n")
		done = true

	case "quit", "exit":
		done = true
		err = errors.New("exited by user")

	default:
		done, err = c.NewExpression(input)
		if err != nil {
			return
		}
	}

	return done, err
}

// Checker.NewExpression handles the parsing and creation of a new user expression.
func (c *Checker) NewExpression(input string) (bool, error) {
	solved, exprErr, err := c.Expression(input)
	if err != nil {
		return true, fmt.Errorf("failed to parse expression: %w", err)
	}

	if exprErr != nil {
		// User inputted a bad command or expression. Continue.
		fmt.Println()
		fmt.Println("Unrecognised or unparsable input.")
		return false, nil
	}

	return c.HandleSolve(solved), nil
}

// Checker.HandleSolve checks if the target value has been reached, and if so, returns the final nested equation.
func (c *Checker) HandleSolve(solved bool) (resolved bool) {
	if !solved {
		return
	}

	fmt.Println()
	fmt.Printf("%v = %v\n", c.target, c.ToString())
	fmt.Println()

	return solved
}
