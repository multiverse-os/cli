package cli

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	color "github.com/multiverse-os/cli/framework/terminal/ansi/color"
	style "github.com/multiverse-os/cli/framework/terminal/ansi/style"
)

// Logging that ideally is not too bloated to get in the way, support for
// overriding by passing your logger's os.Writer but enough complexity to be
// useful in many use cases.

const ansi = "\x1b[@-_][0-?]*[ -/]*[@-~]"

var ansiRegex = regexp.MustCompile(ansi)

type Output struct {
	//Timestamp string
	prefix    string
	stripANSI bool
	file      io.Writer
}

// TODO: When this goes into its own package, this should be moved to its own
// file
type LogLevel int

const (
	INFO LogLevel = iota
	DEBUG
	WARNING
	ERROR
	FATAL
)

func (self LogLevel) String() string {
	switch self {
	case INFO:
		return "INFO"
	case DEBUG:
		return "DEBUG"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	default:
		return Blank
	}
}

// In the future we may want to migrate to a theme system that we define ansi
// code for each log level. Maybe regex colors, and primary, secondary,
// contrast (which will be used when printing values on debug, help text, and
// version)
func merge(text []string) string { return strings.Join(text, Blank) }

// TODO: Use to construct themes
func purple(text ...string) string  { return color.Purple(merge(text)) }
func green(text ...string) string   { return color.Green(merge(text)) }
func lime(text ...string) string    { return color.Lime(merge(text)) }
func silver(text ...string) string  { return color.Silver(merge(text)) }
func blue(text ...string) string    { return color.Blue(merge(text)) }
func skyBlue(text ...string) string { return color.Blue(merge(text)) }
func olive(text ...string) string   { return color.Olive(merge(text)) }
func red(text ...string) string     { return color.Red(merge(text)) }
func maroon(text ...string) string  { return color.Maroon(merge(text)) }
func bold(text ...string) string    { return style.Bold(merge(text)) }

func thin(text ...string) string  { return style.Thin(merge(text)) }
func white(text ...string) string { return color.White(merge(text)) }

func brackets(text ...string) string    { return bold("[") + merge(text) + bold("]") }
func parenthesis(text ...string) string { return bold("(") + merge(text) + bold(")") }

// helpers
func debugInfo(text string) string      { return Brackets(skyBlue(text)) }
func varInfo(name, value string) string { return blue(Brackets(bold(name) + white("=") + green(value))) }

//
// Public Methods
///////////////////////////////////////////////////////////////////////////////

// Value Assignment Chaining //////////////////////////////////////////////////
func (self Output) Prefix(prefix string) Output {
	self.prefix = prefix
	return self
}

func (self Output) StripANSI(strip bool) Output {
	self.stripANSI = strip
	return self
}

// Default Ouput Locations ////////////////////////////////////////////////////
func TerminalOutput() Output {
	//palette, err := colorful.WarmPalette(10)
	//if err != nil {
	//	return nil, err
	//}
	return Output{
		//Theme: Theme{
		//	//Primary:   palette[0],
		//	//Secondary: palette[1],
		//	//Contrast:  palette[2],
		//},
		file:      os.Stdout,
		stripANSI: false,
	}
}

func LogfileOutput(filename string) Output {
	// TODO: Create if does not exist, including path. For now, lets just declare
	// log rotation as out of scope. If its not widely supported you require a
	// third party program to handle it anyways, so just leave it to that.
	return Output{
		stripANSI: true,
	}
}

// Writing To Outputs /////////////////////////////////////////////////////////
// TODO: Improve this by implmenting Fprintf locally, so we can provide similar
// functionality to Ouput and Write.
func (self Output) Write(output ...string) {
	if self.stripANSI {
		for _, text := range output {
			text = ansiRegex.ReplaceAllString(text, Blank)
		}
	}
	fmt.Fprintf(self.file, self.prefix+strings.Join(output, Space))
}

func (self Output) Log(level LogLevel, output ...string) {
	levelOutput := bold(level.String())
	switch level {
	case DEBUG:
		levelOutput = purple(levelOutput)
	case WARNING:
		levelOutput = olive(levelOutput)
	case ERROR:
		levelOutput = red(levelOutput)
	case FATAL:
		levelOutput = maroon(levelOutput)
	default:
		levelOutput = skyBlue(levelOutput)
	}
	self.Write(white(Brackets(purple(levelOutput)), strings.Join(output, Space)))
}
