package gotools

import "strings"

// StartsWith checks if the string starts with any of the substrings in the list.
func StartsWith(s string, substrings []string) map[string]bool {
	result := make(map[string]bool)
	for _, substring := range substrings {
		if strings.HasPrefix(s, substring) {
			result[substring] = true
		} else {
			result[substring] = false
		}
	}
	return result
}
