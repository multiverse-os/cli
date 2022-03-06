package cli

import (
  "time"
  "os"
  "fmt"
)

func (self *CLI) Parse(args []string) *Context {
  defer self.benchmark(time.Now(), "benmarking argument parsing")
  for index, argument := range os.Args[1:] {
    fmt.Println("index: ", index)
    fmt.Printf("determining if flag, command or param: '%v' \n", argument)
    // Flag parse
    if flagType, ok := HasFlagPrefix(argument); ok {
      fmt.Println("argument has flag prefix, determining if long or short")
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
        self.Context.Params, _ = self.Context.Params.Add(argument)

        self.Context.Arguments = self.Context.Arguments.Add(
          self.Context.Params.Last(),
        )
        // TODO: THIS WORKS, but want to try to use Add()
        //self.Context.Arguments = append(
        //  self.Context.Arguments, 
        //  *self.Context.Params.Last(),
        //)
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

  firstCommand := self.Context.Commands.First()
  fmt.Println("Command(first)")
  fmt.Println("  Name:        ", firstCommand.Name)
  fmt.Println("  Alias:       ", firstCommand.Alias)
  fmt.Println("  Description: ", firstCommand.Description)
  fmt.Println("  Hidden:      ", firstCommand.Hidden)
  fmt.Println("  Parent:      ", firstCommand.Parent)
  fmt.Println("  Subcommands: ", firstCommand.Subcommands)
  fmt.Println("  Flags:       ", firstCommand.Flags)
  fmt.Println("  Action:      ", firstCommand.Action)
  fmt.Println("  Hooks:       ", firstCommand.Hooks)

  return self.Context
}
