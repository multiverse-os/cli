package cli

import (
	"fmt"
	"os"
	"strings"
	"time"

	color "github.com/multiverse-os/cli/framework/terminal/ansi/color"
)

// Ontology of a command-line interface
///////////////////////////////////////////////////////////////////////////////
//
//            global flag    command flag             parameters
//              __|___        ____|_____             ____|_____
//             /      \      /          \           /          \
//     app-cli --flag=2 open --file=thing template /path/to/file
//     \_____/          \__/              \______/
//        |              |                   |
//   application       command             subcommand

// TODO: It would be great to impelement a middleware like system to
// make CLI programming similar to web programming. Reusing these conceepts
// should make it more familiar and easier to transpose code
// TODO: Provide a way to register an RSS feed that can be used for checking for
// updates.
type Build struct {
	CompiledAt time.Time
	//Signature  string
	//Source     string
}

// TODO: Ability to have multiple errors, for example we can parse and provide
// all errors at once regarding input so user does not need to trial and
// error to get the information how to fix issues but can instead fix all at
// once and rerun the command.
// NOTE: command is the root command of the command tree which is the name of
// the application. It stores the global flags and functions to hold all the
// global functionality for when commands are not used. This enables us to avoid
// duplicating logic.
type CLI struct {
	Name          string
	Description   string
	Locale        string
	Version       Version
	Build         Build
	RequiredArgs  int      // For simple scripts, like one that converts a file and requires filename
	Command       *Command // Base command that represents the CLI itself
	ParamType     DataType // Filename types should be able to define extension for autcomplete
	DefaultAction Action
	Outputs       []Output
	DebugMode     bool // Controls if Debug output writes are skipped
	//Errors        []error
}

func New(cli *CLI) *CLI {
	//if IsBlank(cli.Name) {}
	//if IsBlank(cli.Locale) {}
	// TODO: Migrate to a system that just lets us add logger as one of the
	// outputs, enabling outputing to x number of locations which can easily be a
	// logfile in addition to stdout
	//if len(cli.Logger.Name) == 0 {
	//	cli.Logger = log.DefaultLogger(cli.Name, true, true)
	//}
	if cli.Version.undefined() {
		cli.Version = Version{Major: 0, Minor: 1, Patch: 0}
	}
	if IsZero(len(cli.Outputs)) {
		cli.Outputs = append(cli.Outputs, TerminalOutput())
	}
	cli.Command = &Command{
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
func (self *CLI) Run(arguments []string) (context *Context, err error) {
	defer self.benchmark(time.Now())
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
	return context, err
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
		Command: &inputCommand{
			Name: self.Name,
		},
		CommandPath: []string{self.Name},
		Flags:       map[string]*inputFlag{},
		Args:        arguments[1:],
	}

	flagGroup := newFlagGroup()
	for index, argument := range context.Args {
		self.Debug(debugInfo("CLI.parse()"), "attempting to parse the", varInfo("argument", argument))
		if argument[0] == shortFlag[0] && IsLessThan(1, len(argument)) {
			if flag, ok := context.parseFlag(argument); ok {
				flagGroup.addFlag(flag)
			}
		} else {
			if command, ok := self.command.Route(append(context.CommandPath, argument)); ok {
				inputCmd := newInputCommand(context.Command, command.Name)
				inputCmd.addSubcommandTree(command.Subcommands)
				if !IsZero(len(*flagGroup)) {
					inputCmd.addFlags(flagGroup)
					flagGroup.reset()
				}
				context.addCommand(inputCmd)
			} else {
				for _, param := range arguments[index:] {
					if param[:1] == shortFlag {
						if flag, ok := context.parseFlag(param); ok {
							flagGroup.addFlag(flag)
						}
					} else {
						// TODO: Need parameter init datatype declaration to do more with
						// parameter otherwise its going to be simple string slice
						context.Params = append(context.Params, param)
					}
				}
				if !flagGroup.isEmpty() {
					context.Command.addFlags(flagGroup)
					flagGroup.reset()
				}
				break
			}
		}
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
//
// TODO: implement a Fprintf style Output for better and more coonsistent Output
// usage. Right now its a bit tedious too use but if it takes interface{} and
// covnerts them automatically, it would become more natural. Eventually, should
// also define a theme, then all softwware in a collection can use a consistent
// theme. And finally icon set package would be great, even if just making it
// easier to access the basic UTF-8 ones
// TODO: Fact that color functions CANT receive anything but string is very bad,
// makes it difficult to use the library effetively
func (self *CLI) Output(text string) {
	for _, output := range self.Outputs {
		output.Write(text + "\n")
	}
}

func (self *CLI) Log(level LogLevel, text string) {
	for _, output := range self.Outputs {
		output.Log(level, text+"\n")
	}
}

func (self *CLI) Debug(text ...string) {
	if !self.DebugMode {
		self.Log(DEBUG, color.Blue(strings.Join(text, " ")))
	}
}

func (self *CLI) Warn(text ...string)          { self.Log(WARNING, color.Silver(strings.Join(text, " "))) }
func (self *CLI) Error(text string, err error) { self.Log(ERROR, color.Silver(text+": "+err.Error())) }
func (self *CLI) Fatal(text string, err error) {
	self.Log(FATAL, color.Red(text+": "+err.Error()))
	os.Exit(1)
}

func (self *CLI) benchmark(startedAt time.Time) {
	elapsed := time.Since(startedAt)
	self.Debug(color.Green("cli command parse and action execution completed in [ " + fmt.Sprintf("%s", elapsed) + " ]"))
}
