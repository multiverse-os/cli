package cli

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"
)

var (
	errInvalidActionType = NewExitError("ERROR invalid Action type. Must be `func(*Context`)` or `func(*Context) error).", 2)
)

// TODO: Move all text into locales so we can support localization

// TODO: Bash complete should be default, and it should just be setup automatically by just iterating over the available
// command names, flag names and such. No reason we should have to re-state information which is already stored to acheive
// this functionality.
type CLI struct {
	// The name of the program. Defaults to path.Base(os.Args[0])
	Name string
	// Full name of command for help, defaults to Name
	HelpName string
	// Color Help ouput
	ColorOutput bool
	// Description of the program.
	Usage string
	// Text to override the USAGE section of help
	UsageText string
	// Description of the program argument format.
	ArgsUsage string
	// Version of the program
	Version
	// Description of the program
	Description string
	// List of commands to execute
	Commands []Command
	// List of flags to parse
	Flags []Flag
	// Simple logging system init - optional, this just initializes it with the same app name
	Logger log.Logger
	// Boolean to enable bash completion commands
	BashCompletion bool
	// Boolean to hide built-in help command
	HideHelp bool
	// Boolean to hide built-in version flag and the VERSION section of help
	HideVersion bool
	// Populate on app startup, only gettable through method Categories()
	categories CommandCategories
	// An action to execute when the bash-completion flag is set
	BashComplete BashCompleteFunc
	// An action to execute before any subcommands are run, but after the context is ready
	// If a non-nil error is returned, no subcommands are run
	Before BeforeFunc
	// An action to execute after any subcommands are run, but after the subcommand has finished
	// It is run even if Action() panics
	After AfterFunc

	// The action to execute when no subcommands are specified
	// Expects a `cli.ActionFunc` but will accept the *deprecated* signature of `func(*cli.Context) {}`
	// *Note*: support for the deprecated `Action` signature will be removed in a future version
	Action interface{}

	// Execute this function if the proper command cannot be found
	CommandNotFound CommandNotFoundFunc
	// Execute this function if an usage error occurs
	OnUsageError OnUsageErrorFunc
	// Compilation date
	CompiledOn time.Time
	// Writer writer to write output to
	Writer io.Writer
	// ErrWriter writes error output
	ErrWriter io.Writer
	// Execute this function to handle ExitErrors. If not provided, HandleExitCoder is provided to
	// function as a default, so this is optional.
	ExitErrHandler ExitErrHandlerFunc
	// Other custom info
	Metadata map[string]interface{}
	// Carries a function which returns app specific info.
	ExtraInfo func() map[string]string
	// CustomCLIHelpTemplate the text template for app help topic.
	// cli.go uses text/template to render templates. You can
	// render custom help text by setting this variable.
	CustomCLIHelpTemplate string

	didSetup bool
}

func compiledOn() time.Time {
	info, err := os.Stat(os.Args[0])
	if err != nil {
		return time.Now()
	}
	return info.ModTime()
}

// New creates a new cli CLIlication with some reasonable defaults for Name,
// Usage, Version and Action.
func New(cmd *CLI) *CLI {
	if cmd.Name == "" {
		cmd.Name = filepath.Base(os.Args[0])
	}
	if cmd.HelpName == "" {
		cmd.HelpName = filepath.Base(os.Args[0])
	}
	if cmd.Usage == "" {
		cmd.Usage = "A new command-line interface"
	}
	if cmd.Version.Major == 0 && cmd.Version.Minor == 0 && cmd.Version.Patch == 0 {
		cmd.Version = Version{
			Major: 0,
			Minor: 1,
			Patch: 0,
		}
	}
	if cmd.BashComplete == nil {
		cmd.BashComplete = DefaultCLIComplete
	}
	if cmd.Action == nil {
		cmd.Action = helpCommand.Action
	}
	if cmd.Writer == nil {
		cmd.Writer = os.Stdout
	}
	cmd.CompiledOn = compiledOn()
	return cmd
}

// Setup runs initialization code to ensure all data structures are ready for
// `Run` or inspection prior to `Run`.  It is internally called by `Run`, but
// will return early if setup has already happened.
func (self *CLI) Setup() {
	if self.didSetup {
		return
	}

	self.didSetup = true

	newCmds := []Command{}
	for _, c := range self.Commands {
		if c.HelpName == "" {
			c.HelpName = fmt.Sprintf("%s %s", self.HelpName, c.Name)
		}
		newCmds = append(newCmds, c)
	}
	self.Commands = newCmds

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

// Run is the entry point to the cli app. Parses the arguments slice and routes
// to the proper flag/args combination
func (self *CLI) Run(arguments []string) (err error) {
	self.Setup()

	// handle the completion flag separately from the flagset since
	// completion could be attempted after a flag, but before its value was put
	// on the command line. this causes the flagset to interpret the completion
	// flag name as the value of the flag before it which is undesirable
	// note that we can only do this because the shell autocomplete function
	// always appends the completion flag at the end of the command
	shellComplete, arguments := checkShellCompleteFlag(self, arguments)

	// parse flags
	set, err := flagSet(self.Name, self.Flags)
	if err != nil {
		return err
	}

	set.SetOutput(ioutil.Discard)
	err = set.Parse(arguments[1:])
	nerr := normalizeFlags(self.Flags, set)
	context := NewContext(self, set, nil)
	if nerr != nil {
		fmt.Fprintln(self.Writer, nerr)
		ShowCLIHelp(context)
		return nerr
	}
	context.shellComplete = shellComplete

	if checkCompletions(context) {
		return nil
	}

	if err != nil {
		if self.OnUsageError != nil {
			err := self.OnUsageError(context, err, false)
			self.handleExitCoder(context, err)
			return err
		}
		fmt.Fprintf(self.Writer, "%s %s\n\n", "Incorrect Usage.", err.Error())
		ShowCLIHelp(context)
		return err
	}

	if !self.HideHelp && checkHelp(context) {
		ShowCLIHelp(context)
		return nil
	}

	if !self.HideVersion && checkVersion(context) {
		PrintVersion(context)
		return nil
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
		if c.HelpName == "" {
			c.HelpName = fmt.Sprintf("%s %s", self.HelpName, c.Name)
		}
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

// Categories returns a slice containing all the categories with the commands they contain
func (self *CLI) Categories() CommandCategories {
	return self.categories
}

// VisibleCategories returns a slice of categories and commands that are
// Hidden=false
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

// VisibleCommands returns a slice of the Commands with Hidden=false
func (self *CLI) VisibleCommands() []Command {
	ret := []Command{}
	for _, command := range self.Commands {
		if !command.Hidden {
			ret = append(ret, command)
		}
	}
	return ret
}

// VisibleFlags returns a slice of the Flags with Hidden=false
func (self *CLI) VisibleFlags() []Flag {
	return visibleFlags(self.Flags)
}

func (self *CLI) hasFlag(flag Flag) bool {
	for _, f := range self.Flags {
		if flag == f {
			return true
		}
	}

	return false
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
