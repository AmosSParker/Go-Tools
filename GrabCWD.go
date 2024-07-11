package gotools

import "os"

// GrabCWD returns the current working directory.
func GrabCWD() string {
	cwd, err := os.Getwd()
	if err != nil {
		return ""
	}
	return cwd
}
