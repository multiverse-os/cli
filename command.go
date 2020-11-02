package cli

import (
	"strings"
)

type Command struct {
	Category    int
	Name        string
	Alias       string
	Description string
	Hidden      bool
	Parent      *Command
	Subcommands []Command
	Flags       []Flag
	Action      Action
}

func (self Command) is(name string) bool { return self.Name == name || self.Alias == name }

func (self Command) visibleSubcommands() (subcommands []Command) {
	for _, subcommand := range self.Subcommands {
		if !subcommand.Hidden {
			subcommands = append(subcommands, subcommand)
		}
	}
	return subcommands
}

func (self Command) visibleFlags() (flags []*Flag) {
	for _, flag := range self.Flags {
		if !flag.Hidden {
			flags = append(flags, &flag)
		}
	}
	return flags
}

func (self Command) usage() (output string) {
	if len(self.Alias) != 0 {
		output += ", " + self.Alias
	}
	return self.Name + output
}

func (self Command) path() []string {
	route := []string{self.Name}
	for parent := self.Parent; parent != nil; parent = parent.Parent {
		route = append(route, parent.Name)
	}
	return route
}

//
// Public Methods
///////////////////////////////////////////////////////////////////////////////
func (self Command) Base() bool { return self.Parent == nil }

func (self Command) HasFlags() bool {
	return 0 < len(self.visibleFlags())
}

func Commands(commands ...Command) []Command { return commands }

func (self Command) Subcommand(name string) (Command, bool) {
	for _, subcommand := range self.Subcommands {
		if subcommand.is(name) {
			return subcommand, true
		}
	}
	return Command{}, false
}

func (self Command) Flag(arg string) (*Flag, bool) {
	arg = strings.ToLower(arg)
	for _, flag := range self.Flags {
		if flag.Name == arg || flag.Alias == arg {
			return &flag, true
		}
	}
	return nil, false
}

func (self Command) Path() []string {
	route := []string{self.Name}
	for parent := self.Parent; parent != nil; parent = parent.Parent {
		route = append(route, parent.Name)
	}
	return route
}
