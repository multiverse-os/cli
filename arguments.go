package cli

import (
  "fmt"
	"path/filepath"
	"strings"
	"time"

  data "github.com/multiverse-os/cli/data"
)

type arguments int

type Arguments interface {
	Get(index int) string
	Flag(index int) string
	Command(index int) string
	Params() []string
	Len() int
	Present() bool
	Slice() []string
}

// TODO: I dont remember if I tried to use just a type Chain []*Commands but I
// feel like I spent a lot of time working that one out -- but with programminmg
// could have changed since then because we are so high up in terms of
// abstraction.
type Chain struct {
	Commands commands
}

func (self *Chain) Length() int { return len(self.Commands) }

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

func (self *Chain) AddCommand(command *Command) {
	self.Commands = append(self.Commands, command)

  // TODO: IMPORTANT
  // This takes all flags from the command then just sets them to default
  // regardless if they are assigned in the cmd line. This is a major bug. Need
  // to detect each and set default
  // acordingly
	//flags := []Flag{}
	//for _, flag := range command.Flags {
	//	//if len(flag.Value) == 0 {
	//	flag.Value = flag.Default
	//	//}
	//	flags = append(flags, flag)
	//}
	//command.Flags = flags

  for _, commandFlag := range command.Flags {
    if data.IsNil(commandFlag.Value) {
      commandFlag.Value = commandFlag.Default
    }
  }
}

func (self *Chain) Last() *Command             { return self.Commands[len(self.Commands)-1] }
func (self *Chain) NoCommands() bool           { return self.IsRoot() && len(self.First().Subcommands) == 0 }
func (self *Chain) HasCommands() bool          { return self.IsRoot() && 0 < len(self.First().Subcommands) }
func (self *Chain) IsRoot() bool               { return len(self.Commands) == 1 }
func (self *Chain) IsNotRoot() bool            { return 1 < len(self.Commands) }
func (self *Chain) PathExample() (path string) { return strings.Join(self.Path(), " ") }

func (self *Chain) HasSubcommands() bool {
	return self.IsNotRoot() && (0 < len(self.Last().Subcommands))
}

func (self *Chain) Flags() (flags map[string]*Flag) {
	for _, command := range self.Commands {
		for _, flag := range command.Flags {
			flags[flag.Name] = flag
		}
	}
	return flags
}

func (self *Chain) Path() (path []string) {
	for _, command := range self.Commands {
		path = append(path, command.Name)
	}
	return path
}

func (self *Chain) Reversed() (commands []*Command) {
	for i := len(self.Commands) - 1; i >= 0; i-- {
		commands = append(commands, self.Commands[i])
	}
	return commands
}

func (self *Chain) ReversedPath() (path []string) {
	for i := len(self.Commands) - 1; i >= 0; i-- {
		path = append(path, self.Commands[i].Name)
	}
	return path
}

func (self *CLI) Parse(arguments []string) (*Context, error) {
	defer self.benchmark(time.Now(), "benmarking argument parsing and action execution")
	cwd, executable := filepath.Split(arguments[0])

	context := &Context{
		CLI:          self,
		CWD:          cwd,
		Command:      &self.Command,
		Executable:   executable,
		CommandChain: &Chain{},
		Params:       Params{},
		Flags:        make(map[string]*Flag),
		Args:         arguments[1:],
	}

	context.CommandChain.AddCommand(&self.Command)

	var parsedFlags flags
	for index, argument := range context.Args {
		if flagType, ok := HasFlagPrefix(argument); ok {
			// TODO: Need to handle skipping next argument when next argument is used
			parsedFlags = append(parsedFlags, context.ParseFlag(flagType, argument, context.NextArgument(index)))

			//context.ParseFlag(index, flagType, &Flag{Name: argument})
		} else {
			if command, ok := context.Command.Subcommand(argument); ok {
				command.Parent = context.Command
				context.Command = command
				context.CommandChain.AddCommand(context.Command)
			} else {
				for _, param := range context.Args[index:] {
					context.Params.Value = append(context.Params.Value, param)
				}
				break
			}
		}
	}

	context.UpdateFlags(parsedFlags)

	self.Debug = context.HasFlag("debug")

  // TODO: This is currently the router, it would be nice to be able to produce
  // a standard URL like output (even have a URI scheme, like 

  //  cli://user@program:/command/subcommand?params
  //  
  //  OR somethjing similar, then be able to route to a defined functions in a
  //  controller section, but additionally and importantly, provide consistent,
  //  specific and useful details to the controller function so that they can be
  //  slim and written similarly. 
  // 

  // TODO: We may want to add the ability to do before hook and after hooks as
  // alternative or in addition to the default action. This woudl also be nice
  // for like sections, or namespaces as it is sometimes referred to in web
  // applications. The facility for this should be considered when building out
  // below. Because it is fairly critical to the fluid design. Global, commands,
  // and command level should all likely have the functionality.

  fmt.Printf("context.Command.Action: %v\n", context.Command.Action)

  if context.Command.is("version") || context.HasFlag("version") {
		self.RenderVersionTemplate()
  } else if context.HasFlag("help") { // TODO: Removed condition where subcommands but no action that should get help output BUT -- should default action run regardless or above happens only when no default
		  context.RenderHelpTemplate(context.Command)
  } else if context.Command.is("help") {
		  context.RenderHelpTemplate(context.Command.Parent)
  } else {
      context.Execute()
	}

	return context, nil
}

