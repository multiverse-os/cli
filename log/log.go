package log

import (
	"errors"
	"fmt"
	"os"

	"github.com/multiverse-os/cli-framework/text/color"
)

// TODO: Should be able to specify log location, default to standard /var/log/{program_name}
// should be able to output JSON logs

type Logger struct {
	AppName   string
	Path      string
	Filename  string
	Verbosity int
	StdOut    bool
	JSON      bool
}

var Logger = Logger{
	AppName:   "",
	Path:      "/var/log/{app_name}",
	Filename:  "{app_name}.log",
	Verbosity: 1,
	StdOut:    true,
	JSON:      false,
}

type LogType int

const (
	DEFAULT LogType = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func Print(logType LogType, text string) {
	switch logType {
	case INFO:
		if Logger.StdOut {
			fmt.Println(color.Info("[INFO] ") + text)
		}
	case WARNING:
		if Logger.StdOut {
			fmt.Println(color.Warning("[Warning] ") + text)
		}
	case ERROR:
		Error(errors.New(text))
	case FATAL:
		FatalError(errors.New(text))
	default:
		PrintLog(text)
	}
}

func PrintLog(text string) {
	// TODO: Ideally this should get info from application name and use that instead of LOG
	fmt.Println(color.Gray("[LOG] ") + text)
}

func Error(err error) {
	fmt.Println(color.Fail("[Error] ") + err.Error())
}

func FatalError(err error) {
	// TODO: Logs then exits with 1
	fmt.Println(color.Fail("[Fatal Error] ") + err.Error())
	os.Exit(1)
}
