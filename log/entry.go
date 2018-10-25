package log

import (
	"encoding/json"
	"encoding/xml"
	"time"

	text "github.com/multiverse-os/cli-framework/text"
	color "github.com/multiverse-os/cli-framework/text/color"
)

type Entry struct {
	logger    *Logger                `json:"-"`
	createdAt time.Time              `json:"created_at"`
	level     LogLevel               `json:"level"`
	message   string                 `json:"message"`
	values    map[string]interface{} `json:"values"`
	errors    []error                `json:"errors"`
	format    Format                 `json:"-"`
}

type Format int

const (
	Default Format = iota
	DefaultWithANSI
	JSON
	XML
)

// Formatting
///////////////////////////////////////////////////////////////////////////////
// Output caps a chained log entry create, it writes the log to the parameter
// defined output without requiring the 'Logger' object. This allows for quick
// and simple log systems that can be called without any prior initialization
// of any variables and with no global variables.

// TODO: Need a generic function for logging to a file without needing a logger
// object

func (self Entry) Output() {
	if self.logger != nil {
		self.AppendToLogger(self.logger)
	} else {
		self.Append()
	}
}

func (self Entry) Format(format Format) Entry {
	self.format = format
	return self
}

func (self Entry) DefaultFormat() Entry {
	self.format = Default
	return self
}

func (self Entry) DefaultFormatWithANSI() Entry {
	self.format = DefaultWithANSI
	return self
}

func (self Entry) JSON() Entry {
	self.format = JSON
	return self
}

func (self Entry) XML() Entry {
	self.format = XML
	return self
}

func (self Entry) FormattedOutput() string {
	return self.FormattedString(self.format)
}

func (self Entry) FormattedString(format Format) string {
	switch format {
	case JSON:
		jsonOutput, err := json.Marshal(self)
		if err != nil {
			FatalError(err)
		}
		return string(jsonOutput)
	case XML:
		xmlOutput, err := xml.Marshal(self)
		if err != nil {
			FatalError(err)
		}
		return string(xmlOutput)
	case DefaultWithANSI:
		return text.Brackets(self.level.String()) + text.Brackets(color.White(self.Timestamp())) + " " + self.message
	}
	// Default; required to be outside because switchcase does not satisfy call
	// conditions to the compiler.
	return text.Brackets(self.level.String()) + text.Brackets(self.Timestamp()) + " " + self.message
}

func (self Entry) DefaultFormattedString() string {
	return self.FormattedString(Default)
}

func (self Entry) DefaultWithANSIFormattedString() string {
	return self.FormattedString(DefaultWithANSI)
}

func (self Entry) JSONFormattedString() string {
	return self.FormattedString(JSON)
}

func (self Entry) XMLFormattedString() string {
	return self.FormattedString(XML)
}

func (self Entry) String() string {
	return self.DefaultFormattedString()
}

// Append (print or write to defined outputs)
///////////////////////////////////////////////////////////////////////////////
func (self Entry) Append() {
	if self.logger != nil {
		self.AppendToLogger(self.logger)
	} else {
		stdout := StdOut{
			format: self.format,
		}
		stdout.Append(self)
	}
}

func (self *Logger) Append(e Entry) {
	e.AppendToLogger(self)
}

func (self Entry) AppendToLogger(logger *Logger) {
	if len(logger.Outputs) == 0 {
		logger.AddStdOutWithFormat(DefaultWithANSI)
	}
	for _, out := range logger.Outputs {
		if logger.Verbosity.IncludesLevel(self.level) {
			out.Append(self)
		}
	}
}

func (self Entry) AppendToFile(format Format, path string) (err error) {
	logFile := LogFile{
		format: format,
		path:   path,
	}
	err = logFile.Open()
	if err != nil {
		return nil
	}
	logFile.Append(self)
	return nil
}

func (self Entry) ToFile(format Format, path string) (err error) {
	return self.AppendToFile(format, path)
}

func (self Entry) WriteToFile(path string) (err error) {
	return self.AppendToFile(self.format, path)
}

func (self Entry) Print() {
	self.Append()
}

func (self Entry) StdOut() {
	self.Append()
}

