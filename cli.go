package cli

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"time"

	log "github.com/multiverse-os/cli-framework/log"
	text "github.com/multiverse-os/cli-framework/text"
	"github.com/multiverse-os/cli-framework/text/color"
)

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
	Commands          []Command
	Subcommands       map[string]Command
	Flags             map[string]Flag
	Logger            log.Logger
	CompiledOn        time.Time
	HideHelp          bool
	HideVersion       bool
	CommandCategories CommandCategories
	Writer            io.Writer
	ErrWriter         io.Writer
	RenderHelpText    PrintHelpText
	// Functions
	//////////////////////////////////////////////////////////////////////////////
	DefaultAction interface{}
	BeforeAction  BeforeFunc
	AfterAction   AfterFunc
	// Autocomplete Function - no need to have this have a bool check, obviously you want intuitive tab autocompletion by default in no situation is intiutive and consistent design undesirable
	BashComplete BashCompleteFunc
	// Error Functions
	CommandNotFound CommandNotFoundFunc
	ExitErrHandler  ExitErrHandlerFunc
	OnUsageError    OnUsageErrorFunc
}

// Setup and New dont seem to have any reason to be separate
func New(cmd *CLI) *CLI {
	// TODO: Should just handle compile time and such together with hashing and signatures
	//cmd.CompiledOn = time.Now()
	// Default to same name 'go build' uses for executable: the working directory name
	if cmd.Name == "" {
		var err error
		cmd.Name, err = filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Print(log.FATAL, "Failed to parse executable working directory in default 'Name' attribute assignment.")
		}
	}
	// Experiment with shorter checking using pointing
	if &cmd.Version == nil {
		cmd.Version = Version{Major: 0, Minor: 1, Patch: 0}
	}
	//if cmd.Version.Major == 0 && cmd.Version.Minor == 0 && cmd.Version.Patch == 0 {
	//	cmd.Version = Version{
	//		Major: 0,
	//		Minor: 1,
	//		Patch: 0,
	//	}
	//}
	// If Action is not specified, the default action should be to render help text
	// TODO: Is this assignment necessary? Can we just check if the default action
	// is empty at runtime? and then just run the help command then instead of wasting this memory space?
	if cmd.DefaultAction == nil {
		cmd.DefaultAction = helpCommand.Action
	}
	// TODO: Where is the reader? Why do we have this attribute? Do we just need the writer?
	if cmd.Writer == nil {
		cmd.Writer = os.Stdout
	}
	if cmd.Logger.AppName == "" {
		cmd.Logger = log.NewLogger(cmd.Name, 1, true, true, false)
	}

	if self.Command(helpCommand.Name) == nil && !self.HideHelp {
		self.Commands = append(self.Commands, helpCommand)
		if (HelpFlag != BoolFlag{}) {
			self.appendFlag(HelpFlag)
		}
	}

	if !self.HideVersion {
		self.appendFlag(VersionFlag)
	}

	self.categories = CommandCategories{}
	for _, command := range self.Commands {
		self.categories = self.categories.AddCommand(command.Category, command)
	}
	sort.Sort(self.categories)

	if self.Metadata == nil {
		self.Metadata = make(map[string]interface{})
	}

	if self.Writer == nil {
		self.Writer = os.Stdout
	}
}

func (self *CLI) Run(arguments []string) (err error) {
	// TODO: Shell completion
	// handle the completion flag separately from the flagset since
	// completion could be attempted after a flag, but before its value was put
	// on the command line. this causes the flagset to interpret the completion
	// flag name as the value of the flag before it which is undesirable
	// note that we can only do this because the shell autocomplete function
	// always appends the completion flag at the end of the command
	//shellComplete, arguments := checkShellCompleteFlag(self, arguments)
	//context.shellComplete = shellComplete

	// TODO: Uhh, what is this?
	set.SetOutput(ioutil.Discard)

	// TODO: This is tail, why did we bother to make the function if we are not
	// going to use it?
	// TODO: This is a big thing, we should make this its own function.
	err = set.Parse(arguments[1:])

	// TODO: So here, is where we would see if any action is called, and if
	// defaultaction is nil, then we just call help, this avoids any used
	// memory by just applying the functionality in order of operations

	context := NewContext(self, set, nil)
	// TODO: What is normalize flags? we junkered most of the blaoted flag code in favor of a single reflect func
	normErr := normalizeFlags(self.Flags, set)
	if normErr != nil {
		fmt.Fprintln(self.Writer, normErr)
		ShowCLIHelp(context)
		return nerr
	}
	if err != nil {
		// TODO: Cant we determine this earlier? like by checking if current parsed command is a command?
		if self.OnUsageError != nil {
			err := self.OnUsageError(context, err, false)
			self.handleExitCoder(context, err)
			return err
		}
		fmt.Fprintf(self.Writer, "%s %s\n\n", "Incorrect Usage.", err.Error())
		ShowCLIHelp(context)
		return err
	}

	if self.After != nil {
		defer func() {
			if afterErr := self.After(context); afterErr != nil {
				if err != nil {
					err = NewMultiError(err, afterErr)
				} else {
					err = afterErr
				}
			}
		}()
	}

	if self.Before != nil {
		beforeErr := self.Before(context)
		if beforeErr != nil {
			fmt.Fprintf(self.Writer, "%v\n\n", beforeErr)
			ShowCLIHelp(context)
			self.handleExitCoder(context, beforeErr)
			err = beforeErr
			return err
		}
	}

	args := context.Args()
	if args.Present() {
		name := args.First()
		c := self.Command(name)
		if c != nil {
			return c.Run(context)
		}
	}

	if self.Action == nil {
		self.Action = helpCommand.Action
	}

	// Run default Action
	err = HandleAction(self.Action, context)

	self.handleExitCoder(context, err)
	return err
}

