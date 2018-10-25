package log

import (
	"os"
	"os/user"
	"strings"
)

func FindOrCreateFile(logFilePath string) bool {
	if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
		os.MkdirAll(logFilePath, 0660)
		os.OpenFile(logFilePath, os.O_RDONLY|os.O_CREATE, 0660)
		if _, err := os.Stat(logFilePath); os.IsNotExist(err) {
			return false
		} else {
			return true
		}
	} else {
		return true
	}
}

func OSLogPath(appName string) string {
	return ("/var/log/" + strings.ToLower(appName) + ".log")
}

func UserLogPath(appName string) string {
	appName = strings.ToLower(appName)
	home := os.Getenv("XDG_CONFIG_HOME")
	if home != "" {
		return (home + "/.local/share/" + appName + "/")
	} else {
		home = os.Getenv("HOME")
		if home != "" {
			return (home + "/.local/share/" + appName + "/")
		} else {
			currentUser, err := user.Current()
			if err != nil {
				FatalError(err)
			}
			home = currentUser.HomeDir
			return (home + "/.local/share/" + appName + "/")
		}
	}
}
