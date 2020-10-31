package cli

import (
	"fmt"
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
	GlobalFlags  map[string]*Flag
	CommandFlags map[string]map[string]*Flag
	Params       Params
	CommandChain *Chain
	Args         []string
}

func (self *Context) HasFlag(name string) bool {
	return (self.HasGlobalFlag(name) || self.HasCommandFlag(name))
}

func (self *Context) HasGlobalFlag(name string) bool {
	_, ok := self.GlobalFlags[name]
	return ok
}

func (self *Context) HasCommandFlag(name string) bool {
	for _, command := range self.CommandChain.Commands {
		for _, flag := range self.CommandFlags[command.Name] {
			if name == flag.Name {
				return true
			}
		}
	}
	return false
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
		flag.Name = flagParts[0][2:]
	}
	self.CLI.Log(DEBUG, "Adding flag to context")
	if 0 < len(flag.Name) {
		self.CLI.Log(DEBUG, "flag.Name:", flag.Name)
		self.CLI.Log(DEBUG, "flag.Value:", flag.Value)
		self.CLI.Log(DEBUG, "Looking up flag to determine what level it is in with path:", self.Command.Path())
		//if command, _, ok := self.CLI.IsFlag(self.Command.Path(), flag.Name); ok {
		//		self.AddFlag(command, flag)

		fmt.Println("self.Command.Name:", self.Command.Name)
		fmt.Println("flag.Name:", flag.Name)
		fmt.Println("len(self.CommandFlags):", len(self.CommandFlags))
		if len(self.CommandFlags[self.Command.Name]) == 0 {
			self.CommandFlags[self.Command.Name] = make(map[string]*Flag)
		}

		fmt.Println("len(self.Command.Name][flag.Name]):", len(self.CommandFlags[self.Command.Name]))

		self.CommandFlags[self.Command.Name][flag.Name] = flag
		//} else {
		self.CLI.Log(DEBUG, "Failed to find command with flag")
		//	}
	} else {
		self.CLI.Log(DEBUG, "Not addding flag because flag name is lenth 0")
	}
}
