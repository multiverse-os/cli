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

type Context struct {
	CLI        *CLI
	Command    Command
	Subcommand Command
	Flags      []Flag
}

type SoftwareBuild struct {
	CompiledOn time.Time
}

// TODO: Should shell be a modificaiton of this, or its own object?
type CLI struct {
	Name          string
	Version       Version
	Description   string
	Usage         string
	Flags         []Flag
	Commands      []Command
	Build         SoftwareBuild
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
	cli.Flags = defaultFlags()
	cli.Commands = defaultCommands()
	return cli
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

func (self *CLI) Run(arguments []string) (err error) {
	context := self.parse(arguments[1:])

	if !context.Command.isEmpty() {
		err = context.Command.Action(context)

	} else {
		self.renderHelp()
		err = self.DefaultAction(context)
	}

	if err != nil {
		self.Logger.Error(err)
	}

	return err
}
