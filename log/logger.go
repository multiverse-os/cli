package log

import (
	"time"
)

// TODO: Write tests for:
//   1) Hook Logic
//   2) Log rotation
//   3) Verbosity Levels
//   4) Write to terminal, write to file

type Logger struct {
	Name                string
	format              Format
	verbosity           Verbosity
	timestampResolution TimestampResolution
	entries             []Entry
	hooks               map[LogLevel]map[HookType][]*Hook
	outputs             []LogOutput
}

func NewLogger(name string, resolution TimestampResolution, verbosity Verbosity) Logger {
	if name == "" {
		Fatal("Name attribute is required to initialize log file")
	}
	return Logger{
		Name:                name,
		verbosity:           verbosity,
		timestampResolution: resolution,
	}
}

func DefaultLogger(name string, stdOut bool, fileOut bool) Logger {
	logger := Logger{
		Name:                name,
		format:              ANSI,
		verbosity:           NORMAL,
		timestampResolution: MINUTE,
	}
	if stdOut {
		logger.OutputToTerminal(ANSI)
	}
	if fileOut {
		logger.OutputToDefaultLogFile(JSON)
	}
	return logger
}

func TerminalLogger(name string, format Format, resolution TimestampResolution, verbosity Verbosity) Logger {
	logger := NewLogger(name, resolution, verbosity)
	logger.OutputToTerminal(format)
	return logger
}

func FileLogger(name string, format Format, logPath string, resolution TimestampResolution, verbosity Verbosity) Logger {
	logger := NewLogger(name, resolution, verbosity)
	if logPath, ok := FindOrCreateFile(logPath); !ok {
		Fatal("Failed to initialize default log path: '" + logPath + "'")
	}
	logger.OutputToFile(format, logPath)
	return logger
}

// Append To Outputs
///////////////////////////////////////////////////////////////////////////////
func (self *Logger) Append(entry Entry) {
	if self.verbosity.Includes(entry.level) {
		if self.HasOutputs() {
			for _, out := range self.outputs {
				out.Append(entry)
			}
		} else {
			Info("Logger has no outputs defined; defaulting to ANSI styled terminal output.")
			self.OutputToTerminal(ANSI)
		}
	}
}

// Outputs
///////////////////////////////////////////////////////////////////////////////
func (self *Logger) OutputToFile(format Format, outputPath string) {
	if outputPath, ok := FindOrCreateFile(outputPath); !ok {
		Fatal("Failed to initialized specified log path: '" + outputPath + "'")
	} else {
		logFile := &LogFile{
			format: format,
			path:   outputPath,
		}
		err := logFile.Open()
		if err != nil {
			Fatal(err.Error())
		} else {
			self.outputs = append(self.outputs, logFile)
		}
	}
}

func (self *Logger) OutputToDefaultLogFile(format Format) {
	if logPath, ok := FindOrCreateFile(DefaultUserLogPath(self.Name)); !ok {
		Fatal("Failed to initialize default user log path: '" + logPath + "'")
	} else {
		self.OutputToFile(format, logPath)
	}
}

func (self *Logger) OutputTo(output Output, format Format) {
	switch output {
	case FILE:
		self.OutputToFile(format, DefaultLogPath(self.Name))
	case TERMINAL:
		self.outputs = append(self.outputs, &Terminal{
			format: format,
		})
	}
}

func (self *Logger) OutputToTerminal(format Format) {
	self.OutputTo(TERMINAL, format)
}

func (self Logger) HasOutputs() bool   { return (len(self.outputs) != 0) }
func (self Logger) HasValues() bool    { return (len(self.outputs) != 0) }
func (self Logger) HasErrors() bool    { return (len(self.outputs) != 0) }
func (self Logger) HasTimestamp() bool { return (self.timestampResolution != DISABLED) }

// Create Log Entries
///////////////////////////////////////////////////////////////////////////////
func (self Logger) Log(level LogLevel, message string) Entry {
	return Entry{
		createdAt:           time.Now(),
		format:              self.format,
		level:               level,
		message:             message,
		timestampResolution: self.timestampResolution,
	}
}
func (self Logger) Info(message string)    { Log(INFO, message).Append() }
func (self Logger) Warning(message string) { Log(WARNING, message).Append() }
func (self Logger) Warn(message string)    { Log(WARN, message).Append() }
func (self Logger) Error(err error)        { Log(ERROR, err.Error()).Append() }
func (self Logger) Fatal(message string)   { Log(FATAL, message).Append() }
func (self Logger) Panic(message string)   { Log(PANIC, message).Append() }

// Graceful Shutdown
///////////////////////////////////////////////////////////////////////////////
func (self *Logger) Shutdown() {
	Info("Shutdown initiated, gracefully closing outputs...")
	for _, output := range self.outputs {
		output.Close()
	}
}
