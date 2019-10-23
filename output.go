package cli

import (
	"fmt"
	"io"
	"os"
	"regexp"
)

const ansi = "\x1b[@-_][0-?]*[ -/]*[@-~]"

var ansiRegex = regexp.MustCompile(ansi)

type Output struct {
	//Timestamp string
	Prefix    string
	StripANSI bool
	File      io.Writer
}

//var stylize = func(text string, color string) string {
//	return prefix + text + suffix
//}

func TerminalOutput() *Output {
	//palette, err := colorful.WarmPalette(10)
	//if err != nil {
	//	return nil, err
	//}
	return &Output{
		//Theme: Theme{
		//	//Primary:   palette[0],
		//	//Secondary: palette[1],
		//	//Contrast:  palette[2],
		//},
		//Styles:         map[string]string{},
		File:      os.Stdout,
		StripANSI: false,
	}
}

func LogfileOutput(filename string) *Output {
	// TODO: Create if does not exist, including path. For now, lets just declare
	// log rotation as out of scope. If its not widely supported you require a
	// third party program to handle it anyways, so just leave it to that.
	return &Output{
		StripANSI: true,
	}
}

func (self *Output) Output(output string) {
	if self.StripANSI {
		ansiRegex.ReplaceAllString(output, "")
	}
	fmt.Fprintf(self.File, self.Prefix+output)
}
