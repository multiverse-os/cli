package ansi

import (
	"strconv"
)

// TODO: What are the codes for these?
func ClearDisplay(code int) string { return prefix + strconv.Itoa(code) + "2J" }

// EraseLine clears part of the line.
// If `n` is zero, clears the cursor to the end of the line.
// If `n` is one, clear from cursor to beginning of the line.
// If `n` is two, clear entire line.
// Cursor position does not change
func EraseLine(code int) string               { return prefix + strconv.Itoa(code) + "K" }
func SelectGraphicsRendition(code int) string { return prefix + strconv.Itoa(code) + "m" }