// Chain-based Assignment
///////////////////////////////////////////////////////////////////////////////
// TODO: Should specifically message and values be validated and considered
// potentially dangerous user input?
func Log(message string) Entry {
	return Entry{
		createdAt: time.Now(),
		level:     LOG,
		format:    DefaultWithANSI,
		message:   message,
	}
}

func AppendEntry(e Entry) {
	entry := Entry{
		level:     e.level,
		createdAt: time.Now(),
		message:   e.message,
		format:    DefaultWithANSI,
		values:    e.values,
	}
	entry.Append()
}

func NewFormattedLog(level LogLevel, message string, format Format) Entry {
	return Entry{
		level:     level,
		message:   message,
		format:    format,
		createdAt: time.Now(),
	}
}

func NewLogWithFormat(e Entry, format Format) Entry {
	return LogWithFormat(e, format)
}

func LogWithFormat(e Entry, format Format) Entry {
	return Entry{
		level:     e.level,
		createdAt: time.Now(),
		message:   e.message,
		format:    format,
		values:    e.values,
	}
}

func NewEntry(e Entry) Entry {
	return Entry{
		level:     e.level,
		createdAt: time.Now(),
		message:   e.message,
		format:    e.format,
		values:    e.values,
	}
}

// Attribute: level
func Level(level LogLevel) Entry {
	return Entry{
		level:     level,
		format:    DefaultWithANSI,
		createdAt: time.Now(),
	}
}

func LogWithValues(level LogLevel, message string, values map[string]interface{}) Entry {
	return Entry{
		createdAt: time.Now(),
		level:     level,
		message:   message,
		format:    DefaultWithANSI,
		values:    values,
	}
}

// Attribute: message
func Message(message string) Entry {
	return Entry{
		createdAt: time.Now(),
		level:     LOG,
		format:    DefaultWithANSI,
		message:   message,
	}
}

// Attribute: values
func WithValues(values map[string]interface{}) Entry {
	return Entry{
		createdAt: time.Now(),
		level:     LOG,
		format:    DefaultWithANSI,
		values:    values,
	}
}

// Aliasing
//// Entry Aliases
func (self Entry) Log(level LogLevel, message string) Entry {
	return LogWithValues(level, message, make(map[string]interface{}))
}

func (self Entry) WithValues(values map[string]interface{}) Entry {
	self.values = values
	return self
}

func (self Entry) Level(level LogLevel) Entry {
	self.level = level
	return self
}

func (self Entry) Message(message string) Entry {
	self.message = message
	return self
}

//// Logger Aliases
func (self Logger) Log(level LogLevel, message string) Entry {
	return self.LogWithValues(level, message, make(map[string]interface{}))
}

func (self Logger) NewLog(level LogLevel, message string) Entry {
	return self.Log(level, message)
}

func (self Logger) LogWithValues(level LogLevel, message string, values map[string]interface{}) Entry {
	return Entry{
		logger:    &self,
		createdAt: time.Now(),
		level:     level,
		message:   message,
		format:    DefaultWithANSI,
		values:    values,
	}
}

func (self Logger) WithValues(values map[string]interface{}) Entry {
	return self.LogWithValues(LOG, "", values)
}

func (self Logger) Message(message string) Entry { return self.Log(LOG, message) }
func (self Logger) Level(level LogLevel) Entry   { return self.Log(level, "") }

//
// Level With Message Aliasing
///////////////////////////////////////////////////////////////////////////////
func (self Entry) Info(message string) Entry {
	self.message = message
	return self
}

func (self Entry) Notice(message string) Entry {
	self.message = message
	return self
}

func (self Entry) Warn(message string) Entry {
	self.message = message
	return self
}

func (self Entry) Warning(message string) Entry {
	self.message = message
	return self
}

func (self Entry) Error(err error) Entry {
	self.message = err.Error()
	self.errors = append(self.errors, err)
	return self
}

func (self Entry) Fatal(message string) Entry {
	self.message = message
	return self
}

func (self Entry) FatalError(err error) Entry {
	self.message = err.Error()
	self.errors = append(self.errors, err)
	return self
}

func (self Entry) Panic(message string) Entry {
	self.message = message
	return self
}
