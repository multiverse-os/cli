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

// TODO: Write tests for basic functionality, specifically around the Parse()
// function + Execute. Fix permissions (public vs private) on functions only leaving
// explicitly the functions used by a developer using the library

// TODO: Expand range of the tests so it test more possible conditions to
// guarantee it works when changes are made

// TODO: Dont render categories if only global flags

// TODO: change receiver variable names on methods from self to the convention

// TODO: Rewrite the README.md

// TODO: Add ability to access Banner/Spinner (and others) text user interface 
// (TUI) tools from actions.
//          context.CLI.Spinner() 

// TODO: Ability to use ansii via CLI.Screen.Clear(), or CLI.Text.Blue("test")

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
  MinimumArgs    int
  Locale         string
}

func (self CLI) Log(output ...string)   { self.Outputs.Log(DEBUG, output...) }
func (self CLI) Warn(output ...string)  { self.Outputs.Log(WARN, output...) }
func (self CLI) Error(output ...string) { self.Outputs.Log(ERROR, output...) }
func (self CLI) Fatal(output ...string) { self.Outputs.Log(FATAL, output...) }

// TODO: Get rid of flag actions by simply catching version or help in a generic
// fallback that looks for these flags. This should also help resolve issues
// requiring hardcoding
func New(appDefinition ...App) (cli *CLI, errs []error) {
  // TODO: Clean this up so its not as ugly
  app := App{}
  if len(appDefinition) != 0 {
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

  // NOTE: If a fallback is not set, we render default help template. 
  if app.Actions.Fallback == nil {
    app.Actions.Fallback = HelpCommand
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
      Hidden: true,
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
      Hidden: true,
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
      } else if (len(argument) == 4 && argument == "help") || 
      (len(argument) == 1 && argument == "h") {
        // TODO: Because using help on a subcommand doesnt parse because help is
        // global. And thats how it should work. Version doesn't need this.
        // But I really hate this hardcoding
        helpCommand := self.Context.Commands.Last().Subcommand("help")
        if helpCommand != nil {
          helpCommand.Parent = self.Context.Commands.First()

          self.Context.Commands = self.Context.Commands.Add(helpCommand)
          self.Context.Flags = append(self.Context.Flags, helpCommand.Flags...)

          self.Context.Arguments = self.Context.Arguments.Add(
            self.Context.Commands.First(),
          )

          self.Context.Command = self.Context.Commands.First()
        }
      }else{
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

  if self.Actions.OnStart != nil {
    self.Context.Actions = self.Context.Actions.Add(self.Actions.OnStart)
  }

  var skipCommandAction bool
  for _, command := range self.Context.Commands {
    for _, flag := range command.Flags {
      if flag.Action != nil && data.IsTrue(flag.Param.value) {
        self.Context.Actions = append(self.Context.Actions, flag.Action)
        if flag.Name == "help" {
          skipCommandAction = true
        }
      }
    }
  }

  if !skipCommandAction {
    if 0 < len(self.Context.Commands) {
      command := self.Context.Commands.First()
      if command.Action != nil {
        self.Context.Actions = append(self.Context.Actions, command.Action)
      }
    }
  }

  if self.Actions.OnExit != nil {
    self.Context.Actions = self.Context.Actions.Add(self.Actions.OnExit)
  }

  // NOTE: Before handing the developer using the library the context we put
  // them in the expected left to right order, despite it being easier for us
  // to access in this function in the reverse order.
  self.Context.Arguments = Reverse(self.Context.Arguments)
  self.Context.Commands = ToCommands(Reverse(self.Context.Commands.Arguments()))
  self.Context.Params = ToParams(Reverse(self.Context.Params.Arguments()))

  return self
}

func (self *CLI) Execute() {
  defer self.benchmark(time.Now(), "benmarking action execution")
  for _, action := range self.Context.Actions {
    action(self.Context)
  }
}
