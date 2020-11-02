package cli

import (
	"strings"

	data "github.com/multiverse-os/cli/data"
	token "github.com/multiverse-os/cli/token"
)

type Context struct {
	PID          int
	CLI          *CLI
	CWD          string
	Executable   string
	Command      *Command
	Flags        map[string]map[string]*Flag
	Params       Params
	CommandChain *Chain
	Args         []string
}

func (self *Context) HasCommandFlag(name string) bool {
	return self.CommandFlag(self.Command.Name, name) != nil
}

func (self *Context) HasFlag(name string) bool {
	return self.Flag(name) != nil
}

func (self *Context) HasGlobalFlag(name string) bool {
	return self.GlobalFlag(name) != nil
}

func (self *Context) HasCommands() bool {
	return self.CommandChain.IsRoot() && self.HasSubcommands()
}

func (self *Context) HasSubcommands() bool {
	return 0 < len(self.Command.Subcommands)
}

func (self *Context) HasSubcommand(name string) bool {
	_, hasSubcommand := self.Subcommand(name)
	return hasSubcommand
}

func (self *Context) Subcommand(name string) (*Command, bool) {
	for _, subcommand := range self.Command.Subcommands {
		if subcommand.is(name) {
			return &subcommand, true
		}
	}
	return nil, false
}

func (self *Context) HasGlobalFlags() bool {
	return 0 < len(self.GlobalFlags())
}

func (self *Context) GlobalFlag(name string) (flag *Flag) {
	for _, flag := range self.GlobalFlags() {
		if flag.is(name) {
			return flag
		}
	}
	return flag
}

func (self *Context) Flag(name string) (flag *Flag) {
	if self.HasCommands() {
		for _, command := range self.CommandChain.Reversed().Commands {
			for _, flag := range self.Flags[command.Name] {
				if flag.is(name) {
					return flag
				}
			}
		}
	}
	return flag
}

func (self *Context) CommandFlag(command, name string) (flag *Flag) {
	for _, flag := range self.Flags[command] {
		if flag.is(name) {
			return flag
		}
	}
	return flag
}

func (self *Context) GlobalFlags() map[string]*Flag {
	flags := make(map[string]*Flag)
	for _, flag := range self.CommandChain.First().Flags {
		flags[flag.Name] = &flag
	}
	return flags
}

func (self *Context) ParseFlag(index int, flagType token.Identifier, flag *Flag) {
	var flagParts []string
	flagParts = strings.Split(flag.Name, token.Equal.String())
	if 1 < len(flagParts) {
		flag.Value = flagParts[1]
	} else {
		if len(self.Args) > index+1 {
			flag.Value = self.Args[index+1]
		} else {
			flag.Value = "1"
			flag.Type = data.Bool
		}
	}
	if flagType == token.Short {
		shortName := flagParts[0][1:]
		// Stacked Flags
		// TODO: Needs to work from specific to global so may need a for loop
		// with minus index i--
		for index, stackedFlag := range shortName {
			// Load flag
			if flagDefinition := self.Flag(string(stackedFlag)); flagDefinition != nil {
				flag.Name = flagDefinition.Name
				if index != (len(flag.Name) - 1) {
					// NOTE: Stacked flag that is not the last element MUST be bool
					flag.Value = "1"
					flag.Type = data.Bool
				} else {
					// NOTE: Stacked flag that is last element needs to use value

				}
				self.ParseFlag(index, flagType, flag)
			}
		}
	} else if flagType == token.Long {
		flag.Name = flagParts[0][2:]
	}
	if 0 < len(flag.Name) {
		if len(self.Flags[self.Command.Name]) == 0 {
			self.Flags[self.Command.Name] = make(map[string]*Flag)
		}
		self.Flags[self.Command.Name][flag.Name] = flag
	}
}
