package log

import (
	"fmt"
	"os"
	"strings"
	"time"

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
	Entries   []Entry
	Path      string
	Filename  string
	Verbosity int
	ToFile    bool
	ToStdOut  bool
	JSON      bool
}

func NewLogger(name string, verbosity int, toFile, toStdOut, json bool) Logger {
	name = strings.ToLower(name)
	logPath := ("/var/log/" + name + "/")
	logFilename := (name + ".log")
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		fmt.Println("file does not exist")
		os.MkdirAll(logPath, 0660)
		os.OpenFile((logPath + logFilename), os.O_RDONLY|os.O_CREATE, 0660)
		// TODO: If this fails to create in var log, we should make the default log directory within user home directory
	} else {
		fmt.Println("file exists")
	}
	return Logger{
		AppName:   name,
		Path:      logPath,
		Entries:   []Entry{},
		Filename:  logFilename,
		Verbosity: verbosity,
		ToFile:    toFile,
		ToStdOut:  toStdOut,
		JSON:      json,
	}
}

func (self *Logger) NewLog(logType LogType, text string) {
	self.Log(logType, text)
}

func (self Logger) Log(logType LogType, text string) {
	logEntry := Entry{
		Logger: &self,
		Type:   logType,
		Time:   time.Now(),
		Text:   text,
	}
	self.Entries = append(self.Entries, logEntry)
	fmt.Println(logEntry.Type.FormattedString(true) + color.White("["+logEntry.Time.String()+"] ") + logEntry.Text)
	if logEntry.Type == FATAL || logEntry.Type == PANIC {
		os.Exit(1)
	}
}

func (self Logger) Info(text string) {
	self.Log(INFO, text)
}

func (self Logger) Warning(text string) {
	self.Log(WARNING, text)
}

func (self Logger) Warn(text string) {
	self.Log(WARNING, text)
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
