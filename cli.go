package cli

import (
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
// CompilerSignature string // This will allow developers to provide signed builds that can be verified to prevent tampering
type CLI struct {
	Name          string
	Version       Version
	Description   string
	Usage         string
	ANSI          bool
	Commands      []Command
	Flags         []Flag
	Logger        log.Logger
	CompiledAt    time.Time
	WriteMutex    sync.Mutex
	Writer        io.Writer
	ErrWriter     io.Writer
	Hooks         map[string]Hook
	DefaultAction Action
}

func New(cli *CLI) *CLI {
	cli.CompiledAt = time.Now()
	if len(cli.Logger.Name) == 0 {
		cli.Logger = log.DefaultLogger(cli.Name, true, true)
	}
	if len(cli.Name) == 0 {
		var err error
		cli.Name, err = filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			cli.Logger.Fatal("failed to assign 'Name' attribute")
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
	input := LoadInput(self, &Command{}, []*Flag{})
	self.renderUI()

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

	err = self.DefaultAction(input)
	if err != nil {
		self.Logger.Error(err)
	}

	return err
}
