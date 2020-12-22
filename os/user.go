package os

import (
	"os"
	"os/user"
)

// TODO: Need to add config path, local data path,
type Group struct {
	Name  string
	Users []*User
	ID    int
}

// TODO: Automatically pull in SSH keys; possibly create a ECDSA ssh key by
// default. Use these for signing, encryption, etc by default
type User struct {
	Name   string
	ID     string
	Groups []*Group

	//Root bool
	//Sudo bool

	Data   Path
	Config Path
	Temp   Path
}

func CurrentUsername() string {
	// os.User in the stdlibrary is defined as:
	// User{
	// 	Uid string // NOTE: Should be int, but for Windows they made it string
	//  Gid string
	//  Username string // TODO: Should have functions to check validity
	//  Name string // Often the display name vs the login name (username)
	//  HomeDir string // NOTE: This really should be HomePath
	// }
	//
	user, err := user.Current()
	if err != nil {
		// TODO: Currently assuming this should essentially never occur, but in
		//       the event it does occur, the details should be recorded in the
		//       the error. This will enable detailed errors to assist the user,
		//       (and eventually the developers) in automatically resolving and
		//       avoiding all error conditions.
		panic(err)
	}
	return username
}

func CurrentUser() *User {
	return &User{
		Name: username,
	}
}

func Home() string {
	home := os.Getenv("XDG_CONFIG_HOME")
	if len(home) != 0 {
		return home
	} else {
		home = os.Getenv("HOME")
		if len(home) != 0 {
			return home
		} else {
			currentUser, err := user.Current()
			home = currentUser.HomeDir
			return home
		}
	}
}
