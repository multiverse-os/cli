package ansi

// TODO: This hasn't been merged in yet, this overlaps with cursor.go

import (
	"bufio"
	"bytes"
	"fmt"
	"regexp"
	"strconv"
)

var COORDINATE_SYSTEM_BEGIN Short = 1

var dsrPattern = regexp.MustCompile(`\x1b\[(\d+);(\d+)R$`)

type Cursor struct {
	In  FileReader
	Out FileWriter
}

func (c *Cursor) Up(n int) {
	fmt.Fprintf(c.Out, "\x1b[%dA", n)
}

func (c *Cursor) Down(n int) {
	fmt.Fprintf(c.Out, "\x1b[%dB", n)
}

func (c *Cursor) Forward(n int) {
	fmt.Fprintf(c.Out, "\x1b[%dC", n)
}

func (c *Cursor) Back(n int) {
	fmt.Fprintf(c.Out, "\x1b[%dD", n)
}

func (c *Cursor) NextLine(n int) {
	fmt.Fprintf(c.Out, "\x1b[%dE", n)
}

func (c *Cursor) PreviousLine(n int) {
	fmt.Fprintf(c.Out, "\x1b[%dF", n)
}

func (c *Cursor) HorizontalAbsolute(x int) {
	fmt.Fprintf(c.Out, "\x1b[%dG", x)
}

func (c *Cursor) Show() {
	fmt.Fprint(c.Out, "\x1b[?25h")
}

func (c *Cursor) Hide() {
	fmt.Fprint(c.Out, "\x1b[?25l")
}

func (c *Cursor) Move(x int, y int) {
	fmt.Fprintf(c.Out, "\x1b[%d;%df", x, y)
}

func (c *Cursor) Save() {
	fmt.Fprint(c.Out, "\x1b7")
}

func (c *Cursor) Restore() {
	fmt.Fprint(c.Out, "\x1b8")
}

func (c *Cursor) MoveNextLine(cur *Coord, terminalSize *Coord) {
	if cur.Y == terminalSize.Y {
		fmt.Fprintln(c.Out)
	}
	c.NextLine(1)
}

func (c *Cursor) Location(buf *bytes.Buffer) (*Coord, error) {
	// ANSI escape sequence for DSR - Device Status Report
	// https://en.wikipedia.org/wiki/ANSI_escape_code#CSI_sequences
	fmt.Fprint(c.Out, "\x1b[6n")

	// There may be input in Stdin prior to CursorLocation so make sure we don't
	// drop those bytes.
	var loc []int
	var match string
	for loc == nil {
		// Reports the cursor position (CPR) to the application as (as though typed at
		// the keyboard) ESC[n;mR, where n is the row and m is the column.
		reader := bufio.NewReader(c.In)
		text, err := reader.ReadSlice('R')
		if err != nil {
			return nil, err
		}

		loc = dsrPattern.FindStringIndex(string(text))
		if loc == nil {
			// Stdin contains R that doesn't match DSR.
			buf.Write(text)
		} else {
			buf.Write(text[:loc[0]])
			match = string(text[loc[0]:loc[1]])
		}
	}

	matches := dsrPattern.FindStringSubmatch(string(match))
	if len(matches) != 3 {
		return nil, fmt.Errorf("incorrect number of matches: %d", len(matches))
	}

	col, err := strconv.Atoi(matches[2])
	if err != nil {
		return nil, err
	}

	row, err := strconv.Atoi(matches[1])
	if err != nil {
		return nil, err
	}

	return &Coord{Short(col), Short(row)}, nil
}

func (cur Coord) CursorIsAtLineEnd(size *Coord) bool {
	return cur.X == size.X
}

func (cur Coord) CursorIsAtLineBegin() bool {
	return cur.X == COORDINATE_SYSTEM_BEGIN
}

// Size returns the height and width of the terminal.
func (c *Cursor) Size(buf *bytes.Buffer) (*Coord, error) {
	// the general approach here is to move the cursor to the very bottom
	// of the terminal, ask for the current location and then move the
	// cursor back where we started

	// hide the cursor (so it doesn't blink when getting the size of the terminal)
	c.Hide()
	// save the current location of the cursor
	c.Save()

	// move the cursor to the very bottom of the terminal
	c.Move(999, 999)

	// ask for the current location
	bottom, err := c.Location(buf)
	if err != nil {
		return nil, err
	}

	// move back where we began
	c.Restore()

	// show the cursor
	c.Show()
	// sice the bottom was calcuated in the lower right corner, it
	// is the dimensions we are looking for
	return bottom, nil
}
