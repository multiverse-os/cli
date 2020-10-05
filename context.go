package cli

import (
	"strings"

	argument "./framework/argument"
	data "./framework/argument/data"
	token "./framework/argument/token"
)

type Context struct {
	PID          int
	CLI          *CLI
	CWD          string
	Executable   string
	Command      *argument.Command
	Flags        map[string]*argument.Flag
	Params       argument.Params
	CommandChain *argument.Chain
	Args         []string
}

func (self *Context) HasFlag(name string) bool {
	_, ok := self.Flags[name]
	return ok
}

func (self *Context) CommandDefinition() *Command {
	return self.Command.Definition.(*Command)
}

func (self *Context) ParseFlag(index int, flagType token.Identifier, flag *argument.Flag) {
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
		for index, stackedFlag := range shortName {
			// Load flag
			if command, flagDefinition, ok := self.CLI.IsFlag(self.Command.Path(), string(stackedFlag)); ok {
				self.CLI.Log(DEBUG, "flag.Name:", flag.Name)
				self.CLI.Log(DEBUG, "flag.Command.Name:", command.Name)
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
		flag.Name = flagParts[0][1:]
	}
	self.CLI.Log(DEBUG, "Adding flag to context")
	if 0 < len(flag.Name) {
		self.CLI.Log(DEBUG, "Looking up flag to determine what level it is in with path:", self.Command.Path())
		//if command, _, ok := self.CLI.IsFlag(self.Command.Path(), flag.Name); ok {
		//		self.AddFlag(command, flag)
		self.Flags[flag.Name] = flag
		//} else {
		self.CLI.Log(DEBUG, "Failed to find command with flag")
		//	}
	} else {
		self.CLI.Log(DEBUG, "Not addding flag because flag name is lenth 0")
	}
}
