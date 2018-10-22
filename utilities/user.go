package main

import (
	"os"
	"os/user"
)

func CurrentUser() user.User {
	currentUser, err := user.Current()
	if err != nil {
		return currentUser
	} else {
		return nil
	}
}

func HomeFolder() string {
	home := os.Getenv("XDG_CONFIG_HOME")
	if home != "" {
		return home
	} else {
		home = os.Getenv("HOME")
		if home != "" {
			return home
		} else {
			currentUser, err := user.Current()
			home = currentUser.HomeDir
			return home
		}
	}
}
