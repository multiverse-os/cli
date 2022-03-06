package cli

import (
  "fmt"
  "os"
  "strings"
  "time"

  data "github.com/multiverse-os/cli/data"
)

func (self *CLI) Parse(args []string) *Context {
  defer self.benchmark(time.Now(), "benmarking argument parsing")
  for _, argument := range os.Args[1:] {
    // Flag parse
    if flagType, ok := HasFlagPrefix(argument); ok {
      argument = flagType.TrimPrefix(argument)
      switch flagType {
        // TODO: and the preventing of non-boolean defaults not allowing
        // param.value to be set 
      case Short:
        // NOTE: One or more short flags (stacked and single)
        for index, flagAlias := range argument {
          // NOTE: Confirm we are not last && next argument is '=' (61) &&
          if len(argument) != index + 1 && argument[index+1] == 61 { 
            if flag := self.Context.Flags.Name(string(flagAlias)); flag != nil {
              if flagParam := argument[index+2:]; len(flagParam) != 0 {
                flag.Set(flagParam)
              }else{
                flag.SetDefault()
              }

              self.Context.Arguments = self.Context.Arguments.Add(flag)

              break
            }
          }else{
            if flag := self.Context.Flags.Name(string(flagAlias)); flag != nil {
              // NOTE: If the default value is not boolean or blank, no
              // assignment occurs to avoid input failures.
              if data.IsBoolean(flag.Default) {
                flag.ToggleBoolean()
              }else if len(flag.Default) == 0 {
                flag.SetTrue()
              }

              self.Context.Arguments = self.Context.Arguments.Add(flag)
            }
          }
        }
      case Long:
        longFlagParts := strings.Split(argument, "=")
        // When there is 1 part its a boolean type flag
        // TODO: DRY this up, its repeated from above; should be flag method
        if flag := self.Context.Flags.Name(string(longFlagParts[0])); flag != nil {
          if len(longFlagParts) == 1 {
            if data.IsBoolean(flag.Default) {
              flag.ToggleBoolean()
            }else if len(flag.Default) == 0 {
              flag.SetTrue()
            }
          }else if len(longFlagParts) == 2 {
            if 0 < len(longFlagParts[1]) {
              flag.Set(longFlagParts[1])
            }else{
              flag.SetDefault()
            }
          }

          self.Context.Arguments = self.Context.Arguments.Add(flag)
        }
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
        // TODO: THIS IS THE LAST PART OF THE PARSING FUCNTION, we jsut need to
        // assign any param we locate to the PREV flag

        // TODO: If PREV argument is type Flag (use type switch not enumerator)
        //       then also assign this param to the previous flags param 
        //       (but use same object, changing one should affect the other)
        //       and then that will cover both = and ' '
        // Param parse
        fmt.Println("parsing param: ", argument)
        self.Context.Params, _ = self.Context.Params.Add(argument)

        if flag := self.Context.Arguments.PreviousFlag(); flag != nil {
          fmt.Println("PREVIOUS FLAG to assign param: ", flag.Name)
          // TODO: Only assign param to flag if the param is empty, for example
          // if we have a flag 'cli --language=en otherparam' we dont want to
          // asign otherparam over en, we only assign it if 'cli --langauge
          // otherparam'

          // TODO: flag.Param.Value != "" && flag.Param.Value != flag.Default
          // then we assign it; but we also still assign it as a param
          if len(flag.Param.Value) == 0 || flag.Param.Value == flag.Default {
            fmt.Println("previous flag is eligible for assigning value")
            flag.Param = self.Context.Params.Last()
          }
        }

        // TODO: When we flip this it will need to be First()
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
