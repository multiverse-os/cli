package log

import (
	"encoding/json"
	"fmt"
	"time"

	color "github.com/multiverse-os/cli/text/ansi/color"
	style "github.com/multiverse-os/cli/text/ansi/style"
)

type Errors []error
type Values map[string]string

type Entry struct {
	format              Format
	timestampResolution TimestampResolution
	logger              *Logger
	level               LogLevel
	createdAt           time.Time
	message             string
	values              Values
	errors              Errors
}

type Format int

const (
	Default Format = iota
	ANSI
	JSON
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
	jsonBytes, err = json.Marshal(entry)
	jsonOutput = string(jsonBytes)
	return jsonOutput, err
}

func (self Entry) Default() Entry { return self.Format(Default) }
func (self Entry) ANSI() Entry    { return self.Format(ANSI) }
func (self Entry) JSON() Entry    { return self.Format(JSON) }

func (self Entry) String() string {
	var values string
	var timestamp string
	switch self.format {
	case ANSI:
		count := 0
		for key, value := range self.values {
			values += (valueStringWithANSI(count, style.Bold(color.Blue(key+"="))) + style.Dim(color.White(value)))
			count++
		}
		if self.HasTimestamp() {
			fmt.Println("Has timestamp, resolution is:", self.timestampResolution)
			timestamp = " " + "[ " + style.Bold(self.Timestamp()) + " ]"
		}
		return self.level.StringWithANSI() + timestamp + " " + values + " " + color.White(self.message)
	case JSON:
		var jsonOutput []byte
		var err error
		jsonOutput, err = json.Marshal(self)
		if err != nil {
			Fatal(err.Error())
		}
		return string(jsonOutput)
	default:
		if self.HasTimestamp() {
			fmt.Println("Has timestamp, resolution is:", self.timestampResolution)
			timestamp = " " + "[ " + self.Timestamp() + " ]"
		}
		for key, value := range self.values {
			values += (key + "=" + value)
		}
		return self.level.String() + timestamp + " " + values + " " + self.message
	}
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
		createdAt:           time.Now(),
		timestampResolution: DISABLED,
		level:               level,
		format:              format,
		message:             message,
		values:              values,
		errors:              errors,
	}
}

func Log(level LogLevel, message string) Entry {
	return NewLog(Default, level, message, Values{}, []error{})
}

func LogWithANSI(level LogLevel, message string) Entry {
	return NewLog(ANSI, level, message, Values{}, []error{})
}

func JSONLog(level LogLevel, message string) Entry {
	return NewLog(JSON, level, message, Values{}, []error{})
}

func Level(level LogLevel) Entry {
	return NewLog(Default, level, "", Values{}, []error{})
}

func Message(message string) Entry {
	return NewLog(Default, LOG, message, Values{}, []error{})
}

func LogWithError(level LogLevel, message string, err error) Entry {
	return NewLog(Default, level, message, Values{}, []error{err})
}

func LogWithErrors(level LogLevel, message string, errors []error) Entry {
	return NewLog(Default, level, message, Values{}, errors)
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
func Info(message string)   { Log(INFO, message).Append() }
func Debug(message string)  { Log(DEBUG, message).Append() }
func Notice(message string) { Log(NOTICE, message).Append() }
func Warn(message string)   { Log(WARN, message).Append() }
func Error(err error)       { Log(ERROR, err.Error()).Append() }
func Fatal(message string)  { Log(FATAL, message).Append() }
func Panic(message string)  { Log(PANIC, message).Append() }

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

func (self Entry) Error(err error) Entry {
	self.message = err.Error()
	self.errors = append(self.errors, err)
	return self
}

func (self Entry) Fatal(message string) Entry {
	self.message = message
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
func (self Entry) HasTimestamp() bool { return (self.timestampResolution != DISABLED) }

func valueStringWithANSI(count int, key string) string {
	switch count {
	case 0:
		return style.Bold(color.Blue(key))
	case 1:
		return style.Bold(color.Green(key))
	case 2:
		return style.Bold(color.Fuchsia(key))
	case 3:
		return style.Bold(color.Yellow(key))
	case 4:
		return style.Bold(color.Cyan(key))
	default:
		return style.Bold(color.Silver(key))
	}
}
