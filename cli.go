package cli

import (
	"path/filepath"
	"strconv"
	"strings"
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
	//Printers       - like Debug() that will put function in and maybe .Value()
	//to chain values
	Debug bool // Controls if Debug output writes are skipped
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
		Name:        cli.Name,
		Subcommands: cli.Commands,
		Flags:       cli.Flags,
		Action:      cli.DefaultAction,
	}
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
	if command, ok := self.Command.Route(context.Command.Path()); ok {
		self.RenderHelpTemplate(command)
	}
	return context, nil
}

//
// Context Creation, Command Routing, Flag Parsing, and Parameter Parsing
/////////////////////////////////////////////////////////////////////////////
func (self *CLI) Parse(arguments []string) *Context {
	cwd, executable := filepath.Split(arguments[0])
	context := &Context{
		CLI:        self,
		CWD:        cwd,
		Executable: executable,
		Command: &argument.Command{
			Arg: self.Name,
		},
		Flags:         map[string]*argument.Flag{},
		ArgumentChain: argument.ParseChain(arguments[1:]),
		Args:          arguments[1:],
	}

	//flagGroup := argument.Flags{}
	for index, arg := range context.Args {
		self.Log(DEBUG, DebugInfo("CLI.parse()"), "attempting to parse the", VarInfo(arg), "at position [", strconv.Itoa(index), "] in the argument chain")

		//if strings.HasPrefix(arg, shortFlag) && data.IsLessThan(1, len(arg)) {
		//	if flag, ok := context.parseFlag(arg); ok {
		//		flagGroup.addFlag(flag)
		//	}
		//} else {
		//	if command, ok := self.command.HasRoute(append(context.CommandPath, arg)); ok {
		//		inputCommand := context.AddCommand(command)
		//		inputCommand.addSubcommandTree(command.Subcommands)
		//		if !data.IsZero(len(*flagGroup)) {
		//			inputCommand.addFlags(flagGroup)
		//			flagGroup.reset()
		//		}
		//		context.addCommand(inputCommand)
		//	} else {
		//		for _, param := range arguments[index:] {
		//			if strings.HasPrefix(param, shortFlag) {
		//				if flag, ok := context.parseFlag(param); ok {
		//					flagGroup.addFlag(flag)
		//				}
		//			} else {
		//				// TODO: Need parameter init datatype declaration to do more with
		//				// parameter otherwise its going to be simple string slice
		//				context.Params = append(context.Params, param)
		//			}
		//		}
		//		if !flagGroup.isEmpty() {
		//			context.Command.addFlags(flagGroup)
		//			flagGroup.reset()
		//		}
		//		break
		//	}
		//}
	}
	return context
}

//
// Ultra Minimal Multiple Output Logging System
///////////////////////////////////////////////////////////////////////////////
// Basic Levels, Debug Mode, Errors, Warnings, Fatals that exit, and flexible
// access so it doesn't get in the way.
//
// Have your own logger? No problem, just append it's io.Writer to CLI.Output.
// Want logging to Terminal & Logfile? No problem, append both to CLI.Output.
// Want both and output to a website? No problem.
func (self *CLI) Output(text ...string) {
	self.Outputs.Write(strings.Join(text, " "))
}

func (self *CLI) Log(level LogLevel, text ...string) {
	self.Outputs.Log(level, strings.Join(text, " "))
}
