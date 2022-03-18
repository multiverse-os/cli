package cli

import (
  "os"
  "strings"
  "time"

  data "github.com/multiverse-os/cli/data"
)

// TODO: How about we jut expose os.Args() through our own function, but also
// have the general Arguments function return fully parsed arguments with flags
// set to their valid values. And Parse() remains just Parse() 

// TODO: Remaining major task here is to move the action logic into a fucntion
// organized in the same order it should be executed. 

// TODO: Have not tested subcommand flag assignment and assignment to higher
// level commands if flag does not exist at the subcommand scope.
// TODO: Flags now only contains assigned flags, and the commands store the
// complete list of available flags (and the psuedo app command stores the
// global ones) -- changes to parse will need to reflect this.

func (self *CLI) Parse(args []string) (*Context, []error) {
  defer self.benchmark(time.Now(), "benmarking argument parsing")

  // TODO: Need to add the check for required flags (new feature just thought of
  // because seems like it would be useful no? But for this to ever occur we
  // need to not just return error but ERRORS- this is actually nice because we
  // would expect parse to return PARSING ERRORS no?-- which will let us have
  // our validations OH VALIDATIONS HOW I LOVE THE

  for _, argument := range os.Args[1:] {
    // Flag parse
    if flagType, ok := HasFlagPrefix(argument); ok {
      argument = flagType.TrimPrefix(argument)
      switch flagType {
      case Short:
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
      if command := self.Context.Command.Subcommand(argument); command != nil {
        // Command parse
        command.Parent = self.Context.Commands.First()

        self.Context.Commands = self.Context.Commands.Add(command)
        self.Context.Flags = append(self.Context.Flags, command.Flags...)

        self.Context.Arguments = self.Context.Arguments.Add(
          self.Context.Commands.First(),
        )

        self.Context.Command = self.Context.Commands.First()
      } else {
        // Params parse
        flag := self.Context.Arguments.PreviousIfFlag()
        if flag != nil {
          if flag.Param.value == flag.Default {
            flag.Param = NewParam(argument)
          }else{
            flag = nil
          }
        }
        if flag == nil {
          self.Context.Params = self.Context.Params.Add(NewParam(argument))
          self.Context.Arguments = self.Context.Arguments.Add(
            self.Context.Params.First(),
          )
        }
      }
    }
  }

  // NOTE: Before handing the developer using the library the context we put
  // them in the expected left to right order, despite it being easier for us
  // to access in this function in the reverse order.
  self.Context.Arguments = Reverse(self.Context.Arguments)
  self.Context.Commands = ToCommands(Reverse(self.Context.Commands.Arguments()))
  self.Context.Flags = ToFlags(Reverse(self.Context.Flags.Arguments()))
  self.Context.Params = ToParams(Reverse(self.Context.Params.Arguments()))

  if self.Actions.OnStart != nil {
    self.Context.Actions = self.Context.Actions.Add(self.Actions.OnStart)
  }

  for _, command := range self.Context.Commands {
    if command.Action != nil {
      self.Context.Actions = append(self.Context.Actions, command.Action)
      for _, flag := range command.Flags {
        if data.IsTrue(flag.Param.value) && flag.Action != nil {
          self.Context.Actions = append(self.Context.Actions, flag.Action)
        }
      }
      // NOTE: Break so only first available action is used. Fallback
      // should only run if no actions were defined by commands
      break
    }
  }

  if self.Actions.OnExit != nil {
    self.Context.Actions = self.Context.Actions.Add(self.Actions.OnExit)
  }

  return self.Context, []error{}
}
