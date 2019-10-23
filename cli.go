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
	Name              string
	RequiredArguments int // For simple scripts, like one that converts a file and requires filename
	Locale            string
	Version           Version
	Description       string
	command           *Command // Base command that represents the CLI itself
	Commands          []Command
	Flags             []Flag
	Build             Build
	Outputs           []Output
	DefaultAction     Action
	Params            DataType
	ParamFileExt      string
	Errors            []error
}

// TODO: We could try a radix tree that is loaded with the commands. Iterating
// through each row edge first, and assigning values using
// command1.command2.command3. Then we take our path and join with . and do a
// prefix search. We can try that later and see if the preformance gain is worth
// the extra overhead but this is not terrible, its technically a bread-first
// search
// NOTE: Public to allow essentially re-running the application without needing to
// start a new process
func (self *Command) Route(path []string) (*Command, bool) {
	if len(path) <= 1 {
		return self, true
	} else {
		command := self
		for _, name := range path[1:] {
			if subcommand, ok := command.Subcommand(name); ok {
				command = &subcommand
			} else {
				return nil, false
			}
		}
		return command, true
	}
	return nil, false
}

func New(cli *CLI) *CLI {
	if IsBlank(cli.Name) {
		cli.Name, _ = os.Executable()
	}
	if IsBlank(cli.Locale) {

	}
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
	cli.command = &Command{
		Name:        cli.Name,
		Subcommands: cli.Commands,
		Flags:       cli.Flags,
		Action:      cli.DefaultAction,
	}
	cli.Build.CompiledAt = time.Now()
	return cli
}

func (self *CLI) Run(arguments []string) (err error) {
	defer self.benchmark(time.Now())
	context := self.parse(arguments)
	if _, ok := context.Flags["version"]; ok {
		self.renderVersion()
	} else if _, ok = context.Flags["help"]; ok {
		//if context.hasNoCommands() {
		//	// TODO: If the command is help remember that it will need to render
		//	// command.Parent
		//	//self.RenderCommandHelp(context.Command())
		//} else {
		//	self.renderApplicationHelp()
		//}
	} else if true == false { //!context.hasNoCommands() {
		//err = context.Command().Action(context)
	} else {
		self.renderApplicationHelp()
		err = self.DefaultAction(context)
	}
	// Use outputs writer and make a method on CLI to do that
	//if err != nil {
	//	self.Logger.Error(err)
	//}
	return err
}

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

// TODO: implement a Fprintf style Output for better and more coonsistent Output
// usage. Right now its a bit tedious too use but if it takes interface{} and
// covnerts them automatically, it would become more natural. Eventually, should
// also define a theme, then all softwware in a collection can use a consistent
// theme. And finally icon set package would be great, even if just making it
// easier to access the basic UTF-8 ones
// TODO: Fact that color functions CANT receive anything but string is very bad,
// makes it difficult to use the library effetively
func (self *CLI) Debug(text ...string)         { self.Log(DEBUG, color.Blue(strings.Join(text, " "))) }
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
