package gotools

import "os/user"

// GrabUsername returns the current user's username.
func GrabUsername() string {
	user, err := user.Current()
	if err != nil {
		return ""
	}
	return user.Username
}
