package color

// CLIHelpTemplate is the text template for the Default help topic.
// cli.go uses text/template to render templates. You can
// render custom help text by setting this variable.
const (
	lightPurple = "\x1b[1;35m"
	blue        = "\x1b[34;1m"
	lightGray   = "\x1b[37;0m"
	white       = "\x1b[0;1m"
	reset       = "\x1b[0;m"
)

const (
	titleColor   = "\x1b[1;35m"
	commandColor = "\x1b[34;1m"
	versionColor = "\x1b[37;0m"
	headerColor  = "\x1b[0;1m"
	resetColor   = "\x1b[0;m"
)

func repeat(c string, times int) string {
	aggregate := ""
	for i := 1; i <= times; i++ {
		aggregate += c
	}
	return aggregate
}
