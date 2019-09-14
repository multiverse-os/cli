package cli

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	log "github.com/multiverse-os/cli/log"
	//radix "github.com/multiverse-os/cli/radix"
)

/// TASKS /////////////////////////////////////////////////////////////////////
// TODO: CLI Framework should also handle:
//       * Configuration, CLI, and Command-line based configuration loading
//       * Load output from template
//       * Put what was context into a trie, to quickly pull data out of memory,
//       and to handle auto completion.
//       * Validation of user input
//       * Ability to define config and local data folders
//       * Print data in a variety of ways, such as: Spark Graphs, Tree, List
///////////////////////////////////////////////////////////////////////////////

type Action func(context *Context) error

// TODO: A problem exist with ordering, its not possible to call global option flags at the end, but long as there is no duplication between
// flag levels which would be best avoided anyways for confusion reasons the global option flag should be callable anywhere. this is the expected
// and normal functionality.
// CompilerSignature string // This will allow developers to provide signed builds that can be verified to prevent tampering
type SoftwareBuild struct {
	Checksum   string
	Signature  string
	CompiledAt time.Time
}

// TODO: Should shell be a modificaiton of this, or its own object?
type CLI struct {
	Name          string
	Version       Version
	Description   string
	Usage         string
	ANSI          bool
	Flags         []Flag
	Commands      []Command
	Build         SoftwareBuild
	Logger        log.Logger
	WriteMutex    sync.Mutex
	Writer        io.Writer
	ErrorWriter   io.Writer
	DefaultAction Action
	//Router        *radix.Tree
}

func New(cli *CLI) *CLI {
	cli.Build.CompiledAt = time.Now()
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
	return cli
}

func (self *CLI) Run(arguments []string) (err error) {
	context := self.parse(arguments[1:])

	fmt.Println("command:", context.Command)
	fmt.Println("subcommand:", context.Subcommand)
	for _, flag := range context.Flags {
		fmt.Println("flag:")
		fmt.Println("  name:", flag.Name)
		fmt.Println(" value:", flag.Value)
	}

	self.renderHelp()

	err = self.DefaultAction(context)
	if err != nil {
		self.Logger.Error(err)
	}

	return err
}

type Context struct {
	CLI        *CLI
	Command    *Command
	Subcommand *Command
	Flags      []*Flag
}

var VersionFlag Flag = Flag{
	Name:    "version",
	Aliases: []string{"v"},
	Usage:   "Print version",
	Hidden:  true,
}

var HelpFlag Flag = Flag{
	Name:    "help",
	Aliases: []string{"h"},
	Usage:   "Print help text",
	Hidden:  true,
}

func defaultCommands() []Command {
	return []Command{
		Command{
			Hidden:  true,
			Name:    "help",
			Aliases: []string{"h"},
			Usage:   "List of available commands or details for a specified command",
			//ArgsUsage: "[command]",
			//Subcommands: InitSubcommands(),
			Action: func() error {
				// TODO: Args need to be loaded into context so its accessible
				//args := c.Args()
				//if args.Present() {
				//	return ShowCommandHelp(c, args.First())
				//}
				//ShowCLIHelp(c)
				return nil
			},
		},
	}
}

func (self Command) InitSubcommands() []Command {
	return []Command{
		Command{
			Name:    "help",
			Aliases: []string{"h"},
			Usage:   "List of available commands or details for a specified command",
			//ArgsUsage:     "[command]",
			ParentCommand: &self,
			Action: func() error {
				// TODO: Fix this because this is all leading to massive bloat
				//args := c.Args()
				//if args.Present() {
				//	return ShowCommandHelp(c, args.First())
				//}
				//return ShowSubcommandHelp(c)
				return nil
			},
		},
	}
}
