package cli

import (
  "time"

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

// TODO: Load all the all the the hooks, global/fallback, command actions in a
// actions slice (within CLI or context) and then build a fucntion Execute() on
// the actions. So in the end it will be 
//      cli.Parse(os.Args).Execute()

// TODO: Add 

// TODO: Add localization support (we should write a library that can be used by
// both this and the webframework

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

func New(app App) *CLI {
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
