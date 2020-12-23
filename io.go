package cli

import (
	"io"
)

type Stdio struct {
	In  FileReader
	Out FileWriter
	Err io.Writer
}

type FileWriter interface {
	io.Writer
	Fd() uintptr
}

type FileReader interface {
	io.Reader
	Fd() uintptr
}
