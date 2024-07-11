package gotools

import (
	"strings"
)

// removes all specified delimiters from the input string
// `delimiters“ is a slice of strings to be removed
// `n` specifies the number of replacements (-1 for all of them)
// returns the cleaned string and an error if any
func EraseDelimiter(input string, delimiters []string, n int) (string, error) {
	result := input
	for _, delimiter := range delimiters {
		result = strings.Replace(result, delimiter, "", n)
	}
	return result, nil
}
