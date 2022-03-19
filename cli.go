package cli

import (
  "time"
  "os"
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
//     app-cli --flag=2 open -f=thing template /path/to/file /path/to-file
//     \_____/          \__/              \______/
//        |              |                   |
//   application       command             subcommand
//
///////////////////////////////////////////////////////////////////////////////
// Alpha Release

// TODO: Provide validation for conflicting flags and commands defined on
// itialization

// TODO: Write tests for basic functionality, specifically around the Parse()
// function + Execute

// TODO: Fix permissions (public vs private) on functions only leaving
// explicitly the functions used by a developer using the library

// TODO: Rewrite the README.md

///////////////////////////////////////////////////////////////////////////////
// TODO: Scaffolding code to hasten development.
// https://golang.org/pkg/go/printer/

// TODO: Ability to have multiple errors, for example we can parse and
// provide all errors at once regarding input so user does not need to trial
// and error to get the information how to fix issues but can instead fix
// all at once and rerun the command.

// TODO: Extend the build aspect of the system. Pull data from last push to the
// public github (it will eventually be our own fucking git hosting and public
// and good). Information about the authors, pgp key (holding email, and such),
// ability to minimize builds or add files. some of the experimental stuff maybe
// added as modules (look back to the chatbot for a good example of
// plugin/module style logic) 

// TODO: Add ability to access Banner/Spinner (and others) text user interface 
// (TUI) tools from actions.
//          context.CLI.Spinner() 

// TODO: Ensure we parse environmental variables

// TODO: Add support for configuration (flag < env < file)
//         Support ~/.local and ~/.config

// TODO: Autocomplete via tab defined during initalization

// TODO: Switch CLI to App, and get rid of New(), switch it to a more sensical
// Chaining definitions, and make it more similar to Gin. So webframework
// developers will have a super easy time learning this CLI framewoork
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

func New(appDefinition ...App) *CLI {
  var app App
  switch len(appDefinition) {
  case 0:
    app = App{}
  default: 
    app = appDefinition[0]
  }

  // NOTE: Sensical defaults to avoid error conditions, simplifying library use
  if data.IsBlank(app.Name) {
    app.Name = "app-cli"
  }
  if app.Version.undefined() {
    app.Version = Version{Major: 0, Minor: 1, Patch: 0}
  }
  if data.IsZero(len(app.Outputs)) {
    app.Outputs = append(app.Outputs, TerminalOutput())
  }
  // NOTE: Correct alias (short) for flags if more than 1 rune
  for _, flag := range app.GlobalFlags {
      if 1 < len(flag.Alias) {
        // TODO: To support localization we will need to handle two byte runes
        // in the future
        flag.Alias = string(flag.Alias[0])
      }
  }

  cli := &CLI{
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

  // TODO: How can we attach actions to these in order to avoid any need to
  // hardcode ANYTHING

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
  }

  cli.Context.Command = cli.Context.Commands.First()

  return cli
}

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
            if flag := self.Context.Flag(string(flagAlias)); flag != nil {
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

  // NOTE: Put the actions into a chain of action in the order they should be
  // ran
  if self.Actions.OnStart != nil {
    self.Context.Actions = self.Context.Actions.Add(self.Actions.OnStart)
  }

  for _, command := range self.Context.Commands.Reverse() {
    for _, flag := range command.Flags {
      if data.IsTrue(flag.Param.value) && flag.Action != nil {
        self.Context.Actions = append(self.Context.Actions, flag.Action)
      }
    }

    if command.Action != nil {
      self.Context.Actions = append(self.Context.Actions, command.Action)
      break
    }
    // NOTE: Break so only first available action is used. Fallback
    // should only run if no actions were defined by commands
  }

  if self.Actions.OnExit != nil {
    self.Context.Actions = self.Context.Actions.Add(self.Actions.OnExit)
  }

  return self.Context, []error{}
}