// TODO: MISSING ABILITY TO PARSE FLAGS THAT ARE USING "QUOTES TO SPACE TEXT".
// TODO: MISSING Flags of slice types can be passed multiple times (-f one -f two -f three)
// TODO: MISSING Collect ALL arguments trailing `--`
// TODO: MISSING ability to stack flag names of any size (right now assumes only
//       1 character size is allowed for short command names).
// NOTE: Check if nextArgument is flag, flag is a boolean if nextArgument is
//       either a flag or is a known command.
// TODO: ==IDEA== Maybe have a expand function that goes over arguments, groups
// up quoted sections, expand out stacked flags, convert " " separators on flags
// with "=" separator.
func (self *Context) ParseFlag(flagType FlagType, argument, nextArgument string) (parsedFlag *Flag) {
	flagParts := strings.Split(StripFlagPrefix(argument), "=")
	parsedFlag.Name = strings.ToLower(flagParts[0])
	if len(flagParts) == 2 {
		parsedFlag.Value = flagParts[1]
	} else if len(flagParts) == 1 {
		if _, ok := HasFlagPrefix(nextArgument); ok {
			parsedFlag.Value = "1"
		} else {
			parsedFlag.Value = nextArgument
		}
	}

	flagFound := false
	for _, command := range self.CommandChain.Reversed() {
		if len(nextArgument) != 0 && command.is(nextArgument) {
			parsedFlag.Value = "1"
		}
		for _, flag := range command.Flags {
			if flag.is(parsedFlag.Name) {
				parsedFlag.Name = flag.Name
				flagFound = true
			}
		}
	}

	if !flagFound {
		// TODO: This means the flag was not located; so HERE we check for the FLAG
		// STACKING. However, the best way to do variable short name length is
		// likely checking 1 2 3, throwing out 1, then again 1 2 3 etc.
		for index, stackedFlag := range parsedFlag.Name {
			for _, subcommand := range self.CommandChain.Reversed() {
				for _, flag := range subcommand.Flags {
					if index == len(parsedFlag.Name)+1 {
						if len(flagParts) == 2 {
							parsedFlag.Value = flagParts[1]
						} else {
							// TODO: Needs to check if nextArgument is viable, if not, then
							//       "1"
						}
					} else if flag.Alias == string(stackedFlag) {
						parsedFlag.Value = "1"
					}
				}

			}
		}
	}

	return parsedFlag
}

func (self *Context) UpdateFlags(parsedFlags flags) {
	for _, parsedFlag := range parsedFlags {
		for _, command := range self.CommandChain.Reversed() {
			for _, commandFlag := range command.Flags {
				if commandFlag.is(parsedFlag.Name) {
					commandFlag.Value = parsedFlag.Value
				}
			}
      // TODO: Was this style required to get the saves of data>?
      //       going to save some evidednce it existed to save headache later if
      //       that turns out why it wont work
			//command.Flags = flags
		}
	}

}

// NOTE: These are here for dev reasons while parsing is being completed; once
// it is these can be moved into the appropriate files like flag.go
func StripFlagPrefix(flagName string) string { return strings.Replace(flagName, "-", "", -1) }

func FlagNameForType(flagType FlagType, argument string) (name string) {
	switch flagType {
	case Short:
		name = argument[1:len(argument)]
	case Long:
		name = argument[2:len(argument)]
	}
	return strings.ToLower(strings.Split(name, "=")[0])
}

func (self *Context) NextArgument(index int) string {
	if index+1 < len(self.Args) {
		return self.Args[index+1]
	}
	return ""
}
