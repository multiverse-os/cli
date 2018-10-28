package fs

import (
	"os"
	"path/filepath"
)

// Our library is designed for usage within core Multiverse OS software, and so
// only Linux (POSIX) systems are relevant. Some functions from the standard Go
// library will be reimplemented to remove Plan9/Windows and other excess to
// provide similar functionality with less codebase overhead.
type Path string

// WorkingDirectory returns path of executable regardless of what folder
// execution occurred from which is more consistent.
func WorkingDirectory() (string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}

func PWD() string {
	return WorkingDirectory()
}
