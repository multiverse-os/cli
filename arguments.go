package cli

import (
	"strings"
	"time"

  //data "github.com/multiverse-os/cli/data"
)

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

func (self arguments) Add(argument Argument) (arguments arguments) {
  return append(arguments, &argument)
}

type Chain struct {
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
//func (self *Chain) Route(path []string) (*Command, bool) {
//	cmd := &Command{}
//	for index, command := range self.Commands {
//		if command.Name == path[index] {
//			if index == len(path) {
//				return command, true
//			} else {
//				cmd = command
//			}
//		} else {
//			return cmd, (len(cmd.Name) == 0)
//		}
//	}
//	return nil, (len(cmd.Name) == 0)
//}


// TODO: We took off the error for command chaining off context
//       for something like cli.Parse(os.Args).Execute() but that
//       may prove unwise and we may add the error back
//       Should it be returning CLI with context or Context with CLI?
func (self *CLI) Parse(arguments []string) *Context {
	defer self.benchmark(time.Now(), "benmarking argument parsing and action execution")

  chain := &Chain{
    Commands: Commands(self.Command),
    Flags: self.Command.Flags,
  }

	var parsedFlags flags
	for index, argument := range arguments[1:] {
		if flagType, ok := HasFlagPrefix(argument); ok {
      // TODO: yo, you like never did this- i mean look a few versions back, you
      // did -but you deleted it. 

      // TODO: we iterate over the flags in chain update it with the flag
      // currently being parsed (if it has a avlue ' ' or '=' after the flag).
      // if no value its a boolean.
      
			// TODO: Need to handle skipping next argument when next argument is used
      // TODO: What about flags with values? This is probably in need of
      // rewriting

      // TODO: TO properly parse a flag, we need to lcoate the defnied flag in
      // the CLI, and then update the value, then add that to the flag chain,
      // and add it to the command's flags (which should exist already but more
      // importantly we need to update it.
      // TODO: THIS IS CRITICAL AND CAN NOT BE SKIPPED ITS THE FLAG PARSING 
			//parsedFlags = append(parsedFlags, chain.ParseFlag(flagType, argument, chain.NextArgument(index)))

			//context.ParseFlag(index, flagType, &Flag{Name: argument})


      // TODO: Add the parsed FLAG to the chain.Arguments too before going
      // through the loop again! 

		} else {
      // Command parse
			if ok, command := self.Commands.Subcommand(argument); ok {
				command.Parent = chain.Commands.Last()
        chain.Commands.Add(command)

      // TODO: Add the parsed COMMAND to the chain.Arguments too before going
      // through the loop again! 

			} else {
        // Param parse
				for _, paramArguments := range arguments[index:] {
          for _, paramArgument := range paramArguments {
            chain.Params = append(chain.Params, &Param{Value: string(paramArgument)})
          }
          // TODO: Add the parsed PARAM to the chain.Arguments too before going
          // through the loop again!
          chain.Arguments.Add(chain.Params.Last())


				}
				break
			}

      // Argument Parse
      // TODO: Parse the argument to establish the chain.arguments
      // Flag Parse
		}

    // TODO: Populate Actions in Chain + Build Execute command (executes actions
    // in order put in the chain.Actions slice)
	}

	//chain.UpdateFlags(parsedFlags)
  // TODO: This needs to be relpaced by a function that iterates over the
  // chain.Flags (and sets their values based on the parsed flags. 
  //   OR
  // we change the values as we iterate through the arguments

	return &Context{
  // TODO: Test to see if this is actually working, and ensure it works with
  // both -debug and -d
  //              chain.Commands.Subcommand("name")
    Debug:        chain.Flags.Name("debug").Bool(),
		CLI:          self,
    Process:      Process(),
		Command:      chain.Commands.Last(),
    Flags:        chain.Flags,
    Params:       chain.Params,
    Chain:        chain,
	}
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
func (self *Chain) ParseFlag(flagType FlagType, argument, nextArgument string) (parsedFlag *Flag) {
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
	for _, command := range self.Commands.Reversed() {
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
			for _, subcommand := range self.Commands.Reversed() {
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

func (self *Chain) UpdateFlags(parsedFlags flags) {
	for _, parsedFlag := range parsedFlags {
		for _, command := range self.Commands.Reversed() {
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
