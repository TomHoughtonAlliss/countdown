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
