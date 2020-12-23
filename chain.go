package cli

import (
	"strings"
)

type Chain struct {
	Commands []*Command
}

func (self *Chain) Route(path []string) (*Command, bool) {
	cmd := &Command{}
	for index, command := range self.Commands {
		if command.Name == path[index] {
			if index == len(path) {
				return command, true
			} else {
				cmd = command
			}
		} else {
			return cmd, (len(cmd.Name) == 0)
		}
	}
	return nil, (len(cmd.Name) == 0)
}

func (self *Chain) First() *Command {
	if 0 < len(self.Commands) {
		return self.Commands[0]
	} else {
		return nil
	}
}

func (self *Chain) AddCommand(command *Command) {
	self.Commands = append(self.Commands, command)

	flags := []Flag{}
	for _, flag := range command.Flags {
		//if len(flag.Value) == 0 {
		flag.Value = flag.Default
		//}
		flags = append(flags, flag)
	}
	command.Flags = flags
}

func (self *Chain) Last() *Command             { return self.Commands[len(self.Commands)-1] }
func (self *Chain) NoCommands() bool           { return self.IsRoot() && len(self.First().Subcommands) == 0 }
func (self *Chain) HasCommands() bool          { return self.IsRoot() && 0 < len(self.First().Subcommands) }
func (self *Chain) IsRoot() bool               { return len(self.Commands) == 1 }
func (self *Chain) IsNotRoot() bool            { return 1 < len(self.Commands) }
func (self *Chain) UnselectedCommand() bool    { return 0 < len(self.Last().Subcommands) }
func (self *Chain) PathExample() (path string) { return strings.Join(self.Path(), " ") }

func (self *Chain) HasSubcommands() bool {
	return self.IsNotRoot() && (0 < len(self.Last().Subcommands))
}

func (self *Chain) Flags() (flags map[string]*Flag) {
	for _, command := range self.Commands {
		for _, flag := range command.Flags {
			flags[flag.Name] = &flag
		}
	}
	return flags
}

func (self *Chain) Path() (path []string) {
	for _, command := range self.Commands {
		path = append(path, command.Name)
	}
	return path
}

func (self *Chain) Reversed() (commands []*Command) {
	for i := len(self.Commands) - 1; i >= 0; i-- {
		commands = append(commands, self.Commands[i])
	}
	return commands
}

func (self *Chain) ReversedPath() (path []string) {
	for i := len(self.Commands) - 1; i >= 0; i-- {
		path = append(path, self.Commands[i].Name)
	}
	return path
}
