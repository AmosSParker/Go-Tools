package gotools

import "os/user"

// GrabHomeDir returns the current user's home directory.
func GrabHomeDir() string {
	user, err := user.Current()
	if err != nil {
		return ""
	}
	return user.HomeDir
}
