package cli

import (
	"os"
	"path/filepath"
)

type process struct {
	ID         int
	CWD        string
	Executable string
}

func Process() process {
	cwd, executable := filepath.Split(os.Args[0])
	return process{
		ID:         os.Getpid(),
		CWD:        cwd,
		Executable: executable,
	}
}
