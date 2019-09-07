package style

const (
	escape = "\x1b"
	prefix = escape + "["
	suffix = "m"
)

const (
	Reset = 0
	// ANSI Styles
	Bold          = 1
	Faint         = 2 // Decreased intensity
	Italic        = 3 // Not widely support, sometimes inverse.
	Underline     = 4
	SlowBlink     = 5 // Less than 150 times per minute
	RapidBlink    = 6 // Over 150 times per minute
	Reverse       = 7 // Swaps fg and bg colors
	Conceal       = 8 // Not widely supported
	Strikethrough = 9 // Crossed-out
)

const (
	Normal   = Reset
	Strong   = Bold
	Emphasis = Italic // These help with those familiar with HTML
	Dim      = Faint
	Thin     = Faint
	Inverse  = Reverse
	Blink    = SlowBlink
	Hide     = Conceal
	CrossOut = Strikethrough
)

func Sequence(code int) string          { return prefix + strconv.Itoa(code) + suffix }
func Text(code int, text string) string { return Sequence(code) + text + Sequence(code+20) }

// Style Text
///////////////////////////////////////////////////////////////////////////////
func Strong(text string) string        { return Style(Strong, text) }
func Bold(text string) string          { return Style(Bold, text) }
func Italic(text string) string        { return Style(Italic, text) }
func Emphasis(text string) string      { return Style(Emphasis, text) }
func Faint(text string) string         { return Style(Faint, text) }
func Dim(text string) string           { return Style(Dim, text) }
func Thin(text string) string          { return Style(Thin, text) }
func Underline(text string) string     { return Style(Underline, text) }
func Strikethrough(text string) string { return Style(Strikethrough, text) }
func Crossout(text string) string      { return Style(CrossOut, text) }
func Blink(text string) string         { return Style(Blink, text) }
func SlowBlink(text string) string     { return Style(SlowBlink, text) }
func FastBlink(text string) string     { return Style(FastBlink, text) }
func Inverse(text string) string       { return Style(Inverse, text) }
func Reverse(text string) string       { return Style(Reverse, text) }
func Conceal(text string) string       { return Style(Conceal, text) }
func Hide(text string) string          { return Style(Hide, text) }
