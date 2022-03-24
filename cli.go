package cli

import (
  "fmt"
  "time"
  "strings"

  data "github.com/multiverse-os/cli/data"
)

///////////////////////////////////////////////////////////////////////////////
// Ontology of a command-line interface
///////////////////////////////////////////////////////////////////////////////
//
//            global flag    command flag             parameters (params)
//              __|___        ____|_____             ____|_____
//             /      \      /          \           /          \
//     app-cli --flag=2 open -f thing template /path/to/file /path/to-file
//     \_____/          \__/              \______/
//        |              |                   |
//   application       command             subcommand
//
///////////////////////////////////////////////////////////////////////////////
// Alpha Release

// TODO: Provide validation for conflicting flags and commands defined on
// itialization

// TODO: When using youtube-dl I noticed they had so many flags they needed
// groupings (or categories); so to fit a wider range of use cases we should add
// this functionality (probably minimum argmuments too).

// TODO: Shouldnt render subcommands section if they dont exist

// TODO: Write tests for basic functionality, specifically around the Parse()
// function + Execute. Fix permissions (public vs private) on functions only leaving
// explicitly the functions used by a developer using the library

// TODO: Rewrite the README.md

// TODO: Add ability to access Banner/Spinner (and others) text user interface 
// (TUI) tools from actions.
//          context.CLI.Spinner() 

type App struct {
  Name           string
  Description    string
  Version        Version
  Debug          bool
  Outputs        Outputs
  GlobalFlags    flags
  Commands       commands
  Actions        Actions
}

type CLI struct {
  Version        Version
  Build          Build
  Debug          bool
  Context        *Context
  Outputs        Outputs
  Actions        Actions
  //Locale         string // Not yet implemented
}

func (self CLI) Log(output ...string)   { self.Outputs.Log(DEBUG, output...) }
func (self CLI) Warn(output ...string)  { self.Outputs.Log(WARN, output...) }
func (self CLI) Error(output ...string) { self.Outputs.Log(ERROR, output...) }
func (self CLI) Fatal(output ...string) { self.Outputs.Log(FATAL, output...) }

func New(appDefinition ...App) (cli *CLI, errs []error) {
  // TODO: Clean this up so its not as ugly
  var app App
  if len(appDefinition) == 0 {
    app = App{}
  }else{
    app = appDefinition[0]
  }
  
  // Validation
  errs = append(errs, app.Commands.Validate()...)
  errs = append(errs, app.GlobalFlags.Validate()...)

  if len(errs) != 0 {
     fmt.Println("number of validation errors for flags and commands:", len(errs))
     return cli, errs
  }

  // NOTE: Sensical defaults to avoid error conditions, simplifying library use
  if data.IsBlank(app.Name) {
    app.Name = "app-cli"
  }
  if app.Version.undefined() {
    app.Version = Version{Major: 0, Minor: 1, Patch: 0}
  }
  if len(app.Outputs) == 0 {
    app.Outputs = append(app.Outputs, TerminalOutput())
  }

  // NOTE: Correct alias (short) for flags if more than 1 rune
  //for _, flag := range app.GlobalFlags {
  //  if 1 < len(flag.Alias) {
  //    // TODO: To support localization we will need to handle two byte runes
  //    // in the future
  //    flag.Alias = string(flag.Alias[0])
  //  }
  //}

  // NOTE: If a fallback is not set, we render default help template. 
  if app.Actions.Fallback == nil {
    app.Actions.Fallback = RenderDefaultHelpTemplate
  }

  cli = &CLI{
    Version:  app.Version,
    Outputs:  app.Outputs,
    Actions:  app.Actions,
    Build: Build{
      CompiledAt: time.Now(),
    },
  }

  // TODO: This is going to be troublesome come localization
  if !app.Commands.HasCommand("help") {
    app.Commands = app.Commands.Add(&Command{
      Name: "help",
      Alias: "h",
      Description: "outputs command and flag details",
      Action: HelpCommand,
      Hidden: false,
    })
  }

  if !app.Commands.HasCommand("version") {
    app.Commands = app.Commands.Add(&Command{
      Name: "version",
      Alias: "v",
      Description: "outputs version",
      Action: RenderDefaultVersionTemplate,
      Hidden: false,
    })
  }

  if !app.GlobalFlags.HasFlag("help") {
    app.GlobalFlags = app.GlobalFlags.Add(&Flag{
      Name: "help",
      Alias: "h",
      Description: "outputs command and flag details",
      Hidden: false,
      Action: RenderDefaultHelpTemplate,
    })
  }

  if !app.GlobalFlags.HasFlag("version") {
    app.GlobalFlags = app.GlobalFlags.Add(&Flag{
      Name: "version",
      Alias: "v",
      Description: "outputs version",
      Hidden: false,
      Action: RenderDefaultVersionTemplate,
    })
  }

  // NOTE: Application psuedo-command to store globals and simplify logic
  appCommand := Command{
    Name:        app.Name,
    Description: app.Description,
    Subcommands: app.Commands,
    Flags:       app.GlobalFlags.SetDefaults(),
    Hidden:      true,
    Action:      app.Actions.Fallback,
  }

  cli.Context = &Context{
    CLI:       cli,
    Process:   Process(),
    Commands:  Commands(appCommand),
    Params:    params{},
    Flags:     appCommand.Flags,
    Arguments: Arguments(appCommand),
    Actions:   actions{},
  }

  cli.Context.Command = cli.Context.Commands.First()

  return cli, errs
}

func (self *CLI) Parse(arguments []string) *CLI {
  defer self.benchmark(time.Now(), "benmarking argument parsing")

  // TODO: Need to add the check for required flags (new feature just thought of
  // because seems like it would be useful no? But for this to ever occur we
  // need to not just return error but ERRORS- this is actually nice because we
  // would expect parse to return PARSING ERRORS no?-- which will let us have
  // our validations OH VALIDATIONS HOW I LOVE THE

  for _, argument := range arguments[1:] {
    // Flag parse
    if flagType, ok := HasFlagPrefix(argument); ok {
      argument = flagType.TrimPrefix(argument)
      switch flagType {
      case Short:
        for index, shortFlag := range argument {
          // NOTE: Confirm we are not last && next argument is '=' (61) &&
          if len(argument) != index + 1 && argument[index+1] == 61 { 
            if flag := self.Context.Flag(string(shortFlag)); flag != nil {
              if flagParam := argument[index+2:]; len(flagParam) != 0 {
                flag.Set(flagParam)
              }else{
                // TODO: Is this necessary or redundant?
                flag.SetDefault()
              }
              self.Context.Arguments = self.Context.Arguments.Add(flag)
              break
            }
          }else{
            if flag := self.Context.Flag(string(shortFlag)); flag != nil {
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
        if flag := self.Context.Flag(string(longFlagParts[0])); flag != nil {
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
    for _, flag := range command.Flags {
      if data.IsTrue(flag.Param.value) && flag.Action != nil {
        self.Context.Actions = append(self.Context.Actions, flag.Action)
      }
    }

    if command.Action != nil {
      self.Context.Actions = append(self.Context.Actions, command.Action)
      // NOTE: Break so only first available action is used. Fallback
      // should only run if no actions were defined by commands
      break
    }
  }

  if self.Actions.OnExit != nil {
    self.Context.Actions = self.Context.Actions.Add(self.Actions.OnExit)
  }

  return self
}

func (self *CLI) Execute() {
  defer self.benchmark(time.Now(), "benmarking action execution")
  for _, action := range self.Context.Actions {
    action(self.Context)
  }
}
