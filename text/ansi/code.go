package main

import (
	"github.com/multiverse-os/ansi/color"
	"github.com/multiverse-os/ansi/style"
)

func Reset(text string) string          { return Sequence(Reset) + text }
func DefaultColor(text string) string   { return Sequence(color.Default) + text }
func DefaultBgColor(text string) string { return Sequence(color.DefaultBg) + text }

func Style(code int, text string) string { return style.Text(code, text) }
func Color(code int, text string) string { return color.Text(code, text) }

func Format(code int, text string) string {
	if code >= 1 && code <= 9 {
		return style.Text(code, text)
	} else {
		return color.Text(code, text)
	}
}
