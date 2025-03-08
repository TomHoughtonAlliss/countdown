package checker

import "fmt"

func (e *expression) compute() (int, error) {
	var res int
	var err error

	switch e.operator {
	case "*":
		res = e.first * e.second
	case "/":
		res = e.first / e.second
	case "+":
		res = e.first + e.second
	case "-":
		res = e.first - e.second
	default:
		err = fmt.Errorf("unknown operator %v", e.operator)
	}

	return res, err
}
