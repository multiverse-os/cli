package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// LoadHostInfo loads "~/.config/bashprompt/host_info", returning box data if it
// exists and can be loaded or nil otherwise.
//
// A valid host_info file will be interpreted to build up a coloured hostname
// box section, allowing per-host themes. File definition is as follows:
//  • one definition per line
//  • anything after '#' is treated as a comment
//  • empty lines are ignored
//  • remaining lines must have exactly two words:
//    · first word is printed verbatim
//    · second word is “[esc]” (escape codes), brackets required
func LoadHostInfo() []byte {
	fname := filepath.Join(ConfigDir, "host_info")
	fp, err := os.Open(fname)
	if err != nil {
		return nil
	}
	defer fp.Close()

	var (
		lineNum int
		buf     bytes.Buffer
		state   = hostInfoState{
			first:  true,
			leftBg: '7',
		}
	)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		lineNum++
		if err := hostInfoLine(scanner.Bytes(), &buf, &state); err != nil {
			err = fmt.Errorf("%s:%d: %v", fname, lineNum, err)
			CaptureError(err)
			return nil
		}
	}
	if err := scanner.Err(); err != nil {
		CaptureError(err)
		return nil
	}

	return buf.Bytes()
}

type hostInfoState struct {
	first   bool
	leftBg  byte
	rightBg byte
}

func hostInfoLine(line []byte, buf *bytes.Buffer, state *hostInfoState) error {
	// strip comments
	if pos := bytes.IndexRune(line, '#'); pos != -1 {
		line = line[:pos]
	}

	// split into whitespace-separated fields
	fi := bytes.Fields(line)
	switch len(fi) {
	case 0:
		// skip blank lines
		return nil

	case 2:
		// expected

	default:
		return errors.New("line must contain exactly two tokens")
	}

	l1 := len(fi[1])
	if l1 <= 2 || fi[1][0] != '[' || fi[1][l1-1] != ']' {
		return errors.New("malformed escape code")
	}

	colour := fi[1][1 : l1-1]
	colours := bytes.Split(colour, []byte(";"))
	for _, c := range colours {
		if len(c) == 2 && c[0] == '4' {
			state.rightBg = c[1]
		}
	}

	if state.first {
		state.first = false
	} else {
		// TODO: colours
		SetColour(buf, fmt.Sprintf("0;3%c;4%c", state.leftBg, state.rightBg))
		buf.WriteRune('') // \ue0b0 vim powerline separator
	}

	state.leftBg = state.rightBg

	SetColour(buf, string(fi[1][1:l1-1]))
	buf.Write(fi[0])
	return nil
}
