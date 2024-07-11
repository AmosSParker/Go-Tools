package gotools

import "strconv"

// StringToInt converts a string to an integer.
func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return i
}
