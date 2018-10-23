package log

import (
	"fmt"
	"os"
	"os/user"
	"strings"
	"sync"
	"time"

	color "github.com/multiverse-os/cli-framework/text/color"
)

// TODO: Add flag to increase resolution of the timestamp for software that needs greater resolution
type Logger struct {
	AppName   string
	Entries   []Entry
	Path      string
	Filename  string
	FileMutex sync.Mutex
	LogFile   *os.File
	FileData  []byte
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
		AppName:   name,
		Path:      logPath,
		Filename:  logFilename,
		Entries:   []Entry{},
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
	fmt.Println(logEntry.Type.FormattedString(true) + color.White("["+logEntry.Time.Format("Jan _2 15:04")+"] ") + logEntry.Text)
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

// TODO: Should switch to streaming method of writing to log file to avoid lag periods
// and bottlenecking on io
func (self *Logger) OpenLog() (err error) {
	defer CloseLog()
	self.FileMutex.Lock()
	self.LogFile, err = os.OpenFile(self.LogFilePath(), os.O_APPEND|os.O_WRONLY, 0660)
	self.FileMutex.Unlock()
	return err
}

// TODO: This should be called in signal cancel and shutdown
func (self *Logger) CloseLog() (err error) {
	self.FileMutex.Lock()
	self.LogFile.Close()
	self.FileMutex.Unlock()
}

func (self Logger) AppendToLog(text string) {
	// TODO: Add the ability to write to file as JSON (and probably XML)
	self.FileMutex.Lock()
	if _, err = file.WriteString(text); err != nil {
		self.FileMutex.Unlock()
		self.ToFile = false
		self.ToStdOut = true
		self.Log(FATAL, err.Error())
	} else {
		self.FileMutex.Unlock()
	}
}

func (self Logger) LogFilePath() string {
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
