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
}

// TODO: Make sure that no existing commands ahve the same name, and since
// delete uses name
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

func (self Command) Subcommand(name string) *Command {
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

func (self commands) Arguments() (commandArguments arguments) {
  for _, command := range self {
    commandArguments = append(commandArguments, Argument(command))
  }
  return commandArguments
}

func (self commands) Names() (commandNames []string) {
  for _, command := range self {
    commandNames = append(commandNames, command.Name)
  }
  return commandNames
}
  
// Commands Public Methods
func HelpCommand(context *Context) error {
  context.Commands = context.Commands.Delete(context.Commands.Name("help"))
  return RenderDefaultHelpTemplate(context)
}

func (self commands) First() *Command { return self[0] }
func (self commands) Last() *Command { return self[self.Count()-1] }

func (self commands) Count() int { return len(self) }
func (self commands) IsZero() bool { return self.Count() == 0 }

func (self commands) HasCommand(name string) bool { 
  return self.Name(name) != nil
}

func (self commands) Name(name string) *Command {
  for _, subcommand := range self {
    if subcommand.is(name) {
      return subcommand
    }
  }
  return nil
}

func (self commands) Delete(command *Command) (newCommands commands) {
  for index, _ := range self {
      if self[index].Name != command.Name {
        newCommands = newCommands.Add(self[index])
      }
  }

  return newCommands
}
 

func (self commands) Visible() (visibleCommands commands) {
  for _, command := range self {
    if !command.Hidden {
      visibleCommands = append(visibleCommands, command)
    }
  }
  return visibleCommands
}

func (self commands) Reverse() (reversedCommands commands) {
  for reversedIndex := self.Count() - 1; reversedIndex >= 0; reversedIndex-- {
    reversedCommands = append(reversedCommands, self[reversedIndex])
  }
  return reversedCommands
}

func (self commands) Add(command *Command) (updatedCommands commands) { 
  command.Flags = command.Flags.SetDefaults()
  return append(append(updatedCommands, command), self...)
}
