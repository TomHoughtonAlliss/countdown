package checker

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/countdown/number-puzzle/helpers"
)

var (
	operators = []string{"*", "/", "+", "-"}
)

type expression struct {
	first    int
	second   int
	operator string
	result   int
}

func NewExpression(expr string) (expression, error) {
	var dud expression

	op, err := findOperator(expr)
	if err != nil {
		return dud, fmt.Errorf("could not find operator: %w", err)
	}

	first, second, err := getValues(expr, op)
	if err != nil {
		return dud, fmt.Errorf("could not split into operands: %w", err)
	}

	e := expression{
		first:    first,
		second:   second,
		operator: op,
		result:   0,
	}

	return e, nil
}

func Find(exprs []expression, n int) (int, bool) {
	var dud int

	for i, expr := range exprs {
		if expr.result == n {
			return i, true
		}
	}

	return dud, false
}

func findOperator(expr string) (string, error) {
	arr := helpers.StringToArray(expr)
	for _, o := range operators {
		if helpers.Index(arr, o) != -1 {
			return o, nil
		}
	}

	return "", fmt.Errorf("failed to find operator in %v", expr)
}

func getValues(expr string, op string) (int, int, error) {
	splits := strings.Split(expr, op)

	if len(splits) != 2 {
		return 0, 0, fmt.Errorf("unexpected number of values in %v", splits)
	}

	first := strings.TrimSpace(splits[0])
	second := strings.TrimSpace(splits[1])

	f, err := strconv.Atoi(first)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse %v: %w", first, err)
	}

	s, err := strconv.Atoi(second)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse %v: %w", second, err)
	}

	return f, s, nil
}
