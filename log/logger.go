package log

import (
	"fmt"
	"os"
	"strings"

	color "github.com/multiverse-os/cli-framework/text/color"
)

type AppLogger interface {
	Log(text string)
	Info(text string)
	Warning(text string)
	Error(text string)
	FatalError(text string)
}

type Logger struct {
	AppName   string
	Path      string
	Filename  string
	Verbosity int
	ToFile    bool
	ToStdOut  bool
	JSON      bool
}

func NewLogger(name string, verbosity int, toFile, toStdOut, json bool) Logger {
	name = strings.ToLower(name)
	logPath = ("/var/log/" + name + "/")
	logFilename = (name + ".log")
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		os.MkdirAll(logPath, 0660)
		os.OpenFile((logPath + logFilename), os.O_RDONLY|os.O_CREATE, 0660)
	}
	return Logger{
		AppName:   name,
		Path:      logPath,
		Filename:  logFilename,
		Verbosity: verbosity,
		ToFile:    toFile,
		ToStdOut:  toStdOut,
		JSON:      json,
	}
}

func (self Logger) Log(logType LogType, text string) {
	switch logType {
	case INFO:
		if self.ToStdOut {
			fmt.Println(color.Info("[INFO] ") + text)
		}
	case WARNING:
		if self.ToStdOut {
			fmt.Println(color.Warning("[Warning] ") + text)
		}
	case ERROR:
		if self.ToStdOut {
			fmt.Println(color.Fail("[Error] ") + text)
		}
	case FATAL:
		if self.ToStdOut {
			fmt.Println(color.Fail("[Fatal Error] ") + text)
		}
		os.Exit(1)
	default:
		if self.ToStdOut {
			fmt.Println(color.Gray("[LOG] ") + text)
		}
	}
}

func (self Logger) Info(text string) {
	self.Log(INFO, text)
}

func (self Logger) Warning(text string) {
	self.Log(WARNING, text)
}

func (self Logger) Error(err error) {
	self.Log(ERROR, err.Error())
}

func (self Logger) FatalError(err error) {
	self.Log(FATAL, err.Error())
}

func (self Logger) AppendToLog(text string) {
	// TODO: Add the ability to write to file as JSON (and probably XML)
	file, err := os.OpenFile((self.Path + self.Filename), os.O_APPEND|os.O_WRONLY, 0660)
	if err != nil {
		self.ToFile = false
		self.ToStdOut = true
		self.Log(FATAL, err.Error())
	}
	defer file.Close()
	if _, err = file.WriteString(text); err != nil {
		self.ToFile = false
		self.ToStdOut = true
		self.Log(FATAL, err.Error())
	}
}
