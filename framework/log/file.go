package log

import (
	"os"
	"os/user"
	"path"
	"strings"
)

func DefaultLogPath(name string) (logPath string) {
	if logPath, ok := FindOrCreateFile(DefaultOSLogPath(name)); !ok {
		if logPath, _ = FindOrCreateFile(DefaultUserLogPath(name)); !ok {
			workingPath, _ := os.Getwd()
			return (workingPath + name + ".log")
		} else {
			return logPath
		}
	} else {
		return logPath
	}
}

func FindOrCreateFile(filePath string) (string, bool) {
	filePath, filename := path.Split(filePath)
	if filename != "" {
		if _, err := os.Stat((filePath + filename)); os.IsNotExist(err) {
			os.MkdirAll(filePath, 0700)
			os.OpenFile((filePath + filename), os.O_RDONLY|os.O_CREATE, 0660)
			if _, err := os.Stat((filePath + filename)); !os.IsNotExist(err) {
				return (filePath + filename), true
			}
		} else {
			return (filePath + filename), true
		}
	}
	return "", false
}

func DefaultOSLogPath(name string) string {
	return ("/var/log/" + strings.ToLower(name) + ".log")
}

func DefaultUserLogPath(name string) string {
	home := os.Getenv("XDG_CONFIG_HOME")
	if home == "" {
		home = os.Getenv("HOME")
		if home == "" {
			currentUser, err := user.Current()
			if err != nil {
				Fatal(err.Error())
			}
			home = currentUser.HomeDir
		}
	}
	name = strings.ToLower(name)
	return home + ("/.local/share/" + name + "/" + name + ".log")
}
