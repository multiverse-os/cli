package cli

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
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
// TODO: Ability to have multiple errors, for example we can parse and provide
// all errors at once regarding input so user does not need to trial and
// error to get the information how to fix issues but can instead fix all at
// once and rerun the command.
// NOTE: command is the root command of the command tree which is the name of
// the application. It stores the global flags and functions to hold all the
// global functionality for when commands are not used. This enables us to avoid
// duplicating logic.
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

type Build struct {
	CompiledAt time.Time
	Source     string
	Commit     string
	Signature  string
	Authors    []Author
}

type Author struct {
	PublicKey string
	Name      string
	Email     string
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
	Flags    []Flag
	Commands []Command
	//Errors        []error
}

func New(cli *CLI) *CLI {
	if data.IsBlank(cli.Name) {
		cli.Name = "example"
	}
	//if data.IsBlank(cli.Locale) {}
	if cli.Version.undefined() {
		cli.Version = Version{Major: 0, Minor: 1, Patch: 0}
	}
	if data.IsZero(len(cli.Outputs)) {
		cli.Outputs = append(cli.Outputs, TerminalOutput())
	}

	cli.Command = Command{
		Global:      true,
		Name:        cli.Name,
		Subcommands: cli.Commands,
		Flags:       cli.Flags,
		Action:      cli.DefaultAction,
	}
	cli.Debug = true
	cli.Build.CompiledAt = time.Now()
	return cli
}

func (self *CLI) IsFlag(path []string, flagName string) (*Command, *Flag, bool) {

	if 0 < len(path) {
		if command, ok := self.Command.Subcommand(path[0]); ok {
			return command.Flag(flagName)
		} else {
			self.IsFlag(path[:(len(path)-1)], flagName)
		}
	}
	return nil, nil, false
}

///////////////////////////////////////////////////////////////////////////////
// TODO:                                                                     //
//   Another important thing would be migrating to an approach which better  //
//   supports managing a service. It makes sense for a CLI tool to support   //
//   daemonization (even if by subpackage), pid creation, service/process    //
//   management.                                                             //
///////////////////////////////////////////////////////////////////////////////
func (self *CLI) Parse(arguments []string) (*Context, error) {
	defer self.benchmark(time.Now(), "benmarking argument parsing and action execution")

	cwd, executable := filepath.Split(arguments[0])

	context := &Context{
		CLI:          self,
		CWD:          cwd,
		Command:      &self.Command,
		Executable:   executable,
		GlobalFlags:  map[string]*Flag{},
		CommandFlags: map[string]map[string]*Flag{},
		CommandChain: &Chain{},
		Params:       Params{},
		Args:         arguments[1:],
	}
	context.CommandChain.AddCommand(&self.Command)

	for index, arg := range context.Args {
		self.Log(DEBUG, DebugInfo("CLI.parse()"), "attempting to parse the", VarInfo(arg), "at position [", strconv.Itoa(index), "] in the argument chain")
		if flagType, ok := HasFlagPrefix(arg); ok {
			context.ParseFlag(index, flagType, &Flag{Name: arg})
		} else {

			self.Log(DEBUG, "parsing either a command or parameters")

			if command, ok := context.Command.Subcommand(arg); ok {

				self.Log(DEBUG, "parsing command:", command.Name)
				command.Parent = context.Command

				context.Command = &command

				self.Log(DEBUG, "new command name is:", context.Command.Name)
				context.CommandChain.AddCommand(context.Command)

				//context.CommandChain.AddCommand(context.Command)
				self.Log(DEBUG, "new command name is:", context.Command.Name)
			} else {
				self.Log(DEBUG, "parsing params")
				for _, param := range context.Args[index:] {
					if flagType, ok := HasFlagPrefix(param); ok {
						self.Log(DEBUG, "parsing params, but it is a flag... loading on last command")
						context.ParseFlag(index, flagType, &Flag{Name: arg})
					} else {
						self.Log(DEBUG, "parsing params, adding to context")
						context.Params.Value = append(context.Params.Value, param)
					}
				}
				break
			}
		}
	}

	self.Log(DEBUG, "what is the LAST command name?", context.CommandChain.Last().Name)
	self.Log(DEBUG, "what is the FIRST command name?", context.CommandChain.First().Name)
	self.Log(DEBUG, "what is the CHAIN PATH:", context.CommandChain.Path())

	for _, command := range context.CommandChain.Commands {
		for _, flag := range context.CommandFlags[command.Name] {
			fmt.Println("flag:", flag.Name)
			// TODO: Each name and alias shoudl be added to the map so it can be
			// used by the developer using the library
			flag.Name = strings.Split(flag.Name, ",")[0]
			fmt.Println("flag.value:", flag.Value)
			fmt.Println("flag.default:", flag.Default)
			if len(flag.Value) == 0 {
				fmt.Println("setting default value for flag")
				flag.Value = flag.Default
			}
			context.CommandFlags[command.Name][flag.Name] = flag
		}
	}

	if context.CommandChain.Unselected() {
		context.Command = &Command{
			Parent: context.Command,
			Name:   "help",
		}
	}

	self.Debug = context.HasFlag("debug")
	if context.HasFlag("version") || context.Command.Name == "version" {
		self.RenderVersionTemplate()
	} else if context.HasFlag("help") || context.Command.Name == "help" {
		context.RenderHelpTemplate()
	} else {
		context.Command.Action(context)
	}
	return context, nil

}
