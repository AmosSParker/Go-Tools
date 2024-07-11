package gotools

import "os"

// GrabExecutablePath returns the path of the current executable
func GrabExecutablePath() string {
	ex, err := os.Executable()
	if err != nil {
		return ""
	}
	return ex
}
