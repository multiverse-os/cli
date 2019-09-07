package style

import (
	"strconv"
)

const (
	prefix = "\x1b["
	suffix = "m"
)

const (
	RESET = 0
	// ANSI Styles
	BOLD             = 1
	DIM              = 2 // Decreased intensity
	ITALIC           = 3 // Not widely support, sometimes inverse.
	UNDERLINE        = 4
	SLOW_BLINK       = 5 // Less than 150 times per minute
	FAST_BLINK       = 6 // Over 150 times per minute
	INVERSE          = 7 // Swaps fg and bg colors
	CONCEAL          = 8 // Not widely supported
	STRIKETHROUGH    = 9 // Crossed-out
	DEFAULT_FONT     = 10
	DOUBLE_UNDERLINE = 21
	FRAMED           = 51
	ENCIRCLE         = 52
	OVERLINE         = 53
)

// Aliasing
const (
	STRONG      = BOLD
	EMPHASIS    = ITALIC
	THIN        = DIM
	CROSSOUT    = STRIKETHROUGH
	BLINK       = SLOW_BLINK
	RAPID_BLINK = FAST_BLINK
	REVERSE     = INVERSE
	HIDE        = CONCEAL
)

func Off(code int) int {
	switch {
	case code == 1:
		return 22
	case (code >= 2 && code <= 10):
		return code + 20
	default:
		return 0
	}
}

func Sequence(code int) string          { return prefix + strconv.Itoa(code) + suffix }
func Reset() string                     { return Sequence(reset) }
func Open(code int) string              { return Sequence(code) }
func Close(code int) string             { return Sequence(Off(code)) }
func Text(code int, text string) string { return Sequence(code) + text + Sequence(Off(code)) }

// Style Text
///////////////////////////////////////////////////////////////////////////////
func Strong(text string) string        { return Text(BOLD, text) }
func Bold(text string) string          { return Text(BOLD, text) }
func Italic(text string) string        { return Text(ITALIC, text) }
func Emphasis(text string) string      { return Text(ITALIC, text) }
func Dim(text string) string           { return Text(DIM, text) }
func Thin(text string) string          { return Text(DIM, text) }
func Underline(text string) string     { return Text(UNDERLINE, text) }
func Strikethrough(text string) string { return Text(STRIKETHROUGH, text) }
func Crossout(text string) string      { return Text(STRIKETHROUGH, text) }
func Blink(text string) string         { return Text(SLOW_BLINK, text) }
func SlowBlink(text string) string     { return Text(SLOW_BLINK, text) }
func FastBlink(text string) string     { return Text(FAST_BLINK, text) }
func RapidBlink(text string) string    { return Text(FAST_BLINK, text) }
func Inverse(text string) string       { return Text(INVERSE, text) }
func Reverse(text string) string       { return Text(INVERSE, text) }
func Conceal(text string) string       { return Text(CONCEAL, text) }
func Hide(text string) string          { return Text(CONCEAL, text) }
func Framed(text string) string        { return Text(FRAMED, text) }
func Encircle(text string) string      { return Text(ENCIRCLE, text) }
func Overline(text string) string      { return Text(OVERLINE, text) }
