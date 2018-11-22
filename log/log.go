package log

import (
	"encoding/json"
	"fmt"
	"time"

	text "github.com/multiverse-os/cli/text"
)

type Errors []error
type Values map[string]string

type Entry struct {
	timestampResolution TimestampResolution
	logger              *Logger
	level               LogLevel
	createdAt           time.Time
	message             string
	values              Values
	errors              Errors
	format              Format
}

type Format int

const (
	DefaultFormat Format = iota
	StyledWithANSI
	JSON
	IndentedJSON
)

// Format Aliasing
const (
	PrettyJSON = IndentedJSON
)

// Formatting
///////////////////////////////////////////////////////////////////////////////
// Output caps a chained log entry create, it writes the log to the parameter
// defined output without requiring the 'Logger' object. This allows for quick
// and simple log systems that can be called without any prior initialization
// of any variables and with no global variables.
func (self Entry) Output() {
	if self.logger != nil {
		self.logger.Append(self)
	} else {
		self.Append()
	}
}

func (self Entry) Format(format Format) Entry {
	self.format = format
	return self
}

func (self Entry) MarshalJSON() (jsonOutput string, err error) {
	type entryJSON struct {
		Level     string    `json:"level"`
		CreatedAt time.Time `json:"created_at"`
		Message   string    `json:"message"`
		Values    Values    `json:"values,omitempty"`
		Errors    Errors    `json:"errors,omitempty"`
	}
	entry := entryJSON{
		Level:     self.level.String(),
		CreatedAt: self.createdAt,
		Message:   self.message,
		Values:    self.values,
		Errors:    self.errors,
	}
	var jsonBytes []byte
	switch self.format {
	case IndentedJSON:
		jsonBytes, err = json.MarshalIndent(entry, "", "  ")
		jsonOutput = string(jsonBytes)
	default:
		jsonBytes, err = json.Marshal(entry)
		jsonOutput = string(jsonBytes)
	}
	return jsonOutput, err
}

func (self Entry) DefaultFormat() Entry { return self.Format(DefaultFormat) }
func (self Entry) ANSI() Entry          { return self.Format(StyledWithANSI) }
func (self Entry) JSON() Entry          { return self.Format(JSON) }
func (self Entry) PrettyJSON() Entry    { return self.Format(PrettyJSON) }
func (self Entry) IndentedJSON() Entry  { return self.Format(IndentedJSON) }

func (self Entry) String() string {
	switch self.format {
	case DefaultFormat, StyledWithANSI:
		return (self.FormattedLevel() + self.FormattedTimestamp() + self.FormattedValues() + self.FormattedMessage())
	case JSON, PrettyJSON:
		var jsonOutput []byte
		var err error
		switch self.format {
		case PrettyJSON:
			jsonOutput, err = json.MarshalIndent(self, "", "  ")
		default: // JSON
			jsonOutput, err = json.Marshal(self)
		}
		if err != nil {
			FatalError(err)
		}
		return string(jsonOutput)
	}
	return ""
}

// Append (print or write to defined outputs)
///////////////////////////////////////////////////////////////////////////////
func (self Entry) Append() {
	if self.logger != nil {
		self.logger.Append(self)
	} else {
		self.StdOut()
	}
}

func (self Entry) File(path string) {
	logFile := LogFile{format: self.format, path: path}
	_ = logFile.Open()
	logFile.Append(self)
}

func (self Entry) Terminal() {
	terminal := Terminal{
		format: self.format,
	}
	terminal.Append(self)
}
func (self Entry) StdOut() { self.Terminal() }

// Chain-based Assignment
///////////////////////////////////////////////////////////////////////////////
// TODO: Should specifically Message and Values be validated and considered
// potentially dangerous user input?
func NewLog(format Format, level LogLevel, message string, values Values, errors []error) Entry {
	return Entry{
		createdAt: time.Now(),
		level:     level,
		format:    format,
		message:   message,
		values:    values,
		errors:    errors,
	}
}

