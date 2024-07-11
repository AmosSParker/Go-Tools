package gotools

import (
	"os"
	"path/filepath"
)

// GrabExutableName returns the name of the current executable
func GrabExecutableName() string {
	ex, err := os.Executable()
	if err != nil {
		return ""
	}
	return filepath.Base(ex)
}
