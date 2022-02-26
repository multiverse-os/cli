package cli

import (
  "time"

  data "github.com/multiverse-os/cli/data"
)


// TODO: Scaffolding code to hasten development.
// https://golang.org/pkg/go/printer/

// Ontology of a command-line interface
///////////////////////////////////////////////////////////////////////////////
//
//            global flag    command flag             parameters (params)
//              __|___        ____|_____             ____|_____
//             /      \      /          \           /          \
//     app-cli --flag=2 open --file=thing template /path/to/file /path/to-file
//     \_____/          \__/              \______/
//        |              |                   |
//   application       command             subcommand
//
///////////////////////////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////////////////////////
// TODO: Ability to have multiple errors, for example we can parse and       //
// provide all errors at once regarding input so user does not need to trial //
// and error to get the information how to fix issues but can instead fix    //
// all at once and rerun the command.                                        //
///////////////////////////////////////////////////////////////////////////////

// TODO: Extend the build aspect of the system. Pull data from last push to the
// public github (it will eventually be our own fucking git hosting and public
// and good). Information about the authors, pgp key (holding email, and such),
// ability to minimize builds or add files. some of the experimental stuff maybe
// added as modules (look back to the chatbot for a good example of
// plugin/module style logic) 
type CLI struct {
  Name           string
  Description    string

  Locale         string // Not yet implemented

  Context        *Context

  Version        Version
  Build          Build
  // TODO: Would like a better solution for this concept, perhaps having to do
  // with the argument interface, through a function implemented on each type,
  // flag, command, and param
  //MinimumArgs    int       // For simple scripts, like one that converts a file and requires filename
  Outputs        Outputs

  Debug          bool

  GlobalFlags    flags
  Commands       commands
  Actions        Actions
  GlobalHooks    Hooks
  //Errors       []error
  //Examples     []Chain
}

type Level uint8

const (
	GlobalLevel Level = iota
	CommandLevel
)

// Helpers
//// logging
func (self CLI) Log(output ...string)   { self.Outputs.Log(DEBUG, output...) }
func (self CLI) Warn(output ...string)  { self.Outputs.Log(WARN, output...) }
func (self CLI) Error(output ...string) { self.Outputs.Log(ERROR, output...) }
func (self CLI) Fatal(output ...string) { self.Outputs.Log(FATAL, output...) }
//// (actions|command|flags|params) chain
func (self CLI) arguments() arguments { return self.Context.Chain.Arguments }
func (self CLI) commands() commands { return self.Context.Chain.Commands }
func (self CLI) flags() flags { return self.Context.Chain.Flags }
func (self CLI) params() params { return self.Context.Chain.Params }

//

// TODO: Move the global flags into the first command in the chain (the root
// command which is the application itself) -- this will allow for much simpler
// processing of flags and actions
// TODO: look at command, then each command in reverse

// TODO: CLI should have spinners, loaders, etc any TUI style things

//    context.CLI.Spinner() 

// TODO: CLI Needs the ability to pull out defined flags 
// TODO: CLI needs the ability to pull out defined commands
// TODO: Need the ability to pull out the action, which should be global and the
// hooks in a actions slice (SHOULD IT?????)

// TODO: Flags renders this kinda obsolete but we ahve to update all associated
// functions. This will temporarily break everything but this is pre-alpha and
// we are getting messy real messy because on the other end of this emss is an
// API we won't be able to touch without stupid levels of time wasting. 
//func (self *CLI) Flags() (flags []*Flag) {
//	for _, flag := range self.GlobalFlags {
//		flags = append(flags, &flag)
//	}
//	return flags
//}

// TODO: Ensure we parse environmental variables
// TODO: Add support for configurations
// TODO: Add support for important paths like ~/.local and ~/.config

func New(cli *CLI) *CLI {

  // TODO: Flag Names & Command Names is validated (and default params?)
  if data.IsBlank(cli.Name) {
    cli.Name = "app-cli"
  }
  if cli.Version.undefined() {
    cli.Version = Version{Major: 0, Minor: 1, Patch: 0}
  }
  if data.IsZero(len(cli.Outputs)) {
    cli.Outputs = append(cli.Outputs, TerminalOutput())
  }

  newCLI := &CLI{
    Name:    cli.Name,
    Version: cli.Version,
    Outputs: cli.Outputs,
    Build: Build{
      CompiledAt: time.Now(),
    },
    GlobalHooks: Hooks{
      BeforeAction: cli.GlobalHooks.BeforeAction,
      AfterAction: cli.GlobalHooks.AfterAction,
    },
    Actions: Actions{
      Global:   cli.Actions.Global,
      Fallback: cli.Actions.Fallback,
    },
  }

  newCLI.Context = &Context{
    CLI:     cli,
    Process: Process(),
    Debug:   false,
    Chain:   &Chain{
      Flags: cli.GlobalFlags,
      Commands: commands([]*Command{&Command{
        Name:        cli.Name,
        Subcommands: cli.Commands,
        Flags:       cli.GlobalFlags,
      }}),
    },
  }

  return newCLI
}