// RunAndExitOnError calls .Run() and exits non-zero if an error was returned
//
// Deprecated: instead you should return an error that fulfills cli.ExitCoder
// to cli.CLI.Run. This will cause the application to exit with the given eror
// code in the cli.ExitCoder
func (self *CLI) RunAndExitOnError() {
	if err := self.Run(os.Args); err != nil {
		fmt.Fprintln(self.errWriter(), err)
		OsExiter(1)
	}
}

// RunAsSubcommand invokes the subcommand given the context, parses ctx.Args() to
// generate command-specific flags
func (self *CLI) RunAsSubcommand(ctx *Context) (err error) {
	// append help to commands
	if len(self.Commands) > 0 {
		if self.Command(helpCommand.Name) == nil && !self.HideHelp {
			self.Commands = append(self.Commands, helpCommand)
			if (HelpFlag != BoolFlag{}) {
				self.appendFlag(HelpFlag)
			}
		}
	}

	newCmds := []Command{}
	for _, c := range self.Commands {
		newCmds = append(newCmds, c)
	}
	self.Commands = newCmds

	// parse flags
	set, err := flagSet(self.Name, self.Flags)
	if err != nil {
		return err
	}

	set.SetOutput(ioutil.Discard)
	err = set.Parse(ctx.Args().Tail())
	nerr := normalizeFlags(self.Flags, set)
	context := NewContext(self, set, ctx)

	if nerr != nil {
		fmt.Fprintln(self.Writer, nerr)
		fmt.Fprintln(self.Writer)
		if len(self.Commands) > 0 {
			ShowSubcommandHelp(context)
		} else {
			ShowCommandHelp(ctx, context.Args().First())
		}
		return nerr
	}

	if checkCompletions(context) {
		return nil
	}

	if err != nil {
		if self.OnUsageError != nil {
			err = self.OnUsageError(context, err, true)
			self.handleExitCoder(context, err)
			return err
		}
		fmt.Fprintf(self.Writer, "%s %s\n\n", "Incorrect Usage.", err.Error())
		ShowSubcommandHelp(context)
		return err
	}

	if len(self.Commands) > 0 {
		if checkSubcommandHelp(context) {
			return nil
		}
	} else {
		if checkCommandHelp(ctx, context.Args().First()) {
			return nil
		}
	}

	if self.After != nil {
		defer func() {
			afterErr := self.After(context)
			if afterErr != nil {
				self.handleExitCoder(context, err)
				if err != nil {
					err = NewMultiError(err, afterErr)
				} else {
					err = afterErr
				}
			}
		}()
	}

	if self.Before != nil {
		beforeErr := self.Before(context)
		if beforeErr != nil {
			self.handleExitCoder(context, beforeErr)
			err = beforeErr
			return err
		}
	}

	args := context.Args()
	if args.Present() {
		name := args.First()
		c := self.Command(name)
		if c != nil {
			return c.Run(context)
		}
	}

	// Run default Action
	err = HandleAction(self.Action, context)
	self.handleExitCoder(context, err)
	return err
}

// Command returns the named command on CLI. Returns nil if the command does not exist
func (self *CLI) Command(name string) *Command {
	for _, c := range self.Commands {
		if c.HasName(name) {
			return &c
		}
	}

	return nil
}

func (self *CLI) Categories() CommandCategories {
	return self.categories
}

func (self *CLI) VisibleFlags() []Flag {
	return visibleFlags(self.Flags)
}

func (self *CLI) HasVisibleFlags() bool {
	return (len(self.Flags) > 0)
}

func (self *CLI) hasFlag(flag Flag) bool {
	for _, f := range self.Flags {
		if flag == f {
			return true
		}
	}
	return false
}

func (self *CLI) VisibleCategories() []*CommandCategory {
	ret := []*CommandCategory{}
	for _, category := range self.categories {
		if visible := func() *CommandCategory {
			for _, command := range category.Commands {
				if !command.Hidden {
					return category
				}
			}
			return nil
		}(); visible != nil {
			ret = append(ret, visible)
		}
	}
	return ret
}

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

func (self *CLI) errWriter() io.Writer {
	// When the app ErrWriter is nil use the package level one.
	if self.ErrWriter == nil {
		return ErrWriter
	}

	return self.ErrWriter
}

func (self *CLI) appendFlag(flag Flag) {
	if !self.hasFlag(flag) {
		self.Flags = append(self.Flags, flag)
	}
}

func (self *CLI) handleExitCoder(context *Context, err error) {
	if self.ExitErrHandler != nil {
		self.ExitErrHandler(context, err)
	} else {
		HandleExitCoder(err)
	}
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

func (self *CLI) LogFile() string {
	return (self.Logger.Path + self.Logger.Filename)
}

func (self *CLI) PrintBanner() {
	fmt.Println(color.Header(self.Name) + "  " + color.Strong("v"+self.Version.String()))
	fmt.Println(color.Light(text.Repeat("=", 80)))
}
