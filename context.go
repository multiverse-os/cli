package cli

import (
	"strings"
)

type CLIAction int

// TODO: The idea here is to establish a router like model to make CLI
// application development more like web development, this will also let us
// easily do things like middleware which will speed up development, allow using
// libraries across a wider scope of applications and simplify the learning
// process
const (
	EmptyAction   CLIAction = iota
	CommandAction           // Command and possibly flags
	FlagAction              // Only flags specified
	VersionAction
	HelpAction
	HelpCommandAction // TODO: an this reasonably merge in with help the way we merged in subcommand?
)

// TODO: ALWAYS consider we want our framework to play nicely with other
// frameworks, for example, we use this in our webframework. WE need to print
// out using a locale but ensure we are handing off responsibility for things
// like configuration off to the application so other frameworks can take over.

// TODO: Decide if flags should be segregated into global flags and
// command flags
// TODO: Command tree doesnt go here, it goes int he CLI program, and is
// generated from the initialization
// NOTE: Configuration should be handled by `cli` framework, it is very basic
// functionality of almost all applications. Should likely be not a string but
// dedicated a Config type similar to Flag. Config should have defaults, then
// override those defaults using ENV, then check default config path (based on
// command name) for a config file, or a flag for config file), then finally
// override values with flag?
type Context struct {
	CLI           *CLI // NOTE: We considered removing this and only giving minimal information defined in context, but its not necessary
	Locales       map[string]string
	Config        map[string]string
	CommandPath   []string              // This provides us with our current routing path, but we only need names not full input commands for each command in path
	Command       *InputCommand         // Endpoint of our input
	Flags         map[string]*InputFlag // SHould be both the application flags AND the CURRENT command flags if they exist
	ParameterType DataType              // Parameter is NOT owned by a command since it can exist WITHOUT a command
	Parameters    []string
	Args          []string // These are more like raw arguments perhaps
}

func (self *Context) CommandName() (name string) {
	if len(self.CommandPath) == 0 {
		name = self.CommandPath[len(self.CommandPath)-1]
	}
	return name
}

///////////////////////////////////////////////////////////////////////////////
// TODO: Since we are migrating to the commandchain concept based on the pathin
// in the command tree specified with the input, hasCommand can not just be
// based on the endpoint command, it
func (self *Context) hasCommand(name string) bool {
	for _, command := range self.CommandPath {
		if command == name {
			return true
		}
	}
	return false
}

func (self *Context) hasVersionArgument() bool { return self.hasFlag("version") }

func (self *Context) hasHelpArgument() bool {
	// TODO: We don't need scope, we find that out by checking IF a command
	// exists. If not then the scope is
	// elsif self.Command.IsNotSubcommand (can only be here IF the command is
	// help, so can this condition even be met? elsif self.Command.
	//return self.hasFlag("help")
	return self.CommandName() == "help" || self.hasFlag("help")
}

func (self *Context) hasNoCommands() bool   { return len(self.CommandPath) == 0 }
func (self *Context) hasNoFlags() bool      { return len(self.Flags) == 0 }
func (self *Context) hasNoArguments() bool  { return len(self.CommandPath) == 0 && len(self.Flags) == 0 }
func (self *Context) hasNoParameters() bool { return len(self.Parameters) == 0 }

func (self *Context) hasFlag(name string) bool {
	// TODO: This should check against the app + the current command flags (or
	// even all commands in path?)
	_, ok := self.Flags[name]
	return ok
}

// TODO: Would be nice to sometime find a nice shorthand for these type of loop
// booleans that could make code way more readable and cut down on something we
// do in literally every application
func (self *Context) hasFlags(names ...string) bool {
	for _, name := range names {
		if !self.hasFlag(name) {
			return false
		}
	}
	return true
}

func (self *CLI) parse(arguments []string) *Context {
	context := &Context{
		CLI:         self,
		CommandPath: []string{},
		Command:     &InputCommand{},
		Flags:       map[string]*InputFlag{},
		Parameters:  []string{},
		Args:        []string{},
	}

	for _, argument := range arguments {
		if argument[0] == "-"[0] && len(argument) > 1 {
			// Flag
			var flag *InputFlag
			if argument[:1] == "--"[:1] {
				// Long Flag - convention is enforcing '=' on Long val
				if strings.Contains(argument, "=") {
					// Not Bool Type
					// TODO: Could do this in less moves by doing this and using length
					flagNameAndValue := strings.Split(argument, "=")
					flag.Name = flagNameAndValue[0]
					flag.Value = flagNameAndValue[1]
				} else {
					// Bool Type
					flag.Name = argument
					flag.Value = Bool
				}
			} else {
				// Short Flag -
				// TODO: we are migrating to supporting only single character short flags
				// TODO: Support multiple small arguments in one declaration, for example:
				// (stacking)
				//         ls -lah

			}
		} else {
			// Command OR Parameters
			// TODO: Depends on if the command exists for the last command in the path
			// TODO: We could scan parameters for spaces and indicator a flag exists
			// to enable trailing flags for ease-of-use
		}
	}
	return context
}
