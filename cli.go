package cli

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	log "github.com/multiverse-os/cli/log"
	text "github.com/multiverse-os/cli/text"
)

// TODO: Support a slice of functions or map of functions for Before and After, so we can have several functions ran before and after any given
// function command subcommand and so on for more complex functionality and modularization of code

// TODO: A problem exist with ordering, its not possible to call global option flags at the end, but long as there is no duplication between
// flag levels which would be best avoided anyways for confusion reasons the global option flag should be callable anywhere. this is the expected
// and normal functionality.

// TODO: Is this redudant between usage text and description?

// TODO: Move all text into locales so we can support localization
// Text to override the USAGE section of help
// Description of the program argument format.

// TODO: No better name than "argsusage"? Becauswe I have no idea what that means

// TODO: Category concept doesnt seem to be used really. Shouldn't be generic "Category" unless its generic, its not.
type CLI struct {
	Name             string
	Version          Version
	Description      string
	NoANSIFormatting bool
	Usage            string
	UsageText        string
	ArgsUsage        string
	// TODO: Store commands and subcommands in a tree object and get rid of this current structure

	Commands        Commands
	Subcommands     Commands
	Flags           map[]Flag

	CommandMap map[string]*Command

	Logger            log.Logger
	CompiledOn        time.Time
	HideHelp          bool
	HideVersion       bool
	CommandCategories CommandCategories
	WriteMutex        sync.Mutex
	Writer            io.Writer
	ErrWriter         io.Writer
	// Functions
	//////////////////////////////////////////////////////////////////////////////
	DefaultAction interface{}
	Hooks         map[string]Hook
	// Error Functions
	// TODO: Why not just make these locales and print from standard error log?
	CommandNotFound CommandNotFoundFunc
	ExitErrHandler  ExitErrHandlerFunc
	OnUsageError    OnUsageErrorFunc
}

func (self *CLI) PrintBanner() {
	fmt.Println(text.Header(self.Name) + "  " + text.Strong("v"+self.Version.String()))
	fmt.Println(text.Light(text.Repeat("=", 80)))
}

// Setup and New dont seem to have any reason to be separate
func New(cli *CLI) *CLI {
	// TODO: Should just handle compile time and such together with hashing and signatures
	//cmd.CompiledOn = time.Now()

	// TODO: Parse ARGs here! So we can use it for nil name assignment etc

	// Default to same name 'go build' uses for executable: the working directory name
	if cli.Name == "" {
		var err error
		cli.Name, err = filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			cli.Logger.Fatal("Failed to parse executable working directory in default 'Name' attribute assignment.")
		}
	}

	if cli.Version.Undefined() {
		cli.Version = Version{Major: 0, Minor: 1, Patch: 0}
	}

	// TODO: I really like this concept, where we could have theoritical different
	// places to write the content. Perhaps even support piping it to HTML or
	// other creative things, but that should be optional and modular, we lay the
	// groundwork though by having this
	// TODO: Add support for 'nohup' like functionality to output all stdout to text file
	// TODO: How can we merge these two? It should be possible and could work very well
	if cli.Writer == nil {
		cli.Writer = os.Stdout
	}
	if cli.Logger.Name == "" {
		fmt.Println("cli.Logger.Name: " + cli.Name)
		cli.Logger = log.DefaultLogger(cli.Name, true, true)
	}
	cli.Commands = InitCommands()
	if !cli.HideVersion {
		// TODO: We should just have an init function that loads hidden version and
		// help flags. we can use a 'bool' to say if they are visible or not, same
		// with commands above
		//cli.appendFlag(VersionFlag)
	}
	cli.CommandCategories = CommandCategories{}
	cli.CommandMap = make(map[string]*Command)
	for _, command := range cli.Commands {
		cli.AddCommandToMap(command)
	}
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

	err = HandleAction(self.DefaultAction, NewContext(self))
	if err != nil {
		self.Logger.Error(err)
	}

	return err
}

func HandleAction(action interface{}, context *Context) error {
	if a, ok := action.(ActionFunc); ok {
		return a(context)
	} else if a, ok := action.(func(*Context) error); ok {
		return a(context)
	} else if a, ok := action.(func(*Context)); ok { // deprecated function signature
		a(context)
		return nil
	}
	return errInvalidActionType
}

// TODO: Using a map to pointers, we can load all the commands into this map,
// then use the name or alias to pull out the pointer for simple lookup.
func (self *CLI) AddCommandToMap(command Command) {
	for _, alias := range command.Names() {
		self.CommandMap[alias] = &command
	}
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
	return (len(self.Flags) > 0)
}

// TODO: We moved flags to a map, like above, we should be opting to use a map
// to pointers of flags, then store all the names and aliases so we can store
// several of the same pointers to different names then just use the map to pull
// things out instead of iterating over each thing and doing bool checks
//func (self *CLI) hasFlag(flag Flag) bool {
//	for _, f := range self.Flags {
//		if flag == f {
//			return true
//		}
//	}
//	return false
//}

func (self *CLI) VisibleCommands() []Command {
	ret := []Command{}
	for _, command := range self.Commands {
		if !command.Hidden {
			ret = append(ret, command)
		}
	}
	return ret
}

func (self *CLI) HasVisibleCommands() bool {
	return (len(self.VisibleCommands()) > 0)
}
