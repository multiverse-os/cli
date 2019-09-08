package vt100_test

import (
	"io"
	"strings"
	"testing"

	. "github.com/jaguilar/vt100"
	"github.com/jaguilar/vt100/vttest"
	"github.com/stretchr/testify/assert"
)

func splitLines(s string) [][]rune {
	ss := strings.Split(s, "\n")
	r := make([][]rune, len(ss))
	for i, line := range ss {
		r[i] = []rune(line)
	}
	return r
}

func esc(s string) string {
	return "\u001b" + s
}

func cmd(s string) Command {
	cmd, err := Decode(strings.NewReader(s))
	if err != nil {
		panic(err)
	}
	return cmd
}

func cmds(s string) []Command {
	var c []Command
	r := strings.NewReader(s)
	for {
		x, err := Decode(r)
		if err == io.EOF {
			return c
		}
		if err != nil {
			panic(err)
		}
		c = append(c, x)
	}
}

func TestPutRune(t *testing.T) {
	v := vttest.FromLines("abc\ndef\nghi")
	v.Cursor.Y = 1
	v.Cursor.X = 1

	assert.Nil(t, v.Process(cmd("z")))
	assert.Equal(t, splitLines("abc\ndzf\nghi"), v.Content)
	assert.Equal(t, 2, v.Cursor.X)
	assert.Equal(t, 1, v.Cursor.Y)
}

func TestMoveCursor(t *testing.T) {
	v := vttest.FromLines("abc\ndef\nghi")
	assert.Nil(t, v.Process(cmd(esc("[3;1H"))))
	assert.Equal(t, 2, v.Cursor.Y)
	assert.Equal(t, 0, v.Cursor.X)
}

func TestCursorDirections(t *testing.T) {
	v := vttest.FromLines("abc\ndef\nghi")

	moves := strings.Join([]string{
		esc("[2B"), // down down
		esc("[2C"), // right right
		esc("[A"),  // up (no args = 1)
		esc("[1D"), // left
	}, "") // End state: +1, +1
	s := strings.NewReader(moves)

	want := []Cursor{
		{Y: 2, X: 0},
		{Y: 2, X: 2},
		{Y: 1, X: 2},
		{Y: 1, X: 1},
	}
	var got []Cursor

	cmd, err := Decode(s)
	for err == nil {
		assert.Nil(t, v.Process(cmd))
		got = append(got, v.Cursor)
		cmd, err = Decode(s)
	}
	if assert.Equal(t, err, io.EOF) {
		assert.Equal(t, want, got)
	}
}

func TestErase(t *testing.T) {
	c := Format{Fg: Yellow, Intensity: Bright}
	var d Format
	for _, tc := range []struct {
		command Command
		want    *VT100
	}{
		{cmd(esc("[K")), vttest.FromLinesAndFormats("abcd\nef  \nijkl", [][]Format{
			{c, c, c, c},
			{c, c, d, d},
			{c, c, c, c},
		})},
		{cmd(esc("[1K")), vttest.FromLinesAndFormats("abcd\n   h\nijkl", [][]Format{
			{c, c, c, c},
			{d, d, d, c},
			{c, c, c, c},
		})},
		{cmd(esc("[2K")), vttest.FromLinesAndFormats("abcd\n    \nijkl", [][]Format{
			{c, c, c, c},
			{d, d, d, d},
			{c, c, c, c},
		})},
		{cmd(esc("[J")), vttest.FromLinesAndFormats("abcd\n    \n    ", [][]Format{
			{c, c, c, c},
			{d, d, d, d},
			{d, d, d, d},
		})},
		{cmd(esc("[1J")), vttest.FromLinesAndFormats("    \n    \nijkl", [][]Format{
			{d, d, d, d},
			{d, d, d, d},
			{c, c, c, c},
		})},
		{cmd(esc("[2J")), vttest.FromLinesAndFormats("    \n    \n    ", [][]Format{
			{d, d, d, d},
			{d, d, d, d},
			{d, d, d, d},
		})},
	} {
		v := vttest.FromLinesAndFormats(
			"abcd\nefgh\nijkl", [][]Format{
				{c, c, c, c},
				{c, c, c, c},
				{c, c, c, c},
			})
		v.Cursor = Cursor{Y: 1, X: 2}
		beforeCursor := v.Cursor

		assert.Nil(t, v.Process(tc.command))
		assert.Equal(t, tc.want.Content, v.Content, "while evaluating ", tc.command)
		assert.Equal(t, tc.want.Format, v.Format, "while evaluating ", tc.command)
		// Check the cursor separately. We don't set it on any of the test cases
		// so we cannot expect it to be equal. It's not meant to change.
		assert.Equal(t, beforeCursor, v.Cursor)
	}
}

var (
	bs = "\u0008" // Use strings to contain these runes so they can be concatenated easily.
	lf = "\u000a"
	cr = "\u000d"
)

func TestBackspace(t *testing.T) {
	v := vttest.FromLines("BA..")
	v.Cursor.Y, v.Cursor.X = 0, 2

	backspace := cmd(bs)
	assert.Nil(t, v.Process(backspace))
	// Backspace doesn't actually delete text.
	assert.Equal(t, vttest.FromLines("BA..").Content, v.Content)
	assert.Equal(t, 1, v.Cursor.X)

	v.Cursor.X = 0
	assert.Nil(t, v.Process(backspace))
	assert.Equal(t, 0, v.Cursor.X)

	v = vttest.FromLines("..\n..")
	v.Cursor.Y, v.Cursor.X = 1, 0
	assert.Nil(t, v.Process(backspace))
	assert.Equal(t, 0, v.Cursor.Y)
	assert.Equal(t, 1, v.Cursor.X)
}

func TestLineFeed(t *testing.T) {
	v := vttest.FromLines("AA\n..")
	v.Cursor.X = 1

	for _, c := range cmds(lf + "b") {
		assert.Nil(t, v.Process(c))
	}
	assert.Equal(t, vttest.FromLines("AA\nb.").Content, v.Content)
}

func TestCarriageReturn(t *testing.T) {
	v := vttest.FromLines("AA\n..")
	v.Cursor.X = 1

	for _, c := range cmds(cr + "b") {
		assert.Nil(t, v.Process(c))
	}
	assert.Equal(t, vttest.FromLines("bA\n..").Content, v.Content)
}

func TestAttributes(t *testing.T) {
	v := vttest.FromLines("....")
	s := strings.NewReader(
		esc("[2ma") + esc("[5;22;31mb") + esc("[0mc") + esc("[4;46md"))
	cmd, err := Decode(s)
	for err == nil {
		assert.Nil(t, v.Process(cmd))
		cmd, err = Decode(s)
	}
	assert.Equal(t, io.EOF, err)
	assert.Equal(t, []rune("abcd"), v.Content[0])
	assert.Equal(t, []Format{
		{Intensity: Dim}, {Blink: true, Fg: Red}, {}, {Underscore: true, Bg: Cyan},
	}, v.Format[0])
}
