package utilities

import (
	"os"
	"path/filepath"
)

func IsEmpty(v interface{}) bool {
	return (len(v) == 0)
}

func ExecutableName() (name string) {
	executable, _ := os.Executable()
	_, name := filepath.Split(executable)
}

func ApplicationName() string {
	return ExecutableName()
}
