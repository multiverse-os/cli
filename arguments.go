package cli

import (
  "fmt"
	"strings"
	"time"

  //data "github.com/multiverse-os/cli/data"
)

// TODO: Why is this int instead of []argument and argument being a raw string
// representation of any given argument in the chain?
//type arguments int

// NOTE: Raw string of the argument after splitting, could be flag, stacked
// flags, command, or param
//type argument string

type ArgumentType int

const (
  CommandArgument ArgumentType = iota
  FlagArgument
  ParamArgument
)

type Argument interface {
  Type()    ArgumentType
}

type arguments []*Argument 

func (self arguments) Last() *Argument { return self[self.Count()-1] }
func (self arguments) Count() int { return len(self) }

// TODO: Later look into how this is being used, is it inline and necessary not
// to produce an error? Or can we produce an error if the index is incorrect?
// But consider: since there is only 1 error condition, if the returned value is
// empty, you know its the only error possible, because an argument can't be
// blank, otherwise its not an argument. 
func (self arguments) Get(index int) *Argument {
  if self.Count() < index {
    return self[index]
  }
  return nil
}

// TODO: I dont remember if I tried to use just a type Chain []*Commands but I
// feel like I spent a lot of time working that one out -- but with programminmg
// could have changed since then because we are so high up in terms of
// abstraction.
type Chain struct {
  // TODO: Should generic argument object be created for storing the full
  // command line as it was entered but as an generic interface for commands,
  // flags and params?
  Arguments arguments 

	Commands commands
  Flags    flags
  Params   params

  // TODO: Our end goal for the CLI framework this session will be developing
  // the code to properly intialize and load all the actions in the order they
  // should be run in into this variable and then executing those actions. 
  Action   actions
}

// TODO: Not sure if this survives 
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

// TODO: We had the idea to switch this to be a method of Commands since this
// functionality is kinda out of scope for the chain object

// TODO: Move Reversed() in arguments to commands object as a method
// Commands.
//func (self *Chain) NoCommands() bool           { return self.IsRoot() && len(self.First().Subcommands) == 0 }
//func (self *Chain) HasCommands() bool          { return self.IsRoot() && 0 < len(self.First().Subcommands) }

// TODO: Was this being used? I think possibly in help but feels wrong.
//func (self *Chain) PathExample() (path string) { return strings.Join(self.Path(), " ") }

//func (self *Chain) HasSubcommands() bool {
//	return self.IsNotRoot() && (0 < len(self.Last().Subcommands))
//}

// TODO: Should move this to commands, but arguments still needs a way to build
// all the flags
//func (self *Chain) Flags() (flags map[string]*Flag) {
//	for _, command := range self.Commands {
//		for _, flag := range command.Flags {
//			flags[flag.Name] = flag
//		}
//	}
//	return flags
//}

func (self *CLI) Parse(arguments []string) (*Context, error) {
	defer self.benchmark(time.Now(), "benmarking argument parsing and action execution")

  // TODO: Build the chain then apply it to the context at the end and return
  // it?

  // TODO: For the hooks we should reverse iterate over the command chain, and
  // puill out each of the hooks and merge tbhem into the hooks held in context
  // then that would resolve how they would be obtained in the execute function,
  // and make the execute function pretty complete leaving most of the logic in
  // there and giving us a clean interaction with commands which are at the
  // heart of this cli framework. in the end its all just about parsing the args
  // to execute a defined action (and its hooks). 

	context.Chain.Commands.Add(&self.Command)

	var parsedFlags flags
	for index, argument := range context.Args {
		if flagType, ok := HasFlagPrefix(argument); ok {
			// TODO: Need to handle skipping next argument when next argument is used
      // TODO: What about flags with values? This is probably in need of
      // rewriting
			parsedFlags = append(parsedFlags, context.ParseFlag(flagType, argument, context.NextArgument(index)))

			//context.ParseFlag(index, flagType, &Flag{Name: argument})
		} else {
			if command, ok := context.Command.Subcommand(argument); ok {
				command.Parent = context.Command
				context.Command = command
        // TODO: Since this is a commands type we can add the AddCommand or just
        // Add so `commands.Add(command)` 
				context.Chain.Commands.Add(context.Command)
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
      context.ExecuteActions()
	}


	return &Context{
		CLI:          self,
    Process:      Process(),
		Command:      &self.Command,
		//Flags:        make(map[string]*Flag),
    Chain:        argumentChain,
		Args:         arguments[1:],
	}, nil

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
	for _, command := range self.Chain.Commands.Reversed() {
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
			for _, subcommand := range self.Chain.Commands.Reversed() {
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
