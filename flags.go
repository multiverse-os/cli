package cli

import (
	"errors"
	"os"
	"strconv"
	"strings"

	style "github.com/multiverse-os/cli/framework/terminal/ansi/style"
)

var True = []string{"t", "true", "y", "yes", "1"}

type FlagType int

const (
	BoolFlag FlagType = iota
	IntFlag
	StringFlag
	DirectoryFlag
	FileFlag
)

type Flag struct {
	Name        string // Primary name
	Aliases     []string
	Description string
	Hidden      bool
	Value       interface{}
}

func (self Flag) Help() string {
	return "    " +
		style.Bold(self.Usage()) +
		strings.Repeat(" ", (18-len(self.Usage()))) +
		style.Dim(self.Description) +
		"\n"
}

func (self Flag) Valid(flagType FlagType) (bool, error) {
	switch flagType {
	//case BoolFlag:
	//case IntFlag:
	//case StringFlag:
	case DirectoryFlag, FileFlag:
		_, err := os.Stat(self.Value.(string))
		return (err == nil), nil
	default:
		return false, errors.New("[error] flag does not have defined type")
	}
}

func (self Flag) String() string   { return self.Value.(string) }
func (self Flag) Path() string     { return self.Value.(string) }
func (self Flag) Filename() string { return self.Value.(string) }

func (flag Flag) Names() []string { return append([]string{flag.Name}, flag.Aliases...) }
func (self Flag) Visible() bool   { return !self.Hidden }

func (self Flag) Float() float64 {
	floatParts := strings.Split(self.Value.(string), ".")
	if len(floatParts) > 1 {
		floatValue, err := strconv.ParseFloat(self.Value.(string), len(floatParts[1]))
		if err != nil {
			return float64(0.00)
		} else {
			return floatValue
		}
	}
	return float64(self.Int())
}

func (self Flag) Bool() bool {
	for _, value := range True {
		if value == self.Value.(string) {
			return true
		}
	}
	return false
}

func (self Flag) Int() int {
	value, err := strconv.Atoi(self.Value.(string))
	if err != nil {
		return 0
	} else {
		return value
	}
}

func (self Flag) Usage() (output string) {
	if len(self.Aliases) > 0 {
		if len(self.Aliases[0]) >= 2 {
			output += "--" + self.Aliases[0]
		} else {
			output += "-" + self.Aliases[0]
		}
	}
	output += ", --" + self.Name
	return output
}

func (self Flag) Is(name string) bool {
	for _, flagName := range self.Names() {
		if flagName == name {
			return true
		}
	}
	return false
}

func defaultFlags() []Flag {
	return []Flag{
		Flag{
			Name:        "version",
			Aliases:     []string{"v"},
			Description: "Print version",
			Hidden:      false,
		},
		Flag{
			Name:        "help",
			Aliases:     []string{"h"},
			Description: "Print help text",
			Hidden:      false,
		},
	}
}

func defaultCommandFlags() []Flag {
	return []Flag{
		Flag{
			Name:        "help",
			Aliases:     []string{"h"},
			Description: "Print help text",
			Hidden:      false,
		},
	}
}
