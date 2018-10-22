package terminal

import (
	"io"
)

func NewAnsiStdout(out FileWriter) io.Writer {
	return out
}

func NewAnsiStderr(out FileWriter) io.Writer {
	return out
}
