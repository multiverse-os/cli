package log

import (
	"io"
	"os"
	"os/user"
	"strings"
)

type Format int

// TODO: The ideal result of this would be making LogFile and StdOut/StdErr
//       share a common Interface, then they can be added via the
//       Logger.WriteStreams(). Anything written to Entries, should be
//       written to each WriteStream(s) using mutexes so in a thread
//       safe way.
//       The most ideal thing would be to write to Entries using mutex,
//       then use this slice to write to the io.Writers from

// TODO: Determine LogLevel/Verbosity groupings

type Logger struct {
	AppName        string
	Verbosity      int
	Entries        []Entry
	TimeResolution TimeResolution
	File           LogFile
	Outputs        []io.Writer
}

func NewLogger(name string, resolution TimeResolution, verbosity int, toFile, toStdOut, json bool) Logger {
	name = strings.ToLower(name)
	logPath := ("/var/log/" + name + "/")
	logFilename := (name + ".log")
	if _, err := os.Stat(logPath); os.IsNotExist(err) {
		os.MkdirAll(logPath, 0660)
		os.OpenFile((logPath + logFilename), os.O_RDONLY|os.O_CREATE, 0660)
		if _, err := os.Stat(logPath); os.IsNotExist(err) {
			if _, err := os.Stat((UserLogPath(name) + logFilename)); os.IsNotExist(err) {
				os.MkdirAll(UserLogPath(name), 0660)
				os.OpenFile((UserLogPath(name) + logFilename), os.O_RDONLY|os.O_CREATE, 0660)
			}
			logPath = UserLogPath(name)
		}
	}
	return Logger{
		AppName:        name,
		Verbosity:      verbosity,
		TimeResolution: resolution,
		Entries:        []Entry{},
		File: LogFile{
			Path:     logPath,
			Filename: logFilename,
		},
	}
}

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

func (self LogFile) FilePath() string {
	return (self.Path + self.Filename)
}

func UserLogPath(appName string) string {
	home := os.Getenv("XDG_CONFIG_HOME")
	if home != "" {
		return (home + "/.local/share/" + strings.ToLower(appName) + "/")
	} else {
		home = os.Getenv("HOME")
		if home != "" {
			return (home + "/.local/share/" + strings.ToLower(appName) + "/")
		} else {
			currentUser, err := user.Current()
			if err != nil {
				FatalError(err)
			}
			home = currentUser.HomeDir
			return (home + "/.local/share/" + strings.ToLower(appName) + "/")
		}
	}
}
