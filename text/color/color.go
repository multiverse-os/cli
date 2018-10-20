package color

// CLIHelpTemplate is the text template for the Default help topic.
// cli.go uses text/template to render templates. You can
// render custom help text by setting this variable.

const (
	prefix = "\x1b["
	suffix = "m"
)

const (
	lightPurple = prefix + "1;35" + suffix
	blue        = prefix + "34;1" + suffix
	lightGray   = prefix + "37;0" + suffix
	white       = prefix + "0;1" + suffix
	orange      = prefix + "" + suffix
	red         = prefix + "" + suffix
	green       = prefix + "" + suffix
	strong      = prefix + "" + suffix
	italics     = prefix + "" + suffix
)

const (
	Header    = white
	Subheader = lightPurple
	Command   = lightGray
	Success   = green
	Warning   = orange
	Fail      = red
	Version   = lightGray
	Reset     = prefix + "0;" + suffix
)
