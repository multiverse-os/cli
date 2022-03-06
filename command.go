package cli

import (
  "strings"

  data "github.com/multiverse-os/cli/data"
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

// Command Private Methods
func (self Command) is(name string) bool { 
  name = strings.ToLower(name)
  return self.Name == name || self.Alias == name
}

func (self Command) usage() (output string) {
  if data.IsBlank(self.Alias) {
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

func (self Command) Subcommand(name string) (*Command, bool) {
  return self.Subcommands.Name(name)
}

// Command Public Methods
func (self Command) Base() bool { return self.Parent == nil }

// TODO: These should be obsoleted by the Flags and Commands structures 
//func (self Command) HasFlags() bool { return 0 < len(self.VisibleFlags()) }
//func (self Command) HasSubcommands() bool { return 0 < len(self.VisibleSubcommands()) }

func (self *Command) HasFlag(name string) bool {
  return self.Flags.Name(name) != nil
}

// TODO: This NEEDs to be using the definedFlags in CLI to build the flag object
//       So it should be a flag passed to the command, not just the name and
//       value !!
func (self Command) AddFlag(flag *Flag) Command {
  self.Flags = append(self.Flags, flag)
  return self
}

// TODO: This was UpdateFlag so hopefully we ahve to fix something and this was
// not a deleterios function!?
func (self Command) SetFlag(name, value string) Command {
  flag, flagExists := self.Flag(name)
  if flagExists {
    flag.Set(value) // Param.Value = value
  }
  return self
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

func (self Command) HasNoAction() bool { return self.Action == nil }
func (self Command) HasAction() bool { return !self.HasNoAction() }

///////////////////////////////////////////////////////////////////////////////
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

// TODO: This isnt used atm, and probably wont be needed
//func (self commands) Hidden() (hiddenCommands commands) {
//  for _, command := range self {
//    if command.Hidden {
//      hiddenCommands = append(hiddenCommands, command)
//    }
//  }
//  return hiddenCommands
//}

func (self commands) Visible() (visibleCommands commands) {
  for _, command := range self {
    if !command.Hidden {
      visibleCommands = append(visibleCommands, command)
    }
  }
  return visibleCommands
}

func (self commands) Add(command *Command) commands  {
  for _, flag := range command.Flags {
    flag.Param = &Param{
      Value: flag.Default,
    }
  }
  return append(self, command)
}


