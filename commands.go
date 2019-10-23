package cli

import (
	"strings"
)

type Command struct {
	Hidden      bool
	Category    int
	Name        string
	Alias       string
	Parent      *Command
	Subcommands []Command
	Flags       []Flag
	Description string
	Action      func(c *Context) error
}

// NOTE: Must cascade through parents recursively if the flag is missing. If no
// commands in the path have the flag defined, then it is ignored.
func (self *Command) Flag(name string) (*Flag, bool) {
	name = strings.ToLower(name)
	for _, flag := range self.Flags {
		if flag.Name == name || flag.Alias == name {
			return &flag, true
		}
		//if self.ParentCommand
	}
	return nil, false
}

func (self *Command) Subcommand(name string) (Command, bool) {
	for _, subcommand := range self.Subcommands {
		if subcommand.Name == name || subcommand.Alias == name {
			return subcommand, true
		}
	}
	return Command{}, false
}

func (self Command) is(name string) bool  { return self.Name == name || self.Alias == name }
func (self Command) visible() bool        { return !self.Hidden }
func (self Command) isRootCommand() bool  { return self.Parent == nil }
func (self Command) hasSubcommands() bool { return len(self.Subcommands) == 0 }

func (self Command) usage() (output string) {
	if len(self.Alias) > 0 {
		output += ", " + self.Alias
	}
	return self.Name + output
}

func (self Command) visibleSubcommands() (commands []Command) {
	for _, command := range self.Subcommands {
		if command.visible() {
			commands = append(commands, command)
		}
	}
	return commands
}

// Public Methods ////
func Commands(commands ...Command) []Command { return commands }