func Log(level LogLevel, message string) Entry {
	return NewLog(StyledWithANSI, level, message, Values{}, []error{})
}

func Level(level LogLevel) Entry {
	return NewLog(StyledWithANSI, level, "", Values{}, []error{})
}

func Message(message string) Entry {
	return NewLog(StyledWithANSI, LOG, message, Values{}, []error{})
}

func LogWithError(level LogLevel, message string, err error) Entry {
	return NewLog(StyledWithANSI, level, message, Values{}, []error{err})
}

func LogWithErrors(level LogLevel, message string, errors []error) Entry {
	return NewLog(StyledWithANSI, level, message, Values{}, errors)
}

func (self Entry) WithValue(key string, v interface{}) Entry {
	switch value := v.(type) {
	case string:
		self.values[key] = value
	case int, uint:
		self.values[key] = fmt.Sprintf("%v", value)
	case bool:
		self.values[key] = fmt.Sprintf("%t", value)
	case float64, float32:
		self.values[key] = fmt.Sprintf("%0.4f", value)
	default:
		self.values[key] = fmt.Sprintf("%+v", value)
	}
	return self
}

func (self Entry) WithError(err error) Entry {
	return self.WithErrors([]error{err})
}

func (self Entry) WithErrors(errors []error) Entry {
	for _, e := range errors {
		self.errors = append(self.errors, e)
	}
	return self
}

// Level With Message Aliasing
///////////////////////////////////////////////////////////////////////////////
func Info(message string)    { Log(INFO, message).Append() }
func Debug(message string)   { Log(DEBUG, message).Append() }
func Notice(message string)  { Log(NOTICE, message).Append() }
func Warn(message string)    { Log(WARN, message).Append() }
func Warning(message string) { Log(WARNING, message).Append() }
func Error(err error)        { Log(ERROR, err.Error()).Append() }
func FatalError(err error)   { Log(FATAL, err.Error()).Append() }
func Fatal(message string)   { Log(FATAL, message).Append() }
func Panic(message string)   { Log(PANIC, message).Append() }

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

// Values and Errors
///////////////////////////////////////////////////////////////////////////////
func (self Entry) HasValues() bool    { return (len(self.values) != 0) }
func (self Entry) HasErrors() bool    { return (len(self.errors) != 0) }
func (self Entry) HasTimestamp() bool { return (self.timestampResolution == DISABLED) }

func (self Values) String() (valuesString string) {
	if len(self) > 0 {
		for key, value := range self {
			valuesString += key + "=" + value
		}
	}
	return valuesString
}

func (self Values) StyledString() (valuesString string) {
	if len(self) > 0 {
		for key, value := range self {
			valuesString += text.Bold(text.Green(key+"=")) + text.White(value)
		}
	}
	return valuesString
}

// Foramtted Output Strings
///////////////////////////////////////////////////////////////////////////////
// These functions use the defined format of an Entry to generate a formatted
// string for a given attribute that is displayed using calling .String()
// on an Entry object. This could allow for future functionality, like
// truncating length of message depending on
func (self Entry) FormattedMessage() (formattedLevel string) {
	return self.message
}

func (self Entry) FormattedLevel() (formattedLevel string) {
	switch self.format {
	case StyledWithANSI:
		return self.level.StyledString()
	default:
		return self.level.String()
	}
}

func (self Entry) FormattedTimestamp() (formattedTimestamp string) {
	formattedTimestamp = self.Timestamp()
	if formattedTimestamp != "" {
		if self.format == StyledWithANSI {
			text.Bold(formattedTimestamp)
		}
		formattedTimestamp = text.Brackets(formattedTimestamp)
	}
	return formattedTimestamp
}

func (self Entry) FormattedValues() (formattedValues string) {
	if self.HasValues() {
		switch self.format {
		case StyledWithANSI:
			formattedValues = self.values.StyledString()
		default:
			formattedValues = self.values.String()
		}
	}
	return formattedValues
}
