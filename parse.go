package cli

import (
  "fmt"
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

func (self *CLI) Parse(args []string) *Context {
  defer self.benchmark(time.Now(), "benmarking argument parsing")
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
        command.Parent = self.Context.Command

        self.Context.Commands = self.Context.Commands.Add(command)
        self.Context.Flags = append(self.Context.Flags, command.Flags...)

        self.Context.Arguments = self.Context.Arguments.Add(self.Context.Commands.First())

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
          self.Context.Arguments = self.Context.Arguments.Add(self.Context.Params.First())
        }
      }
    }
  }


  self.Context.DevOutput()

  // TODO: Need a way to scan for commands or flags for the HELP and VERSION
  // ones which will override essentially everything but OnStart and OnExit

  if self.Actions.OnStart != nil {
    fmt.Println("as expected, cli.Actions.OnStart is NOT nil")
    // TODO :Add it to the action chain
    self.Context.Actions = self.Context.Actions.Add(self.Actions.OnStart)
  }
  fmt.Println("Number of actions after adding onStart if not nil", len(self.Context.Actions))
  // TODO: Are the hidden command and flag currently being added? This might
  // be needed 
  if self.Context.Commands.HasCommand("version") || 
     self.Context.Flags.Assigned().HasFlag("version") {
    // TODO: Instead of just simply printing the render, we should add the
    // function to actions to be executed when Execute() is called
    self.RenderVersionTemplate()
  } else if self.Context.Commands.HasCommand("help") || 
            self.Context.Flags.Assigned().HasFlag("help") {
    // TODO: To simplify these help, could just always do comamnd before last
    // **in fact it should be this! and this may mean we get rid of .Command
    // altogether, dpending on its usefulness in the application layer**, but
    // this will allow simplicity in this function and the most intuitive
    // output from any given input
    // Command should be second to last that is printed help (if command help)
    // and last if flag. 
    self.RenderHelpTemplate(self.Context.Commands.Last()) 
    //self.RenderHelpTemplate(self.Context.Commands.Last())
  } else {
    // TODO: Command check here, look for last command (flipped has occured by
    // now) then check if it has an action and iterate over each one back to the
    // pseudo command. BUT WE NEED TO ASSURE THE THE FALLBACK HAS BEEN ASSIGNED
    // TO THE PSUEDO COMMANDS ACTION (probably should happen in the New()
    // fuction in cli.go

    // NOTE: We iterate over the commands backwards and have the fallback in the 
    // last position in the psuedo-command for the app. And for now we only run
    // one command by breaking after finding the first available action. Instead 
    // of every command in the chain (that might be a desirable option in the future).
    for _, command := range self.Context.Commands {
      if command.Action != nil {
        self.Context.Actions = append(self.Context.Actions, command.Action)
        break
      }
    }
     

  }

  if self.Actions.Fallback != nil {
  }
  // Check if the action is either version or help
  // Then look at last command for action (we could assign fallback to psuedo
  // command for simplicity but to be fair its not all that much simpler)
  // TODO: Add fallback (but need a way to determine if a command action like
  // help or version is being run (or even action flags like help or version)
  // and only add the fallback in the condition those are not run
  if self.Actions.OnExit != nil {
    // TODO: Add it to the action chain
    self.Context.Actions = self.Context.Actions.Add(self.Actions.OnExit)
  }
  fmt.Println("Number of actions after adding onExit if not nil", len(self.Context.Actions))

  // NOTE: Before handing the developer using the library the context we put
  // them in the expected left to right order, despite it being easier for us
  // to access in this function in the reverse order.
  self.Context.Arguments = Reverse(self.Context.Arguments)
  self.Context.Commands = ToCommands(Reverse(self.Context.Commands.Arguments()))
  self.Context.Flags = ToFlags(Reverse(self.Context.Flags.Arguments()))
  self.Context.Params = ToParams(Reverse(self.Context.Params.Arguments()))

  // TODO: Need to add the action parsing, determine if both fallback and global
  // are needed, detect if hooks exist and iterate through
  //
  //  * add them in order to be executed for a very simple excute command.
  //

  // TODO: If we avoid assing any actions until here, like in CLI context. We
  // can just simply iterate over every command including the glboal psuedo
  // command. SO THE PSEUDO COMMAND NEEDS TO BE LOADED in New() under cli.go

  return self.Context
}
