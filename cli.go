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
type Directories struct {
	Working string
	Data    string
	Cache   string
}

type Localisation struct {
	Language string
	Locale   string
	Text     map[string]string
}

type CLI struct {
	Name          string
	Description   string
	Locale        string
	Version       Version
	Build         Build
	RequiredArgs  int       // For simple scripts, like one that converts a file and requires filename
	Command       Command   // Base command that represents the CLI itself
	ParamType     data.Type // Filename types should be able to define extension for autcomplete
	DefaultAction Action
	Outputs       Outputs
	Debug         bool // Controls if Debug output writes are skipped
	// At this point almost entirely for API simplicity
	GlobalFlags []Flag
	Commands    []Command
	//Errors        []error
	//Examples []Chain
}

func (self *CLI) Flags() (flags []*Flag) {
	for _, flag := range self.GlobalFlags {
		flags = append(flags, &flag)
	}
	return flags
}

func New(cli *CLI) *CLI {
	if data.IsBlank(cli.Name) {
		cli.Name = "example"
	}
	if cli.Version.undefined() {
		cli.Version = Version{Major: 0, Minor: 1, Patch: 0}
	}
	if data.IsZero(len(cli.Outputs)) {
		cli.Outputs = append(cli.Outputs, TerminalOutput())
	}

	// Reader:      os.Stdin,
	// Writer:      os.Stdout,
	// 		fmt.Println(a.Writer, "thing")
	// ErrWriter:   os.Stderr,

	return &CLI{
		Name:    cli.Name,
		Version: cli.Version,
		Outputs: cli.Outputs,
		Debug:   true,
		Build: Build{
			CompiledAt: time.Now(),
		},
		Command: Command{
			Name:        cli.Name,
			Subcommands: cli.Commands,
			Flags:       cli.GlobalFlags,
		},
		DefaultAction: cli.DefaultAction,
	}
}
