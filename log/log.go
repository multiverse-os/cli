package log

import (
	"errors"
	"fmt"
	"os"

	"github.com/multiverse-os/cli-framework/text/color"
)

// TODO: Should be able to specify log location, default to standard /var/log/{program_name}
// should be able to output JSON logs

type LogType int

const (
	DEFAULT LogType = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func Print(logType LogType, text string) {
	fmt.Println("")
	switch logType {
	case INFO:
		fmt.Println(color.Blue("[INFO] ") + text)
	case WARNING:
		fmt.Println(color.Blue("[INFO] ") + text)
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
	fmt.Println(color.Red("[Error] ") + err.ToString())
}

func FatalError(err error) {
	// TODO: Logs then exits with 1
	fmt.Println(color.Red("[Fatal Error] ") + err.ToString())
	os.Exit(1)
}
