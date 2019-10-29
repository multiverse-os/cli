package cli

import (
	"path/filepath"
	"strconv"
	"time"

	argument "github.com/multiverse-os/cli/framework/argument"
	data "github.com/multiverse-os/cli/framework/data"
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
	Routes        map[string]*Command
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
	cli.Routes = map[string]*Command{}

	cli.Command = Command{
		Name:        cli.Name,
		Subcommands: cli.Commands,
		Flags:       cli.Flags,
		Action:      cli.DefaultAction,
	}
	cli.Debug = true
	cli.Build.CompiledAt = time.Now()
	return cli
}

// TODO:
// Another important thing would be migrating to an approach which better
// supports managing a service. It makes sense for a CLI tool to support
// daemonization (even if by subpackage), pid creation, service/process
// management.
func (self *CLI) Run(arguments []string) (*Context, error) {
	defer self.benchmark(time.Now(), "benmarking argument parsing and action execution")
	context := self.Parse(arguments)
	//self.Debug = context.HasFlag("debug")
	if context.HasFlag("version") || context.Command.Name == "version" {
		self.RenderVersionTemplate()
	} else if context.HasFlag("help") || context.Command.Name == "help" || context.Command.Parent == nil || context.CommandDefinition().Action == nil {
		self.RenderHelpTemplate(context.CommandDefinition())
	} else {
		context.CommandDefinition().Action(context)
	}
	return context, nil
}

func (self *CLI) IsFlag(path []string, flagName string) (*Command, *Flag, bool) {
	if 0 < len(path) {
		if command, ok := self.Command.Route(path); ok {
			return command.Flag(flagName)
		} else {
			self.IsFlag(path[:(len(path)-1)], flagName)
		}
	}
	return nil, nil, false
}

// Context Creation, Command Routing, Flag Parsing, and Parameter Parsing
/////////////////////////////////////////////////////////////////////////////
func (self *CLI) Parse(arguments []string) *Context {
	defer self.benchmark(time.Now(), "benmarking parse")
	cwd, executable := filepath.Split(arguments[0])
	context := &Context{
		CLI:        self,
		CWD:        cwd,
		Executable: executable,
		Command: &argument.Command{
			Name:       self.Name,
			Definition: &self.Command,
		},
		Flags:        map[string]*argument.Flag{},
		CommandChain: &argument.Chain{},
		Params:       argument.Params{},
		Args:         arguments[1:],
	}

	for index, arg := range context.Args {
		self.Log(DEBUG, DebugInfo("CLI.parse()"), "attempting to parse the", VarInfo(arg), "at position [", strconv.Itoa(index), "] in the argument chain")
		if flagType, ok := argument.HasFlagPrefix(arg); ok {
			context.ParseFlag(index, flagType, &argument.Flag{Name: arg})
		} else {
			self.Log(DEBUG, "parsing either a command or parameters")
			if command, ok := context.CommandDefinition().Route(append(context.Command.Path(), arg)); ok {
				self.Log(DEBUG, "parsing command:", command.Name)
				context.CommandChain.AddCommand(context.Command)
				context.Command = &argument.Command{
					Parent:     context.CommandChain.Last(),
					Name:       command.Name,
					Definition: &command,
				}
				self.Log(DEBUG, "new command name is:", context.Command.Name)
			} else {
				self.Log(DEBUG, "parsing params")
				for _, param := range context.Args[index:] {
					if flagType, ok := argument.HasFlagPrefix(param); ok {
						self.Log(DEBUG, "parsing params, but it is a flag... loading on last command")
						context.ParseFlag(index, flagType, &argument.Flag{Name: arg})
					} else {
						self.Log(DEBUG, "parsing params, adding to context")
						context.Params.Value = append(context.Params.Value, param)
					}
				}
				break
			}
		}
	}
	return context
}
