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

// TODO: Originally only Run() was supported requiring ALL logic to be declared
// in the Run() command. But overtime we decided it would be better to provide
// more flexibility to empower developers to organize their structure and simply
// provide a variety of useful generally useful tools regardless of size or
// structure (and pushing anything too specific into a subpackage).
//
// The start of this is returning the context, so it can be used in the case the
// application continues.
//
// I would prefer if we break up each of the three major configuration loading
// steps so they can be optional or overriden. Environmental Load, Configuration
// Load and finally Argument parsing.
// Then once the developer has the desired state loaded, they can proceed to
// execution based on whatever conditions in the configuration or state of data
// folder.
// Another important thing would be migrating to an approach which better
// supports managing a service. It makes sense for a CLI tool to support
// daemonization (even if by subpackage), pid creation, service/process
// management.
//
// TODO: Antoher key concept of the new design is to re-imagine the
// initializaiton process and make building the tree more like declaring a web
// API via a router. Attach flags to commands using the function name to declare
// their datatype.

// TODO: Run COULD provide ability to stack all basic functionality into one
// function to simplify execution, but keep above in mind when moving forward.
//
// TODO: Also the API is getting close enough to being designed it is now worth
// starting to test the framework as we get closer to freezing the API.
//
// Testing should start with external library testing for the specific purpose
// of working with the API, and shaping it with the tests. Then finally internal
// testing should be done to confirm the innerworkings of the framework work as
// expected and can be confirmed to continue to work after changes.
func (self *CLI) Run(arguments []string) (*Context, error) {
	defer self.benchmark(time.Now(), "benmarking argument parsing and action execution")
	context := self.Parse(arguments)
	//if _, ok := context.Flags["version"]; ok {
	//	self.renderVersion()
	//} else if _, ok = context.Flags["help"]; ok {
	//	//if context.hasNoCommands() {
	//	//	// TODO: If the command is help remember that it will need to render
	//	//	// command.Parent
	//	//	//self.RenderCommandHelp(context.Command())
	//	//} else {
	//	//	self.renderApplicationHelp()
	//	//}
	//} else if true == false { //!context.hasNoCommands() {
	//	//err = context.Command().Action(context)
	//} else {
	//	self.renderApplicationHelp()
	//	err = self.DefaultAction(context)
	//}
	// Use outputs writer and make a method on CLI to do that
	//if err != nil {
	//	self.Logger.Error(err)
	//}

	//if command, ok := self.Command.Route(context.Command.Path()); ok {
	//	self.RenderHelpTemplate(command)
	//	command.Action(context)
	//}
	self.RenderVersionTemplate()
	context.Command.Action.(Action)(context)
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
			Action:     self.Command.Action,
			Definition: self.Command,
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
			// TODO: If we had the ability to route from the last command, we could reduce repeated logic, the thing that is preventing this is that calling in the Command from argument.Command would be a cyclic import
			// to avoid that we could maintain a chain of definitions then pull them out using index. It may be worth migrating the chain code into the main program for this purpose
			if command, ok := context.Command.Definition.(Command).Route(append(context.Command.Path(), arg)); ok {
				context.CommandChain.AddCommand(context.Command)
				context.Command = &argument.Command{
					Name:       command.Name,
					Action:     command.Action,
					Definition: command,
				}
			} else {
				for _, param := range context.Args[index:] {
					if flagType, ok := argument.HasFlagPrefix(param); ok {
						context.ParseFlag(index, flagType, &argument.Flag{Name: arg})
					} else {
						context.Params.Value = append(context.Params.Value, param)
					}
				}
				break
			}
		}
	}
	return context
}
