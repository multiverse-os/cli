package cli

import (
	"strings"
)

type Command struct {
	Global      bool
	Category    int
	Name        string
	Alias       string
	Description string
	Hidden      bool
	Parent      *Command // TODO: Considier reducing this by replacing this dot.path command.subcommand
	Subcommands []Command
	Flags       []Flag
	Action      Action
}

// TODO Consider doing murmur has for comparisons, to speed up everything.

// TODO: Be able to walk over the command tree to output all items, for putting into a map, for hashing and creating ids, etc
// this was already written just look a few commits back

// TODO: Support printing command tree
// TODO: These may be able to go back to being private, were made public when we experimented with having their own package
func (self Command) is(name string) bool { return self.Name == name || self.Alias == name }

func (self Command) visibleSubcommands() (subcommands []Command) {
	for _, subcommand := range self.Subcommands {
		if !subcommand.Hidden {
			subcommands = append(subcommands, subcommand)
		}
	}
	return subcommands
}

func (self Command) visibleFlags() (flags []Flag) {
	for _, flag := range self.Flags {
		if !flag.Hidden {
			flags = append(flags, flag)
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
func Commands(commands ...Command) []Command { return commands }

func (self Command) Subcommand(name string) (Command, bool) {
	for _, subcommand := range self.Subcommands {
		if subcommand.is(name) {
			return subcommand, true
		}
	}
	return Command{}, false
}

// NOTE: Must cascade through parents recursively if the flag is missing. If no
// commands in the path have the flag defined, then it is ignored.
func (self Command) Flag(arg string) (*Command, *Flag, bool) {
	arg = strings.ToLower(arg)
	for _, flag := range self.Flags {
		if flag.Name == arg || flag.Alias == arg {
			return &self, &flag, true
		}
		//if self.ParentCommand
	}
	return nil, nil, false
}

func (self Command) Path() []string {
	route := []string{self.Name}
	for parent := self.Parent; parent != nil; parent = parent.Parent {
		route = append(route, parent.Name)
	}
	return route
}
