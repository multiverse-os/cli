package log

import color "github.com/multiverse-os/cli-framework/text/color"

type LogType int

const (
	TRACE LogType = iota
	DEBUG
	INFO
	WARNING
	ERROR
	FATAL
	PANIC
)

func (self LogType) String() string {
	switch self {
	case TRACE:
		return "Trace"
	case DEBUG:
		return "Debug"
	case INFO:
		return "Info"
	case WARNING:
		return "Warning"
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

func (self LogType) FormattedString(brackets bool) string {
	switch self {
	case TRACE:
		if brackets {
			return color.Blue("[Trace]")
		} else {
			return color.Blue("Trace")
		}
	case DEBUG:
		if brackets {
			return color.White("[Debug]")
		} else {
			return color.White("Debug")
		}
	case INFO:
		if brackets {
			return color.Blue("[Info]")
		} else {
			return color.Blue("Info")
		}
	case WARNING:
		if brackets {
			return color.Warning("[Warning]")
		} else {
			return color.Warning("Warning")
		}
	case ERROR:
		if brackets {
			return color.Fail("[Error]")
		} else {
			return color.Fail("Error")
		}
	case FATAL:
		if brackets {
			return color.Fail("[Fatal Error]")
		} else {
			return color.Fail("Fatal Error")
		}
	case PANIC:
		if brackets {
			return color.Fail("[Panic]")
		} else {
			return color.Fail("Panic")
		}
	default:
		if brackets {
			return color.Gray("[Log]")
		} else {
			return color.Gray("Log")
		}
	}
}
