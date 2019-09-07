package ansi

import (
	"strconv"
)

const (
	escape = "\x1b"
	prefix = escape + "["
	suffix = "m"
	reset  = 0
)

func Sequence(code int) string { return prefix + strconv.Itoa(code) + suffix }
func Reset(text string) string { return Sequence(reset) + text }
