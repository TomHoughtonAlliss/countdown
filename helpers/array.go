package helpers

import "slices"

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

func Remove[T comparable](s []T, elt T) []T {
	i := Index(s, elt)
	return DeleteElt(s, i)
}
