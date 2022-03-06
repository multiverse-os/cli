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

// TODO: Add localization support

// TODO: Fix permissions (public vs private) on functions only leaving
// explicitly the functions used by a developer using the library

// TODO: Reduce flag alias size to 1 character (rune), but command aliases can
// be any length

// TODO: Ensure there are validations on each params, flags, and commands to
//       prevent special characters beyond basics like '-' 'a-z' (except params
//       which should be pretty much everything

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
  GlobalHooks    Hooks
}

type CLI struct {
  Version        Version
  Build          Build
  Debug          bool
  Context        *Context
  Outputs        Outputs
  //Locale         string // Not yet implemented
}

func (self CLI) Log(output ...string)   { self.Outputs.Log(DEBUG, output...) }
func (self CLI) Warn(output ...string)  { self.Outputs.Log(WARN, output...) }
func (self CLI) Error(output ...string) { self.Outputs.Log(ERROR, output...) }
func (self CLI) Fatal(output ...string) { self.Outputs.Log(FATAL, output...) }

func New(app App) *CLI {
  if data.IsBlank(app.Name) {
    app.Name = "app-cli"
  }
  if app.Version.undefined() {
    app.Version = Version{Major: 0, Minor: 1, Patch: 0}
  }
  if data.IsZero(len(app.Outputs)) {
    app.Outputs = append(app.Outputs, TerminalOutput())
  }

  cli := &CLI{
    Version:  app.Version,
    Outputs:  app.Outputs,
    Build: Build{
      CompiledAt: time.Now(),
    },
  }

  appCommand := Command{
    Name:        app.Name,
    Description: app.Description,
    Subcommands: app.Commands,
    Flags:       app.GlobalFlags.SetDefaults(),
    Hidden:      true,
  }

  cli.Context = &Context{
    CLI:     cli,
    Process: Process(),
    Commands: Commands(appCommand),
    Params: params{},
    Flags: appCommand.Flags,
    Arguments: Arguments(appCommand),
  }

  cli.Context.Command = cli.Context.Commands.First()

  // TODO: Take actions+hooks and insert them into the app psuedo-command so
  // they can be consolidated and executed in the Execute() command under
  // actions.

  return cli
}
