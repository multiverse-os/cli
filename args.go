package cli

import (
	"errors"
)

type Argument struct {
	Raw                string
	Value              interface{}
	CommandName        string
	CommandTree        [][]string
	SubcommandName     string
	CommandBreadcrumbs []string
	IsCommand          bool
	IsSubcommand       bool
	IsFlag             bool
	DataType           string
}

func (self Argument) Parse() {
	// TODO: Take raw string value and determine if its command, subcommand, flag, datatype and insert the value
}

type Args []string

func (self Args) HasArgs() bool {
	return (len(self) > 0)
}

func (self Args) AtIndex(index int) string {
	if len(self) > index {
		return self[index]
	}
	return ""
}

func (self Args) First() string {
	return self.AtIndex(0)
}

func (self Args) Second() string {
	return self.AtIndex(1)
}

func (self Args) Third() string {
	return self.AtIndex(2)
}

func (self Args) Fourth() string {
	return self.AtIndex(3)
}

func (self Args) Last() string {
	return self.AtIndex(len(self) - 1)
}

func (self Args) Tail() []string {
	if len(self) > 1 {
		return self[1:]
	}
	return []string{}
}

func (self Args) Swap(a, b int) error {
	if a >= len(self) || b >= len(self) {
		return errors.New("index out of range")
	}
	self[b], self[a] = self[a], self[b]
	return nil
}
