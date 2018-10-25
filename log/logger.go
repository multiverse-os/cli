package log

import (
	"errors"
)

type Logger struct {
	Name           string
	Verbosity      int
	TimeResolution TimeResolution
	Entries        []Entry
	Hooks          map[LogLevel]map[HookType][]*Hook
	Outputs        []LogOutput
}

func NewLogger(name string, resolution TimeResolution, verbosity int) Logger {
	return Logger{
		Name:           name,
		Verbosity:      verbosity,
		TimeResolution: resolution,
		Entries:        []Entry{},
	}
}

func NewFileLogger(name string, resolution TimeResolution, verbosity int, format Format, logPath string) Logger {
	logger := Logger{
		Name:           name,
		Verbosity:      verbosity,
		TimeResolution: resolution,
		Entries:        []Entry{},
	}
	logger.AddFileOutput(format, UserLogPath(name))
	return logger
}

func (self *Logger) InitLogFile(logFilePath string) (string, error) {
	if ok := FindOrCreateFile(logFilePath); ok {
		return logFilePath, nil
	} else {
		userLogFilePath := UserLogPath(self.Name)
		if ok := FindOrCreateFile(userLogFilePath); ok {
			return userLogFilePath, nil
		} else {
			return "", errors.New("Failed to initialize log file")
		}
	}
	return logFilePath, nil
}

//
// Outputs
///////////////////////////////////////////////////////////////////////////////
func (self *Logger) AddFileOutput(format Format, path string) {
	logFilePath, err := self.InitLogFile(path)
	if err != nil {
		logFile := &LogFile{
			format: format,
			path:   logFilePath,
		}
		err := logFile.Open()
		if err != nil {
			FatalError(err)
		} else {
			self.Outputs = append(self.Outputs, logFile)
		}
	} else {
		FatalError(err)
	}
}

func (self *Logger) AddOutput(output Output, format Format) {
	switch output {
	case FILE:
		self.AddFileOutput(format, UserLogPath(self.Name))
	case STDOUT:
		self.Outputs = append(self.Outputs, &StdOut{
			format: format,
		})
	}
}

//
// Graceful Shutdown
///////////////////////////////////////////////////////////////////////////////
func (self *Logger) Shutdown() {
	for _, output := range self.Outputs {
		output.Close()
	}
}

//
// Standard Log Function Aliases
///////////////////////////////////////////////////////////////////////////////
func (self Logger) Info(text string) {
	self.Log(INFO, text)
}

func (self Logger) Warning(text string) {
	self.Log(WARNING, text)
}

func (self Logger) Warn(text string) {
	self.Log(WARN, text)
}

func (self Logger) Error(err error) {
	self.Log(ERROR, err.Error())
}

func (self Logger) FatalError(err error) {
	self.Log(FATAL, err.Error())
}

func (self Logger) Fatal(text string) {
	self.Log(FATAL, text)
}

func (self Logger) Panic(text string) {
	self.Log(FATAL, text)
}
