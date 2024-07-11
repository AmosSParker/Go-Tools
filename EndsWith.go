package gotools

import "strings"

// EndsWith checks if the string ends with any of the substrings in the list.
func EndsWith(s string, substrings []string) map[string]bool {
	result := make(map[string]bool)
	for _, substring := range substrings {
		if strings.HasSuffix(s, substring) {
			result[substring] = true
		} else {
			result[substring] = false
		}
	}
	return result
}
