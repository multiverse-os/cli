package cli

import (
	"errors"
)

// Experimental
type Argument struct {
	Raw          string
	Value        interface{}
	IsCommand    bool
	IsSubcommand bool
	IsFlag       bool
	DataType     string
}

func (self Argument) Parse() {
	// TODO: Take raw string value and determine if its command, subcommand, flag, datatype and insert the value
}

/// End Experiment ////

type Arguments []string

func (self *Context) Arguments() Arguments {
	return (Arguments(self.flagSet.Arguments()))
}

func (self Arguments) HasArguments() bool {
	return (len(Arguments) > 0)
}

func (self Arguments) AtIndex(index int) string {
	if len(self) > index {
		return self[index]
	}
	return ""
}

func (self Arguments) First() string {
	return self.AtIndex(0)
}

func (self Arguments) Second() string {
	return self.AtIndex(1)
}

func (self Arguments) Third() string {
	return self.AtIndex(2)
}

func (self Arguments) Fourth() string {
	return self.AtIndex(3)
}

func (self Arguments) Last() string {
	return self.AtIndex(len(self) - 1)
}

func (self Arguments) Tail() []string {
	if len(self) > 1 {
		return self[1:]
	}
	return []string{}
}

func (self Arguments) Swap(a, b int) error {
	if a >= len(self) || b >= len(self) {
		return errors.New("index out of range")
	}
	self[b], self[a] = self[a], self[b]
	return nil
}
