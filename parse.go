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

    // TODO
    // Right now stacked flags work fine, but a single short flag space
    // separated will not work; but equals works 

    // Flag parse
    if flagType, ok := HasFlagPrefix(argument); ok {
      argument = flagType.TrimPrefix(argument)
      switch flagType {
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
        self.Context.Params = self.Context.Params.Add(argument)

        if flag := self.Context.Arguments.PreviousFlag(); flag != nil {
          fmt.Println("are we getting here because we like ened to be here")

          fmt.Println("flag.Param.Value: ", flag.Param.Value) 
          fmt.Println("len(flag.Param.Value) == 0", len(flag.Param.Value) == 0)

          fmt.Println("...")
          fmt.Println("flag.Default: ", flag.Default)

          if len(flag.Param.Value) == 0 ||
             flag.Param.Value == flag.Default ||
             len(flag.Default) == 0 {

            flag.Param = self.Context.Params.Last()
          }
        }

        // TODO: When we flip this it will need to be First()
        self.Context.Arguments = self.Context.Arguments.Add(
          self.Context.Params.Last(),
        )
      }
    }
  }

  self.Context.DevOutput()

  return self.Context
}
