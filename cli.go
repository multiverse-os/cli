package cli

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	log "github.com/multiverse-os/cli/log"
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

type Action func(input *Input) error

// TODO: A problem exist with ordering, its not possible to call global option flags at the end, but long as there is no duplication between
// flag levels which would be best avoided anyways for confusion reasons the global option flag should be callable anywhere. this is the expected
// and normal functionality.
type CLI struct {
	Name        string
	Version     Version
	Description string
	Usage       string
	ANSI        bool
	Commands    []Command
	Flags       []Flag

	Logger            log.Logger
	CompiledAt        time.Time
	CompilerSignature string // This will allow developers to provide signed builds that can be verified to prevent tampering
	HideHelp          bool
	HideVersion       bool
	CommandCategories CommandCategories
	WriteMutex        sync.Mutex

	Writer    io.Writer
	ErrWriter io.Writer
	//DefaultAction interface{}
	Hooks map[string]Hook
	// Error Functions
	// TODO: Why not just make these locales and print from standard error log?
	CommandNotFound func()
	ExitErrHandler  func()
	OnUsageError    func()
	DefaultAction   Action
}

func New(cli *CLI) *CLI {
	cli.CompiledAt = time.Now()
	if len(cli.Name) == 0 {
		var err error
		cli.Name, err = filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			cli.Logger.Fatal("Failed to parse executable working directory in default 'Name' attribute assignment.")
		}
	}
	if cli.Version.Undefined() {
		cli.Version = Version{Major: 0, Minor: 1, Patch: 0}
	}
	if cli.Writer == nil {
		cli.Writer = os.Stdout
	}
	if len(cli.Logger.Name) == 0 {
		fmt.Println("cli.Logger.Name: " + cli.Name)
		cli.Logger = log.DefaultLogger(cli.Name, true, true)
	}
	cli.Commands = append(cli.Commands, defaultCommands()...)
	if !cli.HideVersion {
		// TODO: We should just have an init function that loads hidden version and
		// help flags. we can use a 'bool' to say if they are visible or not, same
		// with commands above
		//cli.appendFlag(VersionFlag)
	}
	cli.CommandCategories = CommandCategories{}
	return cli
}

func (self *CLI) Run(arguments []string) (err error) {
	// TODO: Add shell completion code (old code used to be here)

	// TODO: So here, is where we would see if any action is called, and if
	// defaultaction is nil, then we just call help, this avoids any used
	// memory by just applying the functionality in order of operations

	// TODO: So previous version this code is started from, would execute actions
	// located inside the command, but then still run default action. This is
	// really poorly designed. we want to make a rich parsing function that loads
	// up a ACTIVE_MAP then use switch case to go through that ACTIVE_MAP and
	// execute the actions
	// Run default Action

	// TODO: Here we want to build the argument/flag/command/subcommand map by
	// parsing the arguments, then we process it. This will also be where we just
	// run the help command if no DefaultAction is defined. We dont need all this
	// extra logic assigning help to default action, or having special functions
	// for showhelp and close vs show help, we just do it all here and reduce our
	// overall codebase signficiantly

	err = HandleAction(self.DefaultAction)
	if err != nil {
		self.Logger.Error(err)
	}

	return err
}

func HandleAction(action interface{}) error {
	if a, ok := action.(func()); ok {
		a()
		return nil
	} else if a, ok := action.(func() error); ok {
		return a()
	} else if a, ok := action.(func()); ok { // deprecated function signature
		a()
		return nil
	}
	return errInvalidActionType
}

func (self *CLI) VisibleFlags() (visibleFlags []Flag) {
	// TODO the first variable assignement ehre is key, it could be used for building
	// a new map of only visible flags
	for _, flag := range self.Flags {
		visibleFlags = append(visibleFlags, flag)
	}
	return visibleFlags
}

func (self *CLI) HasVisibleFlags() bool {
	return len(self.Flags) > 0
}
