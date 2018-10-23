package log

import (
	color "github.com/multiverse-os/cli-framework/text/color"
)

type LogLevel int

const (
	LOG LogLevel = iota
	DEBUG
	INFO
	WARNING
	NOTICE
	ERROR
	TRACE
	FATAL_ERROR
	PANIC
)

// Level Aliasing
const (
	WARN  = WARNING
	FATAL = FATAL_ERROR
)

func (self LogLevel) ColorString() string {
	switch self {
	case DEBUG:
		return color.Debug("Debug")
	case INFO:
		return color.Info("Info")
	case WARN:
		return color.Warning("Warning")
	case NOTICE:
		return color.Notice("Notice")
	case ERROR:
		return color.Error("Error")
	case TRACE:
		return color.Trace("Trace")
	case FATAL:
		return color.Fatal("Fatal Error")
	case PANIC:
		return color.Panic("Panic")
	default:
		return color.Gray("Log")
	}
}

func (self LogLevel) String() string {
	switch self {
	case DEBUG:
		return "Debug"
	case INFO:
		return "Info"
	case WARN:
		return "Warning"
	case NOTICE:
		return "Notice"
	case ERROR:
		return "Error"
	case FATAL:
		return "Fatal Error"
	case PANIC:
		return "Panic"
	default:
		return "Log"
	}
}
