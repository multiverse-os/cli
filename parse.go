package cli

import (
  "time"
  "os"
  "fmt"

  data "github.com/multiverse-os/cli/data"
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
        fmt.Println("stacked flag, going into loop")
        // NOTE: One or more short flags (stacked and single)
        for index, shortFlag := range argument {
          // TODO: Still need to handle condition for flag has param and not
          // using equals but ' ' (space) 
          // if next major argument is not flag or command, then we should
          // just assign that param to both the last flag, and the general
          // params (using a single object)

          // NOTE: Confirm we are not last && next argument is '=' (61) &&
          if len(argument) != index + 1 && argument[index+1] == 61 { 
            if previousFlag := self.Context.Flags.Name(string(argument[index])); previousFlag != nil {
              if flagParam := argument[index+2:]; len(flagParam) != 0 {
                previousFlag.Set(flagParam)
              }else{
                previousFlag.SetDefault()
              }
            }
          }else{
            if flag := self.Context.Flags.Name(string(shortFlag)); flag != nil {
              if data.IsTrue(flag.Default) || data.IsFalse(flag.Default) {
                flag.ToggleBoolean()
              }else{
                flag.SetTrue()
              }
            }
          }
        }
      case Long:
        // Long flag could be boolean == "1" 
        //    || if contains = it has param
        //    || check next argument for flag or command, else assume param

      }
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
        // TODO: If PREV argument is type Flag (use type switch not enumerator)
        //       then also assign this param to the previous flags param 
        //       (but use same object, changing one should affect the other)
        //       and then that will cover both = and ' '
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
