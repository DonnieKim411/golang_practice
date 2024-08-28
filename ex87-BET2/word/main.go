// Package word provides custom functions for working with string.

package word

import "strings"

// UserCount returns the number of times each word is used in a string
// no need to write an example for this one
// writing a test for this one is a bonus challenge; harder
func UseCount(s string) map[string]int {
	xs := strings.Fields(s)
	m := make(map[string]int)
	for _, v := range xs {
		m[v]++
	}
	return m
}

// count number of words in string
func Count(s string) int {
	xs := strings.Fields(s)
	return len(xs)
}
