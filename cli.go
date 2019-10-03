package cli

import (
	"io"
	"os"
	"path/filepath"
	"time"

	log "github.com/multiverse-os/cli/log"
	//radix "github.com/multiverse-os/cli/radix"
)

type Action func(context *Context) error

// TODO: It would be great to impelement a middleware like system to
// make CLI programming similar to web programming. Reusing these conceepts
// should make it more familiar and easier to transpose code

// TODO: Decide if flags should be segregated into global flags and
// command flags
type Context struct {
	CLI        *CLI
	Command    Command
	Subcommand Command
	Flags      map[string]Flag
	Args       []string
}

type Build struct {
	CompiledOn time.Time
	Source     string
	Signature  string
}

// TODO: Should shell be a modificaiton of this, or its own object?
type CLI struct {
	Name          string
	Version       Version
	Description   string
	Usage         string
	Flags         []Flag
	Commands      []Command
	Build         Build
	Logger        log.Logger
	Writer        io.Writer
	DefaultAction Action
	//Router        *radix.Tree
}

func New(cli *CLI) *CLI {
	cli.Build.CompiledOn = time.Now()
	if len(cli.Logger.Name) == 0 {
		cli.Logger = log.DefaultLogger(cli.Name, true, true)
	}
	if len(cli.Name) == 0 {
		var err error
		cli.Name, err = filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			cli.Logger.Fatal(errFailedNameAssignment.Error())
		}
	}
	if cli.Version.Undefined() {
		cli.Version = Version{Major: 0, Minor: 1, Patch: 0}
	}
	if cli.Writer == nil {
		cli.Writer = os.Stdout
	}
	cli.Flags = append(cli.Flags, defaultFlags()...)
	cli.Commands = append(cli.Commands, defaultCommands()...)
	return cli
}

func (self *CLI) visibleCommands() (commands []Command) {
	for _, command := range self.Commands {
		if command.Visible() {
			commands = append(commands, command)
		}
	}
	return commands
}

func (self *CLI) visibleFlags() (flags []Flag) {
	for _, flag := range self.Flags {
		if flag.Visible() {
			flags = append(flags, flag)
		}
	}
	return flags
}

func (self *CLI) isFlag(flagName string) (bool, Flag) {
	for _, flag := range self.Flags {
		if flag.Is(flagName) {
			return true, flag
		}
	}
	return false, Flag{}
}

func (self *CLI) isCommand(commandName string) (bool, Command) {
	for _, command := range self.Commands {
		if command.Is(commandName) {
			return true, command
		}
	}
	return false, Command{}
}

func (self *CLI) isSubcommand(command Command, subcommandName string) (bool, Command) {
	for _, subcommand := range command.Subcommands {
		if subcommand.Is(subcommandName) {
			return true, subcommand
		}
	}
	return false, Command{}
}

func (self *CLI) isCommandFlag(command Command, flagName string) (bool, Flag) {
	for _, flag := range command.Flags {
		if flag.Is(flagName) {
			return true, flag
		}
	}
	return false, Flag{}
}

func (self *CLI) Run(arguments []string) (err error) {
	context := self.parse(arguments[1:])
	if _, ok := context.Flags["version"]; ok {
		self.RenderVersion()
	} else if _, ok = context.Flags["help"]; ok {
		if context.Command.NotEmpty() {
			self.RenderCommandHelp(context.Command)
		} else {
			self.RenderHelp()
		}
	} else if context.Command.NotEmpty() {
		err = context.Command.Action(context)
	} else {
		self.RenderHelp()
		err = self.DefaultAction(context)
	}

	if err != nil {
		self.Logger.Error(err)
	}

	return err
}
