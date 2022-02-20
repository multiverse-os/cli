package cli

import (
	"time"

	data "github.com/multiverse-os/cli/data"
)

type Action func(context *Context) error

// TODO: Scaffolding code to hasten development.
// https://golang.org/pkg/go/printer/

// Ontology of a command-line interface
///////////////////////////////////////////////////////////////////////////////
//
//            global flag    command flag             parameters (params)
//              __|___        ____|_____             ____|_____
//             /      \      /          \           /          \
//     app-cli --flag=2 open --file=thing template /path/to/file
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


type localization struct {
	Language string
	Locale   string
	Text     map[string]string
}

type Actions struct {
  Global    Action
  Fallback  Action
  OnStart   Action
  OnExit    Action
  // OnExit? Or Close? or this just covered by After?
}

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
	Outputs        Outputs
	Debug          bool // Controls if Debug output writes are skipped
	// At this point almost entirely for API simplicity
	GlobalFlags   flags
	Commands      commands
	//Errors      []error
	//Examples    []Chain
}

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
		Debug:   true,
		Build: Build{
			CompiledAt: time.Now(),
		},
		Actions: Actions{
      Global:   cli.Actions.Global,
      Fallback: cli.Actions.Fallback,
      OnStart:  cli.Actions.OnStart,
      OnExit:   cli.Actions.OnExit,
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
