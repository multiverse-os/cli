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

func (self Output) Prefix(prefix string) Output {
	self.prefix = prefix
	return self
}

func (self Output) StripANSI(strip bool) Output {
	self.stripANSI = strip
	return self
}

// TODO: Improve this by implmenting Fprintf locally, so we can provide similar
// functionality to Ouput and Write.
func (self Output) Write(output ...string) {
	if self.stripANSI {
		for _, text := range output {
			text = ansiRegex.ReplaceAllString(text, "")
		}
	}
	fmt.Fprintf(self.file, self.prefix+strings.Join(output, " "))
}

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
		return ""
	}
}

func (self Output) Log(level LogLevel, output ...string) {
	levelOutput := style.Bold(level.String())
	switch level {
	case INFO:
		levelOutput = color.Silver(levelOutput)
	case DEBUG:
		levelOutput = color.Purple(levelOutput)
	case WARNING:
		levelOutput = color.Olive(levelOutput)
	case ERROR:
		levelOutput = color.Silver(levelOutput)
	case FATAL:
		levelOutput = color.Silver(levelOutput)
	}

	self.Write(fmt.Sprintf("[%v] %s", levelOutput, strings.Join(output, " ")))
}
