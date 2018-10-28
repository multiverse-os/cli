package log

import (
	text "github.com/multiverse-os/cli-framework/text"
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

type VerbosityLevel int

const (
	NORMAL VerbosityLevel = iota
	VERBOSE
	VERY_VERBOSE
	MAX_VERBOSE
)

const (
	QUIET VerbosityLevel = -1
)

func (self VerbosityLevel) Includes(level LogLevel) bool {
	switch self {
	case QUIET:
		// Quiet; No Logs
		return false
	case VERBOSE:
		// Verbose; -v
		switch level {
		case LOG, INFO, WARNING, ERROR, FATAL_ERROR, PANIC:
			return true
		default:
			return false
		}
	case VERY_VERBOSE:
		// Very Verbose; -vv
		switch level {
		case LOG, DEBUG, INFO, WARNING, ERROR, FATAL_ERROR, PANIC:
			return true
		default:
			return false
		}
	case MAX_VERBOSE:
		// Debug; -vvv
		switch level {
		case LOG, DEBUG, INFO, NOTICE, WARNING, ERROR, TRACE, FATAL_ERROR, PANIC:
			return true
		default:
			return false
		}
	default:
		// Normal; verbosity == 0
		return true
	}
}

// Level Aliasing
const (
	WARN  = WARNING
	FATAL = FATAL_ERROR
)

func (self LogLevel) StyledString() string {
	switch self {
	case DEBUG:
		return text.Debug("DEBUG")
	case INFO:
		return text.Info("INFO")
	case WARN:
		return text.Warning("WARNING")
	case NOTICE:
		return text.Notice("NOTICE")
	case ERROR:
		return text.Error("ERROR")
	case TRACE:
		return text.Trace("TRACE")
	case FATAL:
		return text.Fatal("FATAL")
	case PANIC:
		return text.Panic("PANIC")
	default:
		return text.Gray("LOG")
	}
}

func (self LogLevel) String() string {
	switch self {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARN:
		return "WARNING"
	case NOTICE:
		return "NOTICE"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	case PANIC:
		return "PANIC"
	default:
		return "LOG"
	}
}
