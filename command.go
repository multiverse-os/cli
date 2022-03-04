package cli

import (
  "strings"
)

type Command struct {
  IsRoot      bool
  Category    int
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
    // NOTE: 
    // a = 97
    // z = 122
    // - = 45
    if (97 <= commandRune && commandRune <= 122) || commandRune == 45 {
      return errInvalidArgumentFormat
    }
  }
  return nil
}

func (self Command) IsValid() bool { return ValidateCommand(self) != nil }

func (self Command) Type() ArgumentType { return CommandArgument }

type commands []*Command

func Commands(commands ...Command) (commandPointers commands) { 
  for index, _ := range commands {
    commandPointers = append(commandPointers, &commands[index])
  }
  return commandPointers
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

func (self commands) Reversed() (commands commands) {
  for i := self.Count() - 1; i >= 0; i-- {
    commands = append(commands, self[i])
  }
  return commands
}

func (self commands) Path() (path []string) {
  for _, command := range self {
    path = append(path, command.Name)
  }
  return path
}

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

func (self commands) Add(command *Command) commands  {
  for _, flag := range command.Flags {
    flag.Param = &Param{
      Value: flag.Default,
    }
  }

  return append(self, command)
}

// Command Private Methods
func (self Command) is(name string) bool { 
  name = strings.ToLower(name)
  return self.Name == name || self.Alias == name
}

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


//func (self *Command) Flags() (flags map[string]*Flag) {
//	for _, flag := range self.Flags {
//		flags[flag.Name] = flag
//	}
//	return flags
//}
// c.Flags["port"].Int()

// c.Flags.Name("port").Int() 

// c.Flag("port").Int()

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

