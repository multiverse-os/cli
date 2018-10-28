package text

// CLIHelpTemplate is the text template for the Default help topic.
// cli.go uses text/template to render templates. You can
// render custom help text by setting this variable.

const (
	prefix = "\x1b["
	suffix = "m"
)

const (
	// Text Style Codes
	RESET         = prefix + "0" + suffix
	STRONG        = prefix + "1" + suffix
	LIGHT         = prefix + "2" + suffix
	ITALIC        = prefix + "3" + suffix
	UNDERLINE     = prefix + "4" + suffix
	BLINK         = prefix + "5" + suffix
	HIDDEN        = prefix + "8" + suffix
	STRIKETHROUGH = prefix + "9" + suffix
	DEFAULT       = prefix + "10" + suffix
	/////////////////////////////////////////////////////////////////////////////
	// 3/4 Bit Color Codes (Reference: en.wikipedia.org/wiki/ANSI_escape_code)
	// Foreground (3/4 Bit) Color Codes
	BLACK   = prefix + "30" + suffix
	RED     = prefix + "31" + suffix
	GREEN   = prefix + "32" + suffix
	YELLOW  = prefix + "33" + suffix // Or orange or brown; software dependent
	BLUE    = prefix + "34" + suffix
	MAGENTA = prefix + "35" + suffix
	CYAN    = prefix + "36" + suffix
	GRAY    = prefix + "37" + suffix // Or light-gray; software dependent
	// "Bright" Foreground (3/4 Bit) Color Codes
	DARK_GRAY     = prefix + "90" + suffix // Officially "Bright Black"
	LIGHT_RED     = prefix + "91" + suffix
	LIGHT_GREEN   = prefix + "92" + suffix
	LIGHT_YELLOW  = prefix + "93" + suffix
	LIGHT_BLUE    = prefix + "94" + suffix
	LIGHT_MAGENTA = prefix + "95" + suffix
	LIGHT_CYAN    = prefix + "96" + suffix
	WHITE         = prefix + "97" + suffix // Officially "Bright White"
)

// Aliasing
const (
	HEADER       = WHITE
	SUBHEADER    = STRONG
	EMPHASIS     = ITALIC
	BOLD         = STRONG
	THIN         = LIGHT
	SUCCESS      = GREEN
	LOG          = GRAY
	DEBUG        = GRAY
	INFO         = BLUE
	WARNING      = YELLOW
	WARN         = WARNING
	FAIL         = RED
	ERROR        = RED
	FATAL_ERROR  = RED
	FATAL        = RED
	PANIC        = RED
	TRACE        = BLUE
	H1           = HEADER
	H2           = SUBHEADER
	PURPLE       = MAGENTA
	LIGHT_PURPLE = LIGHT_MAGENTA
	LIGHT_BLACK  = DARK_GRAY
	LIGHT_WHITE  = WHITE
)

//
// Standardized Theme Color Functions
///////////////////////////////////////////////////////////////////////////////

func RandomColors() []string {
	var valueColors = []string{GREEN, BLUE, YELLOW, RED, PURPLE, CYAN, LIGHT_RED}
	return valueColors
}

func ANSI(style, text string) string {
	return (style + text + RESET)
}

func Default(text string) string {
	return (RESET + text)
}

func Header(text string) string {
	return (HEADER + text + RESET)
}

func Subheader(text string) string {
	return (SUBHEADER + text + RESET)
}

func Emphasis(text string) string {
	return (EMPHASIS + text + RESET)
}

// TODO: Add Weight(fontWeight int, text string) string {} func like with CSS
func Strong(text string) string {
	return (STRONG + text + RESET)
}

func Bold(text string) string {
	return (STRONG + text + RESET)
}

func Light(text string) string {
	return (LIGHT + text + RESET)
}

func Thin(text string) string {
	return (LIGHT + text + RESET)
}

func Success(text string) string {
	return (SUCCESS + text + RESET)
}

func Log(text string) string {
	return (LOG + text + RESET)
}

func Notice(text string) string {
	return (LOG + text + RESET)
}

func Info(text string) string {
	return (INFO + text + RESET)
}

func Debug(text string) string {
	return (LOG + text + RESET)
}

func Warning(text string) string {
	return (WARNING + text + RESET)
}

func Warn(text string) string {
	return (WARN + text + RESET)
}

func Trace(text string) string {
	return (TRACE + text + RESET)
}

func Fail(text string) string {
	return (FAIL + text + RESET)
}

func Error(text string) string {
	return (FAIL + text + RESET)
}

func Fatal(text string) string {
	return (FAIL + text + RESET)
}

func FatalError(text string) string {
	return (FAIL + text + RESET)
}

func Panic(text string) string {
	return (FAIL + text + RESET)
}

//
// Standard Color Functions
///////////////////////////////////////////////////////////////////////////////
func Black(text string) string {
	return (BLACK + text + RESET)
}

func Red(text string) string {
	return (RED + text + RESET)
}

func Green(text string) string {
	return (GREEN + text + RESET)
}

func Yellow(text string) string {
	return (YELLOW + text + RESET)
}

func Blue(text string) string {
	return (BLUE + text + RESET)
}

func Magenta(text string) string {
	return (MAGENTA + text + RESET)
}

func Purple(text string) string {
	return (PURPLE + text + RESET)
}

func Cyan(text string) string {
	return (CYAN + text + RESET)
}

func Gray(text string) string {
	return (GRAY + text + RESET)
}

func DarkGray(text string) string {
	return (DARK_GRAY + text + RESET)
}

func LightRed(text string) string {
	return (LIGHT_RED + text + RESET)
}

func LightGreen(text string) string {
	return (LIGHT_GREEN + text + RESET)
}

func LightYellow(text string) string {
	return (LIGHT_YELLOW + text + RESET)
}

func LightBlue(text string) string {
	return (LIGHT_BLUE + text + RESET)
}

func LightMagenta(text string) string {
	return (LIGHT_MAGENTA + text + RESET)
}

func LightPurple(text string) string {
	return (LIGHT_PURPLE + text + RESET)
}

func LightCyan(text string) string {
	return (LIGHT_CYAN + text + RESET)
}

func White(text string) string {
	return (WHITE + text + RESET)
}
