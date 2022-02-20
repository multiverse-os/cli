package cli

import (
	"strings"
)

type CommandActions struct {
  Before Action
  After  Action
}

type Command struct {
	Category    int
	Name        string
	Alias       string
	Description string
	Hidden      bool
	Parent      *Command
  Subcommands commands
	Flags       flags
	Action      Action
  Hooks       CommandActions
}

type commands []*Command


func Commands(definedCommands ...Command) (commandPointers commands) { 
  for _, command := range definedCommands {
    commandPointers = append(commandPointers, &command)
  }
  return commandPointers
}

// Commands Public Methods
func (self commands) Count() int { return len(self) }
func (self commands) IsZero() bool { return self.Count() == 0 }

// NOTE: This allows for Command.Subcommands.Visible()
func (self commands) Hidden() (commands commands) {
  for _, command := range self {
    if command.Hidden {
      commands = append(commands, command)
    }
  }
  return commands
}

func (self commands) Visible() (commands commands) {
  for _, command := range self {
    if !command.Hidden {
      commands = append(commands, command)
    }
  }
  return commands
}

// Command Private Methods
func (self Command) is(name string) bool { return self.Name == name || self.Alias == name }

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

// Command Public Methods
func (self Command) Base() bool { return self.Parent == nil }

// TODO: These should be obsoleted by the Flags and Commands structures 
//func (self Command) HasFlags() bool { return 0 < len(self.VisibleFlags()) }
//func (self Command) HasSubcommands() bool { return 0 < len(self.VisibleSubcommands()) }

// TODO: This already exists in context, but should be taken out. Context should
// be for meta logic that combines the individaul or list logic into more
// complex helpers. Tasks like looking up subcommands should be built in by
// nature of the fact that even the base command is technically a command and
// they nest infinitely.
func (self Command) Subcommand(name string) (*Command, bool) {
	for _, subcommand := range self.Subcommands {
		if subcommand.is(name) {
			return subcommand, true
		}
	}
	return nil, false
}

func (self Command) Flag(name string) (*Flag, bool) {
	for _, flag := range self.Flags {
		if flag.is(strings.ToLower(name)) {
			return flag, true
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

// TODO: This can now be accomplished with 
//
//                  command.Subcommands.Visible()
//
//       which is definitely a much nicer API for developers to interact
//       with. Commend this out and kinda keep it a bit to provide examples
//       for this new trick we are bringing into the standard object build. 
//
// func (self Command) VisibleSubcommands() (subcommands []Command) {
// 	for _, subcommand := range self.Subcommands {
// 		if !subcommand.Hidden {
// 			subcommands = append(subcommands, subcommand)
// 		}
// 	}
// 	return subcommands
// }

// TODO: This should be obsoleted by 
//func (self Command) VisibleFlags() (flags []*Flag) {
//	for _, flag := range self.Flags {
//		if !flag.Hidden {
//			flags = append(flags, &flag)
//		}
//	}
//	return flags
//}

func (self Command) HasNoAction() bool { return self.Action == nil }
func (self Command) HasAction() bool { return !self.HasNoAction() }

