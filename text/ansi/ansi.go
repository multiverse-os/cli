package ansi

import (
	"strconv"

	color "github.com/multiverse-os/text/ansi/color"
	style "github.com/multiverse-os/text/ansi/style"
)

const (
	escape = "\x1b"
	prefix = escape + "["
	suffix = "m"
	reset  = 0
)

func Sequence(code int) string { return prefix + strconv.Itoa(code) + suffix }
func Reset(text string) string { return Sequence(reset) + text }

func ColorOpen(code int) string          { return color.Open(code) }
func ColorClose(code int) string         { return color.Close(code) }
func Color(code int, text string) string { return color.Open(code) + text + color.Close(code) }

func StyleOpen(code int) string          { return style.Open(code) }
func StyleClose(code int) string         { return style.Close(code) }
func Style(code int, text string) string { return style.Open(code) + text + style.Close(code) }
