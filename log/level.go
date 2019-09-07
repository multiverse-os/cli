package log

import (
	color "github.com/multiverse-os/cli/text/ansi/color"
)

type LogLevel int

const (
	LOG LogLevel = iota
	DEBUG
	INFO
	WARNING
	NOTICE
	ERROR
	FATAL
	TRACE
	PANIC
)

type Verbosity int

const (
	NORMAL Verbosity = iota
	VERBOSE
	VERY_VERBOSE
	MAX_VERBOSE
)

const (
	QUIET Verbosity = -1
)

func (self Verbosity) Includes(level LogLevel) bool {
	switch self {
	case QUIET:
		// Quiet; No Logs
		return false
	case VERBOSE:
		// Verbose; -v
		switch level {
		case LOG, INFO, WARNING, ERROR, FATAL, PANIC:
			return true
		default:
			return false
		}
	case VERY_VERBOSE:
		// Very Verbose; -vv
		switch level {
		case LOG, DEBUG, INFO, WARNING, ERROR, FATAL, PANIC:
			return true
		default:
			return false
		}
	case MAX_VERBOSE:
		// Debug; -vvv
		switch level {
		case LOG, DEBUG, INFO, NOTICE, WARNING, ERROR, TRACE, FATAL, PANIC:
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
	WARN = WARNING
)

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
	case TRACE:
		return "TRACE"
	case FATAL:
		return "FATAL"
	case PANIC:
		return "PANIC"
	default:
		return "LOG"
	}
}

func (self LogLevel) StringWithANSI() string {
	switch self {
	case DEBUG:
		return color.White("DEBUG")
	case INFO:
		return color.Green("INFO")
	case WARN:
		return color.Yellow("WARNING")
	case NOTICE:
		return color.Silver("NOTICE")
	case ERROR:
		return color.Red("ERROR")
	case TRACE:
		return color.Blue("TRACE")
	case FATAL:
		return color.Maroon("FATAL")
	case PANIC:
		return color.Maroon("PANIC")
	default:
		return color.White("LOG")
	}
}
