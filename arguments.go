package cli

import (
  "fmt"
  "os"
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
  IsValid() bool
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
  Actions  actions
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
func (self *CLI) ParseArgs() *Context {
  defer self.benchmark(time.Now(), "benmarking argument parsing")
  // TODO: Keep this field in context? If we are not using it below why are we
  // saving it exactly?
  self.Context.Args = os.Args[1:]


  fmt.Printf("%v \n", self.commands())

  fmt.Printf("commands:\n")
  for _, command := range self.commands() {
    fmt.Printf("command name: %v \n", command.Name)
  }

  for index, argument := range os.Args[1:] {
    // TODO: Must not ToLower Params type arguments
    argument = strings.ToLower(argument)

    fmt.Printf("parsing argument: '%v' \n", argument)

    // Flag parse
    if flagType, ok := HasFlagPrefix(argument); ok {
      fmt.Printf("%v\n", flagType.Name())



      //parsedFlags = append(parsedFlags, chain.ParseFlag(flagType, argument, chain.NextArgument(index)))

      //context.ParseFlag(index, flagType, &Flag{Name: argument})


      // TODO: Add the parsed FLAG to the chain.Arguments too before going
      // through the loop again! 

    } else {
      fmt.Printf("else (argument is not a flag, must be command or param)\n")
      // Command parse
      if command, ok := self.commands().Name(argument); ok {
        fmt.Printf("found subcommand %v\n", command.Name)
        //command.Parent = self.commands().Last()
        //self.commands().Add(command)

        // TODO: Add the parsed COMMAND to the chain.Arguments too before going
        // through the loop again! 

      } else {
        fmt.Printf("parsing param")
        // Param parse
        for _, paramArguments := range self.Context.Args[index:] {
          // TODO: Would like to be able to add flags to the end after params
          //         add index, then after first one begin checking if they are
          //         flags if not flag then its param

          for _, paramArgument := range paramArguments {
            self.Context.Chain.Params = append(
              self.params(), 
              &Param{
                Value: string(paramArgument),
              },
            )
          }
          // TODO: Add the parsed PARAM to the chain.Arguments too before going
          // through the loop again!
          self.arguments().Add(
            self.params().Last(),
          )
        }
        break
      }
    }
    // Argument Parse
    // TODO: Parse the argument to establish the chain.arguments
    // Flag Parse

    // TODO: Populate Actions in Chain + Build Execute command (executes actions
    // in order put in the chain.Actions slice)
  }

  //chain.UpdateFlags(parsedFlags)
  // TODO: This needs to be relpaced by a function that iterates over the
  // chain.Flags (and sets their values based on the parsed flags. 
  //   OR
  // we change the values as we iterate through the arguments

  //var debugValue bool
  //debugFlag, err := chain.Comamnds.First().Flag("debug")
  //if err == nil {
  //  debugValue := debugFlag.Bool()
  //}

  // TODO: Here is where we will cache the chain objects in the context before
  // passing context object out of the parse function
  return self.Context
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
func (self *Chain) ParseFlag(flagType FlagType, argument string, nextArgument string) (parsedFlag *Flag) {

  fmt.Printf("flagType: %v \n", flagType)
  flagParts := strings.Split(flagType.TrimPrefix(argument), "=")
  parsedFlag.Name = strings.ToLower(flagParts[0])
  if len(flagParts) == 2 {
    parsedFlag.Param.Value = flagParts[1]
  } else if len(flagParts) == 1 {
    if _, ok := HasFlagPrefix(nextArgument); ok {
      parsedFlag.Param.Value = "1"
    } else {
      parsedFlag.Param.Value = nextArgument
    }
  }

  flagFound := false
  for _, command := range self.Commands.Reversed() {
    if len(nextArgument) != 0 && command.is(nextArgument) {
      parsedFlag.Param.Value = "1"
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
              parsedFlag.Param.Value = flagParts[1]
            } else {
              // TODO: Needs to check if nextArgument is viable, if not, then
              //       "1"
            }
          } else if flag.Alias == string(stackedFlag) {
            parsedFlag.Param.Value = "1"
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
          commandFlag.Param.Value = parsedFlag.Param.Value
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
// TODO: Bug exists here, if the flag contains a dash in the middle like
// app-path then it will turn into -

func FlagNameForType(flagType FlagType, argument string) (name string) {
  // TODO: use strings.HasPrefix()
  argument = flagType.TrimPrefix(argument)

  return strings.Split(name, "=")[0]
}

func (self *Context) NextArgument(index int) string {
  if index+1 < len(self.Args) {
    return self.Args[index+1]
  }
  return ""
}
