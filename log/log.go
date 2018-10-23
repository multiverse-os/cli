package log

import (
	"fmt"
	"os"
	"time"
)

//
// Minimialistic Logging
///////////////////////////////////////////////////////////////////////////////
// Simple log.Print(INFO, "Text") and associated alias functions such as
// Trace("Test trace") or Info("Print info to StdOut") that does not require
// a Logger object or creation of LogEntry objects for very quick access to
// consistent StdOut logging when needed.
func Print(level LogLevel, message string) {
	logEntry := Entry{
		Level:     level,
		CreatedAt: time.Now(),
		Message:   message,
	}
	// TODO: Format & Print LOG_ENTRY
	fmt.Println(logEntry.Message)
	switch level {
	case FATAL, PANIC:
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
	Print(WARN, text)
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
