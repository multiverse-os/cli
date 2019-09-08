package os

import (
	"os"
	"path/filepath"
)

func ExecutableName() (name string) {
	executable, _ := os.Executable()
	_, name := filepath.Split(executable)
}

func ApplicationName() string {
	return ExecutableName()
}
