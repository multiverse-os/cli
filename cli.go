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
	Locale         string
	Version        Version
	Build          Build
	RequiredArgs   int       // For simple scripts, like one that converts a file and requires filename
	Command        Command   // Base command that represents the CLI itself
	ParamType      data.Type // Filename types should be able to define extension for autcomplete
  Actions        Actions
  GlobalHooks    Hooks
	Outputs        Outputs
	Debug          bool
	GlobalFlags    flags
	Commands       commands
	//Errors       []error
	//Examples     []Chain
}

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

func New(cli *CLI) *CLI {
	if data.IsBlank(cli.Name) {
		cli.Name = "app-cli"
	}
	if cli.Version.undefined() {
		cli.Version = Version{Major: 0, Minor: 1, Patch: 0}
	}
	if data.IsZero(len(cli.Outputs)) {
		cli.Outputs = append(cli.Outputs, TerminalOutput())
	}

	return &CLI{
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
		Command: Command{
			Name:        cli.Name,
			Subcommands: cli.Commands,
      // TODO: UNless the global flags append as they parse which they may this
      // may be inadequate (TESTS LOL righta fter this massive change ins
      // trcuture tests for realies!) 
			Flags:       cli.GlobalFlags,
		},
	}
}
