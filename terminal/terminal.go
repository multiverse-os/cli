package main

import (
	"fmt"
)

const ioctlReadTermios = syscall.TCGETS
const ioctlWriteTermios = syscall.TCSETS

const (
	BashNPStart = "\x01"
	BashNPEnd   = "\x02"
	CSI         = "\x1B["
)

type TerminalWriter interface {
	Write(b []byte) (int, error)
	WriteRune(r rune) (int, error)
	WriteString(s string) (int, error)
}

func SetColour(w TerminalWriter, raw string) {
	w.Write([]byte(BashNPStart + CSI))
	w.WriteString(raw)
	w.Write([]byte("m" + BashNPEnd))
}

type RGB struct {
	R, G, B int
}

var RGBUnset = RGB{-1, -1, -1}

func SetRGB(w TerminalWriter, fg, bg RGB) {
	if fg != RGBUnset {
		w.Write([]byte(BashNPStart + CSI))
		fmt.Fprintf(w, "38;2;%d;%d;%d", fg.R, fg.G, fg.B)
		w.Write([]byte("m" + BashNPEnd))
	}
	if bg != RGBUnset {
		w.Write([]byte(BashNPStart + CSI))
		fmt.Fprintf(w, "48;2;%d;%d;%d", bg.R, bg.G, bg.B)
		w.Write([]byte("m" + BashNPEnd))
	}
}

// PrintableLength returns the length of the printable characters in the given
// string. It will strip out anything between ASCII control chars 1 (SOH) and 2
// (STX), as per readline.
func PrintableLength(str string) int {
	var (
		length   int
		inEscape bool
	)
	for _, r := range str {
		switch {
		case inEscape:
			if r == '\x02' {
				inEscape = false
			}
		case r == '\x01':
			inEscape = true
		default:
			length++
		}
	}
	return length
}
