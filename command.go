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
// TODO: Will need to be from CLI to it can get access to the context in order
// to validate uniquness of command but that may just be a waste for checking
// against bad developer knowledge using the library
func ValidateCommand(command Command) error {
	if 32 < len(command.Name) {
		return ErrInvalidArgumentLength
	}
	//for _, commandRune := range command.Name {
	//  // NOTE: a = 97; z = 122; - = 45
	//  if (97 <= commandRune && commandRune <= 122) || commandRune == 45 {
	//    return ErrInvalidArgumentFormat
	//  }
	//}
	return nil
}

func (cmd Command) IsValid() bool { return ValidateCommand(cmd) != nil }

func (cmd Command) is(name string) bool {
	name = strings.ToLower(name)
	return (len(cmd.Name) == len(name) && cmd.Name == name) ||
		(len(cmd.Alias) == len(name) && cmd.Alias == name)
}

func (cmd Command) Subcommand(name string) *Command {
	return cmd.Subcommands.Name(name)
}

func (cmd Command) Flag(name string) *Flag {
	return cmd.Flags.Name(name)
}

func (cmd Command) IsRoot() bool { return cmd.Parent == nil }

// /////////////////////////////////////////////////////////////////////////////
// TODO: These should be consist with linked list even if we dont use it (and we
// probably should)
type commands []*Command

func Commands(commands ...Command) (commandPointers commands) {
	for index, _ := range commands {
		commandPointers = append(commandPointers, &commands[index])
	}
	return commandPointers
}

func (cmds commands) Arguments() (commandArguments arguments) {
	for _, command := range cmds {
		commandArguments = append(commandArguments, Argument(command))
	}
	return commandArguments
}

func (cmds commands) Names() (commandNames []string) {
	for _, command := range cmds {
		commandNames = append(commandNames, command.Name)
	}
	return commandNames
}

// Commands Public Methods
func HelpCommand(context *Context) error {
	context.Commands = context.Commands.Remove("help")
	//context.Commands = context.Commands.Reverse()
	return RenderDefaultHelpTemplate(context)
}

func VersionCommand(context *Context) error {
	return RenderDefaultVersionTemplate(context)
}

func (cmds commands) First() *Command { return cmds[0] }
func (cmds commands) Last() *Command  { return cmds[cmds.Count()-1] }
func (cmds commands) Count() int      { return len(cmds) }
func (cmds commands) IsZero() bool    { return cmds.Count() == 0 }

// TODO: Exists() should just be
func (cmds commands) HasCommand(name string) bool {
	return cmds.Name(name) != nil
}

// TODO: Do we need this aliasing?
func (cmds commands) Exists(name string) bool { return cmds.HasCommand(name) }

func (cmds commands) Name(name string) *Command {
	for _, subcommand := range cmds {
		if subcommand.is(name) {
			return subcommand
		}
	}
	return nil
}

func (cmds commands) Index(name string) int {
	for index, subcommand := range cmds {
		if subcommand.is(name) {
			return index
		}
	}
	return -1
}

func (cmds commands) Remove(name string) (newCommands commands) {
	for _, subcommand := range cmds {
		if !subcommand.is(name) {
			newCommands = append(newCommands, subcommand)
		}
	}
	return newCommands
}

func (cmds commands) Validate() (errs []error) {
	for _, command := range cmds {
		if err := ValidateCommand(*command); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func (cmds commands) Visible() (visibleCommands commands) {
	for _, command := range cmds {
		if !command.Hidden {
			visibleCommands = append(visibleCommands, command)
		}
	}
	return visibleCommands
}

func (cmds commands) Reverse() (reversedCommands commands) {
	for reversedIndex := cmds.Count() - 1; reversedIndex >= 0; reversedIndex-- {
		reversedCommands = append(reversedCommands, cmds[reversedIndex])
	}
	return reversedCommands
}

func (cmds *commands) Add(command *Command) {
	command.Flags = command.Flags.SetDefaults()
	*cmds = append(append(commands{}, command), *cmds...)
}
