package cli

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	ansi "github.com/multiverse-os/cli/terminal/ansi"
)

const prefixSize = 2
const tabSize = 4

func Prefix() string { return strings.Repeat(" ", 2) }
func Tab() string    { return strings.Repeat(" ", tabSize) }

type Stdio struct {
	In  FileReader
	Out FileWriter
	Err io.Writer
}

type FileWriter interface {
	io.Writer
	Fd() uintptr
}

type FileReader interface {
	io.Reader
	Fd() uintptr
}

type Outputs []Output

// TODO: This provide the formatter logic for extending colors ontop of fmt by
// adding new %X type logic. we can then add %{blue}%{bold} or like css
// %{color:blue;ansi:bold;}
// https://github.com/nhooyr/color
// This is important because its also the founation for a nice implementaiton of
// locales without relying on outside depndencies

// The hex output for this formatter is nice, rest is ok
// https://github.com/go-ffmt/ffmt

// A lot of logic should be handled by either terminal or text libraries (and
// these may be merged again)
// https://github.com/jedib0t/go-pretty
// Text
//
// Utility functions to manipulate text with or without ANSI escape sequences. Most of the functions available are used in one or more of the other packages here.
//     Align text horizontally or vertically
//         text/align.go and text/valign.go
//     Colorize text
//         text/color.go
//     Cursor Movement
//         text/cursor.go
//     Format text (convert case)
//         text/format.go
//     String Manipulation (Pad, RepeatAndTrim, RuneCount, Trim, etc.)
//         text/string.go
//     Transform text (UnixTime to human-readable-time, pretty-JSON, etc.)
//         text/transformer.go
//     Wrap text
//         text/wrap.go

// TODO: Formatters like this will provide the best way forward for not just
// doing ANSI better, but also provide simple way to handle localizaiton.
// https://github.com/kr/pretty/blob/main/formatter.go

// TODO: Consider rebuilding the table library, do it in such a way that each of
// the visual compoenents are mapped to a type that can be loaded like a theme.
// https://github.com/konojunya/go-frame/blob/master/frame.go

// The same should be with tree structures. Also the tree structure should
// support horizontal and vertical printing.

// THe concept of wrapping at 80 is great and would simplify our help output code
// https://github.com/PraserX/afmt

// TODO: Use terminal library to read width and add ability to format to one side
// Logging that ideally is not too bloated to get in the way, support for
// overriding by passing your logger's os.Writer but enough complexity to be
// useful in many use cases.

type Output struct {
	//Timestamp string
	prefix    string
	stripANSI bool
	file      io.Writer
}

// TODO: When this goes into its own package, this should be moved to its own file
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

// In the future we may want to migrate to a theme system that we define ansi
// code for each log level. Maybe regex colors, and primary, secondary,
// contrast (which will be used when printing values on debug, help text, and
// version)
func VarInfo(value interface{}) string {
	return ansi.Bold(ansi.White("[")) + ansi.Bold(ansi.Blue(fmt.Sprintf("%T", value))) + ansi.Bold(ansi.White("=")) + ansi.Green(fmt.Sprintf("%s", value)) + ansi.Bold(ansi.White("]"))
}

func DebugInfo(functionName string) string {
	return ansi.Bold(ansi.White("[")) + ansi.SkyBlue(functionName) + ansi.Bold(ansi.White("]"))
}

func merge(textParts ...interface{}) (output string) {
	for _, part := range textParts {
		output += fmt.Sprint(part, " ")
	}
	return output
}

// Value Assignment Chaining //////////////////////////////////////////////////
func (self Output) Prefix(prefix string) Output {
	self.prefix = prefix
	return self
}

// TODO: By using the method of coloring that does it via % using fmt override.
// WE can just not place the ANSI in. Since filelogging is more likely usecase
// it will use less resources in the long run to do that approach over the strip
// approach
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
func (self Outputs) Write(text ...string) {
	for _, output := range self {
		if output.stripANSI {
			fmt.Fprintf(output.file, "%s", output.prefix+ansi.Regex.ReplaceAllString(fmt.Sprint(strings.Join(text, " ")), "")+"\n")
		} else {
			fmt.Fprint(output.file, output.prefix, strings.Join(text, " "), "\n")
		}
	}
}

// TODO: Would be much better to just allow providing a closure, or a
// surrounding fucntion, whatever that function does. And in our case it will be
// coloring. That way its more flexible, needs less color specific code, and
// allows developers to pick colors.
func (self Outputs) Log(level LogLevel, output ...string) {
	var levelOutput string
	switch level {
	case DEBUG:
		levelOutput = ansi.Blue(level.String())
	case WARNING:
		levelOutput = ansi.Olive(level.String())
	case ERROR:
		levelOutput = ansi.Red(level.String())
	case FATAL:
		levelOutput = ansi.Maroon(level.String())
	default:
		levelOutput = ansi.Purple(level.String())
	}
	self.Write(ansi.Bold(ansi.White("[")) + levelOutput + ansi.Bold(ansi.White("]")) + strings.Join(output, " "))
}

//
// Ultra Minimal Multiple Output Logging System
///////////////////////////////////////////////////////////////////////////////
// Basic Levels, Debug Mode, Errors, Warnings, Fatals that exit, and flexible
// access so it doesn't get in the way.
//
// Have your own logger? No problem, just append it's io.Writer to CLI.Output.
// Want logging to Terminal & Logfile? No problem, append both to CLI.Output.
// Want both and output to a website? No problem.
func (self *CLI) Output(text ...interface{}) {
	self.Outputs.Write(merge(text))
}

func (self *CLI) Log(level LogLevel, text ...interface{}) {
	if level != DEBUG || level == DEBUG && self.Debug {
		self.Outputs.Log(level, merge(text, " "))
	}
}

// TODO: Can make this even better by having it return a function then we only need to pass the desription
func (self *CLI) benchmark(startedAt time.Time, description string) {
	self.Outputs.Log(DEBUG, ansi.Bold(DebugInfo("Benchmark")), ansi.Green(description), ansi.Bold(ansi.White("[")), ansi.Green(fmt.Sprintf("%v", time.Since(startedAt))), ansi.Bold(ansi.White("]")))
}
