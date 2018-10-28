package binary

import (
	"os"
	"path/filepath"
)

type Binary struct {
	// TODO: Use an embedded struct `File` from fs/file.go that can be used to
	// provide basic functionality for any file.
	Name        string
	Path        string
	Permissions int
	Owner       string
	Group       string
	Data        []byte
}

func ExecutableName() (name string) {
	executable, _ := os.Executable()
	_, name := filepath.Split(executable)
}

// Aliasing
func ApplicationName() string { return ExecutableName() }
func BinaryName() string      { return ExecutableName() }
