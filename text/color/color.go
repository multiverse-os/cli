package color

// CLIHelpTemplate is the text template for the Default help topic.
// cli.go uses text/template to render templates. You can
// render custom help text by setting this variable.

const (
	prefix = "\x1b["
	suffix = "m"
)

const (
	//blue   = prefix + "34;1" + suffix
	reset           = prefix + "0;" + suffix
	strong          = prefix + "0;1" + suffix
	light           = prefix + "0;2" + suffix
	italic          = prefix + "0;3" + suffix
	underline       = prefix + "0;4" + suffix
	blink           = prefix + "0;5" + suffix
	hidden          = prefix + "0;8" + suffix
	strikethrough   = prefix + "0;9" + suffix
	test            = prefix + "0;11" + suffix
	purple          = prefix + "1;35" + suffix
	lightgray       = prefix + "37;0" + suffix
	black           = prefix + "0;30" + suffix
	red             = prefix + "0;31" + suffix
	green           = prefix + "0;32" + suffix
	yellow          = prefix + "0;33" + suffix
	cyan            = prefix + "0;36" + suffix
	blue            = prefix + "0;34" + suffix
	white           = prefix + "0;37" + suffix
	whiteBackground = prefix + "0;7" + suffix
)

const (
	HeaderCode    = strong
	SubheaderCode = white
	EmphasisCode  = italic
	StrongCode    = strong
	LightCode     = light
	SuccessCode   = green
	WarningCode   = yellow
	FailCode      = red
	InfoCode      = blue
)

// Reset aliasing
const (
	ResetCode = reset
	Reset     = reset
)

//
// Standardized Theme Color Functions
///////////////////////////////////////////////////////////////////////////////

func Default(text string) string {
	return (Reset + text)
}

func Header(text string) string {
	return (HeaderCode + text + reset)
}

func Subheader(text string) string {
	return (SubheaderCode + text + reset)
}

func Emphasis(text string) string {
	return (EmphasisCode + text + reset)
}

func Strong(text string) string {
	return (StrongCode + text + reset)
}

func Light(text string) string {
	return (LightCode + text + reset)
}

func Success(text string) string {
	return (SuccessCode + text + reset)
}

func Warning(text string) string {
	return (WarningCode + text + reset)
}

func Fail(text string) string {
	return (FailCode + text + reset)
}

func Info(text string) string {
	return (InfoCode + text + reset)
}

//

// Standard Color Functions
///////////////////////////////////////////////////////////////////////////////

func Purple(text string) string {
	return (purple + text + reset)
}

func Cyan(text string) string {
	return (cyan + text + reset)
}

func Blue(text string) string {
	return (blue + text + reset)
}

func White(text string) string {
	return (white + text + reset)
}

func Green(text string) string {
	return (green + text + reset)
}

func Yellow(text string) string {
	return (yellow + text + reset)
}

func Black(text string) string {
	return (black + text + reset)
}

func Red(text string) string {
	return (red + text + reset)
}
