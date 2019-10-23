package cli

import (
	"io"
	"os"
	"time"
	//log "github.com/multiverse-os/cli/framework/log"
	//radix "github.com/multiverse-os/cli/radix"
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

type Action func(context *Context) error

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

// TODO: Should shell be a modificaiton of this, or its own object?
// Output should be a generic thing we write to, this needs to support for
// example writing to both console and log file. And support writing to
// arbitrary locations to be flexible as possible
// TODO: Organize commands into a command tree for better lookup. With the root
// of the tree being the cli name *(This change is pretty important because it
// allows us to move a lot of logic previously duplicated on both Commands and
// CLI to only the command secion)*
// TODO: Ability to have multiple errors, for example we can parse and provide
// all errors at once regarding input so user does not need to trial and
// error to get the information how to fix issues but can instead fix all at
// once and rerun the command.
type CLI struct {
	Name          string
	ArgsRequired  int // For simple scripts, like one that converts a file and requires filename
	Locale        string
	Version       Version
	Description   string
	Commands      []Command
	Flags         []Flag
	CommandTree   *Command
	Build         Build
	Outputs       []io.Writer
	DefaultAction Action
	Errors        []error
	//Router        *radix.Tree
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
		cli.Outputs = append(cli.Outputs, os.Stdout)
	}
	cli.CommandTree = &Command{
		Name:        cli.Name,
		Subcommands: cli.Commands,
		Flags:       cli.Flags,
		Action:      cli.DefaultAction,
	}
	cli.Build.CompiledAt = time.Now()
	return cli
}

func (self *CLI) Run(arguments []string) (err error) {
	context := self.parse(arguments[1:])
	if _, ok := context.Flags["version"]; ok {
		self.renderVersion()
	} else if _, ok = context.Flags["help"]; ok {
		if context.hasNoCommands() {
			// TODO: If the command is help remember that it will need to render
			// command.Parent
			//self.RenderCommandHelp(context.Command())
		} else {
			self.renderApplicationHelp()
		}
	} else if !context.hasNoCommands() {
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
