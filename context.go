package cli

import (
	"fmt"
	"path/filepath"
	"strings"
)

// TODO: Maybe we should just call this process, and also include the PID, and
// helpers for kill, signal handling, child process spawning, resource
// information, and on-the-fly daemonization. This would enable the receiver to
// do a lot with the resulting context including completely manage the process.
// TODO: Should be able to provide configPath based on CLI.Name, as well as
// cache folder.
type Context struct {
	CLI         *CLI
	CWD         string
	Executable  string
	CommandPath []string
	Command     *inputCommand
	Flags       map[string]*inputFlag
	ParamType   DataType
	Params      []string
	Args        []string
}

// TODO: This is the object we are giving to the resulting action defined in the
// command.
// TODO: Don't save the defintion, just provide a function for looking up the
// definition
type inputCommand struct {
	//definition  *Command
	Name        string
	Flags       map[string]*inputFlag
	Parent      *inputCommand
	Subcommands []*inputCommand
}

func newInputCommand(parent *inputCommand, name string) *inputCommand {
	return &inputCommand{
		Name:        name,
		Parent:      parent,
		Flags:       make(map[string]*inputFlag),
		Subcommands: []*inputCommand{},
	}
}

func (self *inputCommand) addFlags(flags *inputFlags) {
	for _, flag := range *flags {
		self.Flags[flag.Name] = flag
	}
}

func (self *inputCommand) addSubcommandTree(subcommands []Command) {
	for _, subcommand := range subcommands {
		inputCmd := &inputCommand{Name: subcommand.Name, Parent: self}
		self.Subcommands = append(self.Subcommands, inputCmd)
		if !IsZero(len(subcommand.Subcommands)) {
			inputCmd.addSubcommandTree(subcommand.Subcommands)
		}
	}
}

type inputFlag struct {
	Name    string
	Type    DataType
	Command *inputCommand
	Value   string
}

func newInputFlag() *inputFlag {
	return &inputFlag{
		Value: "1",
	}
}

type inputFlags []*inputFlag

func newFlagGroup() *inputFlags                  { return &inputFlags{} }
func (self *inputFlags) addFlag(flag *inputFlag) { *self = append(*self, flag) }
func (self *inputFlags) reset()                  { self = &inputFlags{} }
func (self *inputFlags) isEmpty() bool           { return IsZero(len(*self)) }

func (self *Context) addCommand(command *inputCommand) *Context {
	fmt.Println("in add command func on context with inputCommand.Name: ", command.Name)
	if self.Command != nil {
		command.Parent = self.Command
		self.Command = newInputCommand(command.Parent, command.Name)
		self.CommandPath = append(self.CommandPath, self.Command.Name)
	} else if command.Parent != nil && !IsZero(len(command.Parent.Subcommands)) {
		// TODO: Maybe the validation for adding command should check the
		// command.Parent subcommands include the name. Maybe each group of flags
		// should be saved to context too, or saved to map with
		// dot.path.command.subcommand key
		for _, parentSubcommand := range command.Parent.Subcommands {
			fmt.Println("checking if parentSubcommand:[", parentSubcommand.Name, "] == command:[", command.Name, "]")
			if parentSubcommand.Name == command.Name {
				self.Command = newInputCommand(command.Parent, command.Name)
				self.CommandPath = append(self.CommandPath, self.Command.Name)
			}
		}
	}
	return self
}

func (self *Context) addFlag(flag *inputFlag) { self.Flags[flag.Name] = flag }

func (self *CLI) parse(arguments []string) *Context {
	cwd, executable := filepath.Split(arguments[0])
	context := &Context{
		CLI:        self,
		CWD:        cwd,
		Executable: executable,
		Command: &inputCommand{
			Name: self.Name,
		},
		Flags: map[string]*inputFlag{},
		Args:  arguments[1:],
	}
	// TODO: Currently we do not support -flag value OR -f value. It requires an
	// '='. But this needs to be fixed for both (even though its not convention
	// for long).

	// TODO:
	//  BUG: Command parsing
	//  1) If list and add are two global subcommands, then doing ./example list
	//  add accepts both, but will not take an extra one as in 'list add add'
	//  keeps only [example list add] path. **(But this may mean we are not
	//  allowing commands with the same name to exist in different scopes!!!)**
	flagGroup := newFlagGroup()
	for index, argument := range arguments {
		self.Debug("[context:parse()] attempting to parse the argument:[", argument, "]")
		if argument[0] == "-"[0] && len(argument) > 1 {
			if flag, ok := context.parseFlag(argument); ok {
				flagGroup.addFlag(flag)
			}
		} else {
			if command, ok := self.command.Route(append(context.CommandPath, argument)); ok {
				inputCmd := newInputCommand(context.Command, command.Name)
				inputCmd.addSubcommandTree(command.Subcommands)
				if !IsZero(len(*flagGroup)) {
					inputCmd.addFlags(flagGroup)
					flagGroup.reset()
				}
				context.addCommand(inputCmd)
			} else {
				for _, param := range arguments[index:] {
					if param[:1] == "-" {
						if flag, ok := context.parseFlag(param); ok {
							flagGroup.addFlag(flag)
						}
					} else {
						// TODO: Need parameter init datatype declaration to do more with
						// parameter otherwise its going to be simple string slice
						context.Params = append(context.Params, param)
					}
				}
				if !flagGroup.isEmpty() {
					context.Command.addFlags(flagGroup)
					flagGroup.reset()
				}
				break
			}
		}
	}
	return context
}

func (self *Context) parseFlag(argument string) (*inputFlag, bool) {
	parsed := newInputFlag()
	if argument[:1] == "--"[:1] {
		// Long Flag - convention is enforcing '=' on Long val
		parsed.Name = argument[2:]
		if strings.Contains(argument, "=") {
			// Not Bool Type
			flagParts := strings.Split(parsed.Name, "=")
			if IsGreaterThan(1, len(flagParts)) {
				parsed.Name = flagParts[0]
				parsed.Value = flagParts[1]
			}
		}
		if flag, ok := self.CLI.command.Flag(parsed.Name); ok {
			parsed.Name = flag.Name
			parsed.Type = flag.Type
			return parsed, true
		}
	} else {
		// Short Flag (or Alias) (ex. ls -a)
		for index, alias := range argument[1:] {
			// Stacked Short Flags (ex. `la -lah`)
			flagParts := strings.Split(string(alias), "=")
			if flag, ok := self.CLI.command.Flag(string(alias)); ok {
				parsed.Name = flag.Name
				if (len(argument[1:]) - 1) == index {
					if !IsZero(len(flagParts)) {
						parsed.Value = flagParts[1]
						parsed.Type = flag.Type
					}
				}
				return parsed, true
			}
		}
	}
	return nil, false
}
