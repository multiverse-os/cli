package cli

import (
	"fmt"
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

func (self *Chain) Last() *Command {
	return self.Commands[len(self.Commands)-1]
}

func (self *Chain) AddCommand(command *Command) {
	fmt.Println("adding command to chain")
	self.Commands = append(self.Commands, command)
}

func (self *Chain) Path() (path []string) {
	for _, command := range self.Commands {
		path = append(path, command.Name)
	}
	return path
}
