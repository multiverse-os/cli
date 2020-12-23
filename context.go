package cli

import "fmt"

type Context struct {
	PID          int
	CLI          *CLI
	CWD          string
	Executable   string
	Command      *Command
	Params       Params
	CommandChain *Chain
	Flags        map[string]*Flag
	Args         []string
}

func (self *Context) UpdateFlag(name, value string) {
	if self.HasCommands() {
		for _, command := range self.CommandChain.Reversed() {
			for _, flag := range command.Flags {
				if flag.is(name) {
					flag.Value = value
				}
			}
		}
	}
}

func (self *Context) HasFlag(name string) bool { return self.Flag(name) != nil }
func (self *Context) HasCommandFlag(name string) bool {
	return self.CommandFlag(self.Command.Name, name) != nil
}

func (self *Context) Flag(name string) *Flag {
	for _, command := range self.CommandChain.Reversed() {
		fmt.Println("command flags:", len(command.Flags))
		for _, flag := range command.Flags {
			if flag.is(name) {
				if len(flag.Value) == 0 {
					flag.Value = flag.Default
				}
				return &flag
			}
		}
	}
	return &Flag{
		Name:  name,
		Value: "0",
	}
}

func (self *Context) CommandFlag(commandName, flagName string) (flag *Flag) {
	for _, command := range self.CommandChain.Commands {
		if command.is(commandName) {
			for _, flag := range command.Flags {
				if flag.is(flagName) {
					if len(flag.Value) == 0 {
						flag.Value = flag.Default
					}
					return &flag
				}
			}
		}
	}
	return flag
}

func (self *Context) HasCommands() bool    { return self.CommandChain.IsRoot() && self.HasSubcommands() }
func (self *Context) HasSubcommands() bool { return 0 < len(self.Command.Subcommands) }

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
