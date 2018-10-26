package cli

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	log "github.com/multiverse-os/cli-framework/log"
	text "github.com/multiverse-os/cli-framework/text"
	color "github.com/multiverse-os/cli-framework/text/color"
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

	Commands    Commands
	Subcommands Commands
	Flags       map[string]Flag

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
	BeforeActions map[string]BeforeFunc
	AfterActions  map[string]AfterFunc
	BeforeAction  BeforeFunc
	AfterAction   AfterFunc
	// Error Functions
	// TODO: Why not just make these locales and print from standard error log?
	CommandNotFound CommandNotFoundFunc
	ExitErrHandler  ExitErrHandlerFunc
	OnUsageError    OnUsageErrorFunc
}

func (self *CLI) PrintBanner() {
	fmt.Println(color.Header(self.Name) + "  " + color.Strong("v"+self.Version.String()))
	fmt.Println(color.Light(text.Repeat("=", 80)))
}

// Setup and New dont seem to have any reason to be separate
func New(cli *CLI) *CLI {
	// TODO: Should just handle compile time and such together with hashing and signatures
	//cmd.CompiledOn = time.Now()
	// Default to same name 'go build' uses for executable: the working directory name
	if cli.Name == "" {
		var err error
		// TODO: We should only be using args AFTER parsing so we can limit by data type, validate and such
		cli.Name, err = filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Print(log.FATAL, "Failed to parse executable working directory in default 'Name' attribute assignment.")
		}
	}
	// Experiment with shorter checking using pointing
	if &cli.Version == nil {
		cli.Version = Version{Major: 0, Minor: 1, Patch: 0}
	}
	//if cmd.Version.Major == 0 && cmd.Version.Minor == 0 && cmd.Version.Patch == 0 {
	//	cmd.Version = Version{
	//		Major: 0,
	//		Minor: 1,
	//		Patch: 0,
	//	}
	//}

	// TODO: Add support for 'nohup' like functionality to output all stdout to text file
	// TODO: How can we merge these two? It should be possible and could work very well
	if cli.Writer == nil {
		cli.Writer = os.Stdout
	}
	if cli.Logger.Name == "" {
		cli.Logger = log.NewSimpleLogger(cli.Name, log.JSON, true)
	}
	cli.Commands = InitCommands()
	if !cli.HideVersion {
		// TODO: We should just have an init function that loads hidden version and
		// help flags. we can use a 'bool' to say if they are visible or not, same
		// with commands above
		//cli.appendFlag(VersionFlag)
	}
	cli.CommandCategories = CommandCategories{}
	for _, command := range cli.Commands {
		cli.AddCommandToMap(command)
	}
	return cli
}

func (self *CLI) Run(arguments []string) (err error) {
	// TODO: Add shell completion code (old code used to be here)

	// TODO: This is tail, why did we bother to make the function if we are not
	// going to use it?
	// TODO: This is a big thing, we should make this its own function.
	// TODO: Parse should probably return to a switch case, then that switch case
	// can resolve the actions that need to be called. instead of assigning. then
	// doing 5 if checks

	// TODO: So here, is where we would see if any action is called, and if
	// defaultaction is nil, then we just call help, this avoids any used
	// memory by just applying the functionality in order of operations

	// TODO: Cant we determine this earlier? like by checking if current parsed command is a command?

	// TODO: arguments come from context, we really need to parse
	// THESE SHOULD BE MOVED OUTSIDE THIS, AND PASSED TO THIS FUCN
	//args = os.Args
	//argCount = len(args)

	// TODO: This should be called from a switch/case that handles the data from a
	// parse commmand parsing the args
	//if self.Action == nil {
	//  self.Action = helpCommand.Action
	//}

	// Run default Action
	err = HandleAction(self.DefaultAction, NewContext(self))
	if err != nil {
		fmt.Println("[Error] " + err.Error())
	}

	return err
}

// TODO: Using a map to pointers, we can load all the commands into this map,
// then use the name or alias to pull out the pointer for simple lookup.
func (self *CLI) AddCommandToMap(command Command) {
	self.CommandMap[command.Name] = &command
	for _, alias := range command.Aliases {
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

// HandleAction attempts to figure out which Action signature was used.  If
// it's an ActionFunc or a func with the legacy signature for Action, the func
// is run!
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
