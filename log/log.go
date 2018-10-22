package log

import (
	"fmt"
	"os"
	"time"

	color "github.com/multiverse-os/cli-framework/text/color"
)

// TODO: Add flag to increase resolution of the timestamp for software that needs greater resolution

func Print(logType LogType, text string) {
	fmt.Println(logType.FormattedString(true) + color.White("["+time.Now().Format("Jan _2 15:04")+"] ") + text)
	if logType == FATAL || logType == PANIC {
		os.Exit(1)
	}
}

func Trace(text string) {
	Print(TRACE, text)
}

func Info(text string) {
	Print(INFO, text)
}

func Warning(text string) {
	Print(WARNING, text)
}

func Warn(text string) {
	Print(WARNING, text)
}

func Error(err error) {
	Print(ERROR, err.Error())
}

func FatalError(err error) {
	Print(FATAL, err.Error())
}

func Fatal(text string) {
	Print(FATAL, text)
}

func Panic(text string) {
	Print(PANIC, text)
}
