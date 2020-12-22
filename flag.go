package cli

import (
	"strconv"
	"strings"

	data "github.com/multiverse-os/cli/data"
)

type FlagType int

const (
	Short FlagType = iota + 1
	Long
	NotAvailable
)

func (self FlagType) Is(flagType FlagType) bool { return self == flagType }
func (self FlagType) Length() int               { return int(self) }
func (self FlagType) String() string            { return strings.Repeat("-", self.Length()) }

type FlagSeparator int

const (
	Whitespace FlagSeparator = iota
	Equal
)

func (self FlagSeparator) Is(flagSeparator FlagSeparator) bool { return self == flagSeparator }

func (self FlagSeparator) String() string {
	if self.Is(Equal) {
		return "="
	} else {
		return " "
	}
}

type FlagLevel uint8

const (
	GlobalFlag FlagLevel = iota
	CommandFlag
)

// TODO: Be able to define the file extension that would be selected for when generating an autocomplete file
type Flag struct {
	Command     *Command
	Level       FlagLevel
	Name        string
	Alias       string
	Description string
	Hidden      bool
	Default     string
	Value       string
	Type        data.Type
}

func HasFlagPrefix(flag string) (FlagType, bool) {
	if strings.HasPrefix(flag, Long.String()) &&
		data.IsGreaterThan(len(flag), Long.Length()) {
		return Long, true
	} else if strings.HasPrefix(flag, Short.String()) &&
		data.IsGreaterThan(len(flag), Short.Length()) {
		return Short, true
	} else {
		return NotAvailable, false
	}
}

func (self Flag) is(name string) bool { return self.Name == name || self.Alias == name }

func (self Flag) flagNames() (output string) {
	return output
}

func (self Flag) help() string {
	usage := Long.String() + self.Name
	if data.NotBlank(self.Alias) {
		usage += ", " + Short.String() + self.Alias
	}
	var defaultValue string
	if len(self.Default) != 0 {
		defaultValue = " [≅ " + self.Default + "]"
	}
	return strings.Repeat(" ", 4) + usage + strings.Repeat(" ", 18-len(usage)) + self.Description + defaultValue + "\n"
}

func Flags(flags ...Flag) []Flag { return flags }

func (self Flag) String() string { return self.Value }

func (self Flag) Int() int {
	intValue, err := strconv.Atoi(self.Value)
	if err != nil {
		return 0
	} else {
		return intValue
	}
}

func (self Flag) Bool() bool { return data.IsTrue(self.Value) }
