package log

import (
	"fmt"
	"os"

	"github.com/multiverse-os/cli-framework/text/color"
)

func Print(logType LogType, text string) {
	switch logType {
	case INFO:
		fmt.Println(color.Info("[INFO] ") + text)
	case WARNING:
		fmt.Println(color.Warning("[Warning] ") + text)
	case ERROR:
		fmt.Println(color.Fail("[Error] ") + text)
	case FATAL:
		fmt.Println(color.Fail("[Fatal Error] ") + text)
		os.Exit(1)
	default:
		fmt.Println(color.Gray("[LOG] ") + text)
	}
}

func Log(text string) {
	Print(DEFAULT, text)
}

func Info(text string) {
	Print(INFO, text)
}

func Warning(text string) {
	Print(WARNING, text)
}

func Error(err error) {
	Print(ERROR, err.Error())
}

func FatalError(err error) {
	Print(FATAL, err.Error())
}
