package style

import (
	"strconv"
)

const (
	prefix = "\x1b["
	suffix = "m"
)

const (
	reset = 0
	// ANSI Styles
	bold          = 1
	dim           = 2 // Decreased intensity
	italic        = 3 // Not widely support, sometimes inverse.
	underline     = 4
	slowBlink     = 5 // Less than 150 times per minute
	fastBlink     = 6 // Over 150 times per minute
	reverse       = 7 // Swaps fg and bg colors
	conceal       = 8 // Not widely supported
	strikethrough = 9 // Crossed-out
)

const (
	strong   = bold
	emphasis = italic // These help with those familiar with HTML
	thin     = dim
	inverse  = reverse
	blink    = slowBlink
	hide     = conceal
	crossOut = strikethrough
)

func Off(code int) int                  { return code + 20 }
func Sequence(code int) string          { return prefix + strconv.Itoa(code) + suffix }
func Text(code int, text string) string { return Sequence(code) + text + Sequence(Off(code)) }

// Style Text
///////////////////////////////////////////////////////////////////////////////
func Strong(text string) string        { return Text(strong, text) }
func Bold(text string) string          { return Text(bold, text) }
func Italic(text string) string        { return Text(italic, text) }
func Emphasis(text string) string      { return Text(emphasis, text) }
func Dim(text string) string           { return Text(dim, text) }
func Thin(text string) string          { return Text(thin, text) }
func Underline(text string) string     { return Text(underline, text) }
func Strikethrough(text string) string { return Text(strikethrough, text) }
func Crossout(text string) string      { return Text(crossOut, text) }
func Blink(text string) string         { return Text(blink, text) }
func SlowBlink(text string) string     { return Text(slowBlink, text) }
func FastBlink(text string) string     { return Text(fastBlink, text) }
func Inverse(text string) string       { return Text(inverse, text) }
func Reverse(text string) string       { return Text(reverse, text) }
func Conceal(text string) string       { return Text(conceal, text) }
func Hide(text string) string          { return Text(hide, text) }
