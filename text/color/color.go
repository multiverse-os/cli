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
	purple          = prefix + "1;35" + suffix
	lightgray       = prefix + "37;0" + suffix
	black           = prefix + "0;30" + suffix
	red             = prefix + "0;31" + suffix
	green           = prefix + "0;32" + suffix
	yellow          = prefix + "0;33" + suffix
	cyan            = prefix + "0;36" + suffix
	blue            = prefix + "0;34" + suffix
	white           = prefix + "0;37" + suffix
	strikethrough   = prefix + "0;9" + suffix
	hidden          = prefix + "0;8" + suffix
	whiteBackground = prefix + "0;7" + suffix
	blink           = prefix + "0;5" + suffix
	underline       = prefix + "0;4" + suffix
	strong          = prefix + "0;1" + suffix
	italic          = prefix + "0;3" + suffix
	reset           = prefix + "0;" + suffix
)

const (
	Header    = strong
	Subheader = white
	Emphasis  = italic
	Strong    = strong
	Success   = green
	Warning   = yellow
	Fail      = red
	Reset     = reset
)

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
