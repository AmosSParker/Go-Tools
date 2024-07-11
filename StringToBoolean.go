package gotools

import "strconv"

// StringToBoolean converts a string to a boolean.
func StringToBoolean(s string) bool {
	b, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}
	return b
}
