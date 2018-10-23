package log

import (
	"os"
	"time"
)

type Entry struct {
	createdAt time.Time              `json:"created_at"`
	level     LogLevel               `json:"type"`
	message   string                 `json:"message"`
	errors    []error                `json:"errors"`
	values    map[string]interface{} `json:"-"`
}

func (self *Logger) Level(level LogLevel) *Logger {

}

func (self *Logger) Message(message string) *Logger {
	self.Log(level, message)
}

func (self *Logger) Entry(logEntry Entry) *Logger {
	entry := Entry{
		Level:     logEntry.Level,
		CreatedAt: time.Now(),
		Message:   logEntry.Message,
	}
	// TODO: Format then append to 'Logger' object or
	entry.Format()
	entry.AppendToLog()
	return self
}

func (self Logger) Log(level LogLevel, message string) {
	logEntry := Entry{
		Logger:    &self,
		Level:     level,
		CreatedAt: time.Now(),
		Message:   message,
	}
	self.Entries = append(self.Entries, logEntry)
	switch logEntry.Level {
	case FATAL, PANIC:
		os.Exit(1)
	}
}

func (self Logger) LogWithValues(level LogLevel, message string, values map[string]interface{}) {
	logEntry := Entry{
		Logger:    &self,
		Level:     level,
		CreatedAt: time.Now(),
		Message:   message,
		Values:    values,
	}
	// TODO: Format & Print; look at logrus for example on good log+values text
	// output.
	switch logEntry.level {
	case FATAL, PANIC:
		os.Exit(1)
	}
}
