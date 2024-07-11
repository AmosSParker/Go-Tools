package gotools

import "strings"

// Contains checks if the string contains any of the substrings in the list.
func Contains(s string, substrings []string) map[string]bool {
	result := make(map[string]bool)
	for _, substring := range substrings {
		if strings.Contains(s, substring) {
			result[substring] = true
		} else {
			result[substring] = false
		}
	}
	return result
}
