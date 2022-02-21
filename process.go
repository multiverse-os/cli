package cli

import (
	"path/filepath"
  "os"
)

type process struct {
  ID          int
  CWD         string
  Executable  string
}

func Process() process {
	cwd, executable := filepath.Split(arguments[0])
  return Process{
    ID: os.Getpid(),
    CWD: cwd,
    Executable: executable,
  }
}
