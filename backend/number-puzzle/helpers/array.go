package helpers

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

// DeleteElt is a wrapper on slices.Delete, to avoid having to call it with two indexes.
func DeleteElt[T any](s []T, i int) []T {
	j := i + 1
	return slices.Delete(s, i, j)
}

// Index fetches the index of elt in slice s.
func Index[T comparable](s []T, elt T) int {
	for i, e := range s {
		if e == elt {
			return i
		}
	}

	return -1
}

// StringToArray turns a string value into an array of single-character strings.
func StringToArray(s string) []string {
	a := make([]string, len(s))
	for i, c := range s {
		a[i] = string(c)
	}
	return a
}

// Remove finds the element and deletes it from the slice.
func Remove[T comparable](s []T, elt T) ([]T, error) {
	i := Index(s, elt)

	if i == -1 {
		return nil, fmt.Errorf("value %v not found in %v", elt, s)
	}

	return DeleteElt(s, i), nil
}

// CommaSeparate converts an array of integers to strings and returns them separated by ", ".
func CommaSeparate(nums []int) string {
	s := make([]string, len(nums))
	for i, n := range nums {
		s[i] = strconv.Itoa(n)
	}

	return strings.Join(s, ", ")
}
