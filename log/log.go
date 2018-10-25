package log

import (
	"fmt"
	"os"
	"time"
)

//
// Minimialistic Debug (Stdout) Logging
///////////////////////////////////////////////////////////////////////////////
// Simple log.Print(INFO, "Text") and associated alias functions such as
// Trace("Test trace") or Info("Print info to StdOut") that does not require
// a Logger object or creation of LogEntry objects for very quick access to
// consistent StdOut logging when needed.
func Print(level LogLevel, message string) {
	entry := Entry{
		createdAt: time.Now(),
		level:     level,
		message:   message,
	}
	// TODO: Format & Print LOG_ENTRY
	fmt.Println(entry.Message)
	// TODO: When the before/acter action hook is built, we can just
	// leverage that system
	switch entry.level {
	case FATAL, PANIC:
		os.Exit(1)
	}
}

func Info(text string) {
	Print(INFO, text)
}

func Notice(text string) {
	Print(NOTICE, text)
}

func Warning(text string) {
	Print(WARNING, text)
}

func Warn(text string) {
	Print(WARN, text)
}

func Trace(text string) {
	Print(TRACE, text)
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
