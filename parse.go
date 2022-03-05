package cli

import (
  "time"
  "os"
  "fmt"
)

//type chain struct {
//  Arguments arguments 
//  Commands commands
//  Flags    flags
//  Params   params
//  Actions  actions
//}

// TODO: We took off the error for command chaining off context
//       for something like cli.ParseArgs().Execute() but that
//       may prove unwise and we may add the error back
//       Should it be returning CLI with context or Context with CLI?
func (self *CLI) Parse(args []string) *Context {
  defer self.benchmark(time.Now(), "benmarking argument parsing")
  for index, argument := range os.Args[1:] {
    fmt.Println("index: ", index)
    fmt.Printf("determining if flag, command or param: '%v' \n", argument)
    // Flag parse
    if flagType, ok := HasFlagPrefix(argument); ok {
      fmt.Println("argument has flag prefix, determining if long or short")
      fmt.Printf("%v\n", flagType.Name())
      argument = flagType.TrimPrefix(argument)

      // TODO: This loops over the flags twice, it will likely be better to
      // combine the logic of Name() and Reversed()
      // TODO: Implement a is() function specifically for checking against 
      // alias or name to reduce the unncessary string comparisons
      switch flagType {
      case Short:
        if len(argument) == 1 {
          // Not stacked; So, we know value is boolean and == "1"

          flag := self.Context.Flags.Reversed().Name(argument)
          if flag != nil {
            // TODO: Update the flag param value (from default) and confirm it is
            // working by checking after the parse function is ran
            flag.SetTrue()
          }
        }else{
          // Stacked
          for index, shortFlag := range argument {
            fmt.Println("stacked short flag parsed:", string(shortFlag))
            // Last Flag (should both functions be at the top?)
            // When argument length == index +1 
            //  || flag BEFORE equals sign 
            if len(argument) == index + 1 {
              // Last short flag could be boolean == "1" 
              //    || check next argument for flag or command, else assume param

              // TODO: Check if the next argument is COMMAND or FLAG or NIL
              //             In this condition flag value == "1"

              //      ELSE
              //       next argument is the new flag value
            }
            // TODO: When we hit the = sign, we take the last flag (which was
            // set to true) and we replace it with argument[index+1:]
            // NOTE: '=' = 61;
            if shortFlag == 61 {
              fmt.Println("equals sign found, assigning param to last flag added")
              // TODO: But this wont work yet because we are not yet actually
              // locating the defined flags and adding them to the flag chain
              flag := self.Context.Flags.Last()
              flag.Set(argument[index+1:])
              break
            }

            // Every flag before the last one value is boolean and == "1"
            flag := self.Context.Flags.Reversed().Name(string(shortFlag))
            if flag != nil {
              fmt.Println("flag ", string(shortFlag), " exists, setting it to true")
            // TODO: Update the flag param value (from default) and confirm it is
            // working by checking after the parse function is ran
              flag.SetTrue()
            }

          }
        }
      case Long:
        // Long flag could be boolean == "1" 
        //    || if contains = it has param
        //    || check next argument for flag or command, else assume param
        
      }
      // TODO: All these conditions need to be supported
      // -fl --flag
      // -flag=param
      // -flag param

      //parsedFlags = append(parsedFlags, chain.ParseFlag(flagType, argument, chain.NextArgument(index)))

      //context.ParseFlag(index, flagType, &Flag{Name: argument})

      // TODO: Recursively call the Parse function using arguments skipping up to 
      //       the index+1
      // (for the next argument when used by a flag) 
    } else {
      if command, ok := self.Context.Command.Subcommand(argument); ok {
        // Command parse
        command.Parent = self.Context.Command

        self.Context.Commands = self.Context.Commands.Add(command)
        self.Context.Flags = append(self.Context.Flags, command.Flags...)

        self.Context.Arguments = self.Context.Arguments.Add(
          self.Context.Commands.Last(),
        )

        self.Context.Command = self.Context.Commands.Last()
      } else {
        // Param parse
        fmt.Println("parsing param: ", argument)
        self.Context.Params = self.Context.Params.Add(argument)
        self.Context.Arguments = self.Context.Arguments.Add(
          self.Context.Params.Last(),
        )
      }
    }
  }

  fmt.Println("================")
  fmt.Println("parsing COMPLETED!") 
  fmt.Println("arguments parsed: ", len(self.Context.Arguments))
  fmt.Println("                  ", self.Context.Arguments)
  fmt.Println("commands parsed:  ", len(self.Context.Commands))
  fmt.Println("                  ", self.Context.Commands)
  fmt.Println("flags parsed:     ", len(self.Context.Flags))
  fmt.Println("                  ", self.Context.Flags)
  fmt.Println("params parsed:    ", len(self.Context.Params))
  fmt.Println("                  ", self.Context.Params)
  fmt.Println("---------------")
  for index, _ := range self.Context.Arguments {
    fmt.Println((*self.Context.Arguments[index]).Type())
    fmt.Println(self.Context.Arguments[index])
  }
  fmt.Println("---------------")
  fmt.Println("command:          ", &self.Context.Command)


  fmt.Println("\n\n")
  fmt.Println("================")


  fmt.Println("must test if changing the command affects the command stored in")
  fmt.Println("arguments and commands. and vice versa\n")
  fmt.Println("



  return self.Context
}

// TODO: MISSING ABILITY TO PARSE FLAGS THAT ARE USING "QUOTES TO SPACE TEXT".
// TODO: MISSING Flags of slice types can be passed multiple times (-f one -f two -f three)
// TODO: MISSING ability to stack flag names of any size (right now assumes only
//       1 character size is allowed for short command names).
// NOTE: Check if nextArgument is flag, flag is a boolean if nextArgument is
//       either a flag or is a known command.
//func (self *chain) ParseFlag(flagType FlagType, argument string, nextArgument string) (parsedFlag *Flag) {
//
//  fmt.Printf("flagType: %v \n", flagType)
//
//  flagParts := strings.Split(flagType.TrimPrefix(argument), "=")
//
//  // TODO: This assumes the flag should be created even if the flag is not
//  // defined by in the global flags or the current commands flags at
//  // initialization
//  parsedFlag.Name = flagParts[0]
//
//  if len(flagParts) == 2 {
//    parsedFlag.Param.Value = flagParts[1]
//  } else if len(flagParts) == 1 {
//    // TODO: This is wrong; currently it assumes if the next argument is a flag
//    // then the current flag is a boolean type flag. But if the next argument is
//    // a command, then it also a boolean type flag (also if it is a param but
//    // the param distinction, as in if it is a flag param vs cli app param is
//    // going to be difficult to distinguish and a fallback will need to be
//    // implemented)
//    if _, ok := HasFlagPrefix(nextArgument); ok {
//      parsedFlag.Param.Value = "1"
//    } else {
//      parsedFlag.Param.Value = nextArgument
//    }
//  }
//
//  flagFound := false
//  for _, command := range self.chain.Commands.First.Reversed() {
//    // TODO: Or merge with above and do the check on next argument to see if its
//    // a flag
//    if data.IsBlank(nextArgument) && command.is(nextArgument) {
//      parsedFlag.Param.Value = "1"
//    }
//    for _, flag := range command.Flags {
//      if flag.is(parsedFlag.Name) {
//        parsedFlag.Name = flag.Name
//        flagFound = true
//      }
//    }
//  }
//
//  if !flagFound {
//    // TODO: This means the flag was not located; so HERE we check for the FLAG
//    // STACKING. However, the best way to do variable short name length is
//    // likely checking 1 2 3, throwing out 1, then again 1 2 3 etc.
//  
//    // stacked flags: tar -xvf param 
//
//    for index, stackedFlag := range parsedFlag.Name {
//      for _, subcommand := range self.chain.Commands.First().Subcommands.Reversed() {
//        for _, flag := range subcommand.Flags {
//          if index == len(parsedFlag.Name)+1 {
//            if len(flagParts) == 2 {
//              parsedFlag.Param.Value = flagParts[1]
//            } else {
//              // TODO: Needs to check if nextArgument is viable, if not, then
//              //       "1"
//            }
//          } else if flag.Alias == string(stackedFlag) {
//            parsedFlag.Param.Value = "1"
//          }
//        }
//
//      }
//    }
//  }
//
//  return parsedFlag
//}
