package cli

import (
  "strings"
)

type Command struct {
  Name        string
  Alias       string
  Description string
  Hidden      bool
  Parent      *Command
  Subcommands commands
  Flags       flags
  Action      Action
  Hooks       Hooks
}

func ValidateCommand(command Command) error {
  if 32 < len(command.Name) {
    return errInvalidArgumentLength
  }
  for _, commandRune := range command.Name {
    // NOTE: a = 97; z = 122; - = 45
    if (97 <= commandRune && commandRune <= 122) || commandRune == 45 {
      return errInvalidArgumentFormat
    }
  }
  return nil
}

func (self Command) IsValid() bool { return ValidateCommand(self) != nil }

func (self Command) is(name string) bool { 
  name = strings.ToLower(name)
  return self.Name == name || self.Alias == name
}

func (self Command) Subcommand(name string) (*Command, bool) {
  return self.Subcommands.Name(name)
}

func (self Command) Base() bool { return self.Parent == nil }
///////////////////////////////////////////////////////////////////////////////
// TODO: These should be consist with linked list even if we dont use it (and we
// probably should)
type commands []*Command

func Commands(commands ...Command) (commandPointers commands) { 
  for index, _ := range commands {
    commandPointers = append(commandPointers, &commands[index])
  }
  return commandPointers
}

func (self commands) Names() (commandNames []string) {
  for _, command := range self {
    commandNames = append(commandNames, command.Name)
  }
  return commandNames
}
  
// Commands Public Methods
func (self commands) First() *Command { return self[0] }
func (self commands) Last() *Command { return self[self.Count()-1] }

func (self commands) Count() int { return len(self) }
func (self commands) IsZero() bool { return self.Count() == 0 }

func (self commands) Name(name string) (*Command, bool) {
  for _, subcommand := range self {
    if subcommand.is(name) {
      return subcommand, true
    }
  }
  return nil, false
}

func (self commands) Reversed() (reversedCommands commands) {
  for i := self.Count() - 1; i >= 0; i-- {
    reversedCommands = append(reversedCommands, self[i])
  }
  return reversedCommands
}

func (self commands) Visible() (visibleCommands commands) {
  for _, command := range self {
    if !command.Hidden {
      visibleCommands = append(visibleCommands, command)
    }
  }
  return visibleCommands
}

func (self commands) Add(command *Command) commands  {
  command.Flags = command.Flags.SetDefaults()
  // TODO: For now add a reversed before returning, so we can have the newest
  // ones up front. Later we can switch it to container/list package or linked
  // list so we can append, prepend, etc
  //return append(self, command).Reversed()
  // To add this we need to remove 2 reverses from parse
  return append(self, command)
}
