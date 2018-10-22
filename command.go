package cli

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

// Command is a subcommand for a cli.CLI.
type Command struct {
	// The name of the command
	Name string
	// short name of the command. Typically one character (deprecated, use `Aliases`)
	ShortName string
	// A list of aliases for the command
	Aliases []string
	// A short description of the usage of this command
	Usage string
	// Custom text to show on USAGE section of help
	UsageText string
	// A longer explanation of how the command works
	Description string
	// A short description of the arguments of this command
	ArgsUsage string
	// The category the command is part of
	Category string
	// The function to call when checking for bash command completions
	BashComplete BashCompleteFunc
	// An action to execute before any sub-subcommands are run, but after the context is ready
	// If a non-nil error is returned, no sub-subcommands are run
	Before BeforeFunc
	// An action to execute after any subcommands are run, but after the subcommand has finished
	// It is run even if Action() panics
	After AfterFunc
	// The function to call when this command is invoked
	Action interface{}
	// TODO: replace `Action: interface{}` with `Action: ActionFunc` once some kind
	// of deprecation period has passed, maybe?

	// Execute this function if a usage error occurs.
	OnUsageError OnUsageErrorFunc
	// List of child commands
	Subcommands Commands
	// List of flags to parse
	Flags []Flag
	// Treat all flags as normal arguments if true
	SkipFlagParsing bool
	// Skip argument reordering which attempts to move flags before arguments,
	// but only works if all flags appear after all arguments. This behavior was
	// removed n version 2 since it only works under specific conditions so we
	// backport here by exposing it as an option for compatibility.
	SkipArgReorder bool
	// Boolean to hide built-in help command
	HideHelp bool
	// Boolean to hide this command from help or completion
	Hidden bool
	// Boolean to enable short-option handling so user can combine several
	// single-character bool arguements into one
	// i.e. foobar -o -v -> foobar -ov
	UseShortOptionHandling bool

	// Full name of command for help, defaults to full command name, including parent commands.
	HelpName        string
	commandNamePath []string

	// CustomHelpTemplate the text template for the command help topic.
	// cli.go uses text/template to render templates. You can
	// render custom help text by setting this variable.
	CustomHelpTemplate string
}

type CommandsByName []Command

func (c CommandsByName) Len() int {
	return len(c)
}

func (c CommandsByName) Less(i, j int) bool {
	return lexicographicLess(c[i].Name, c[j].Name)
}

func (c CommandsByName) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

// FullName returns the full name of the command.
// For subcommands this ensures that parent commands are part of the command path
func (c Command) FullName() string {
	if c.commandNamePath == nil {
		return c.Name
	}
	return strings.Join(c.commandNamePath, " ")
}

// Commands is a slice of Command
type Commands []Command

// Run invokes the command given the context, parses ctx.Args() to generate command-specific flags
func (c Command) Run(ctx *Context) (err error) {
	if len(c.Subcommands) > 0 {
		return c.startCLI(ctx)
	}

	if !c.HideHelp && (HelpFlag != BoolFlag{}) {
		// append help to flags
		c.Flags = append(
			c.Flags,
			HelpFlag,
		)
	}

	set, err := c.parseFlags(ctx.Args().Tail())

	context := NewContext(ctx.CLI, set, ctx)
	context.Command = c
	if checkCommandCompletions(context, c.Name) {
		return nil
	}

	if err != nil {
		if c.OnUsageError != nil {
			err := c.OnUsageError(context, err, false)
			context.CLI.handleExitCoder(context, err)
			return err
		}
		fmt.Fprintln(context.CLI.Writer, "Incorrect Usage:", err.Error())
		fmt.Fprintln(context.CLI.Writer)
		ShowCommandHelp(context, c.Name)
		return err
	}

	if checkCommandHelp(context, c.Name) {
		return nil
	}

	if c.After != nil {
		defer func() {
			afterErr := c.After(context)
			if afterErr != nil {
				context.CLI.handleExitCoder(context, err)
				if err != nil {
					err = NewMultiError(err, afterErr)
				} else {
					err = afterErr
				}
			}
		}()
	}

	if c.Before != nil {
		err = c.Before(context)
		if err != nil {
			ShowCommandHelp(context, c.Name)
			context.CLI.handleExitCoder(context, err)
			return err
		}
	}

	if c.Action == nil {
		c.Action = helpSubcommand.Action
	}

	err = HandleAction(c.Action, context)

	if err != nil {
		context.CLI.handleExitCoder(context, err)
	}
	return err
}

func (c *Command) parseFlags(args Args) (*flag.FlagSet, error) {
	set, err := flagSet(c.Name, c.Flags)
	if err != nil {
		return nil, err
	}
	set.SetOutput(ioutil.Discard)

	if c.SkipFlagParsing {
		return set, set.Parse(append([]string{"--"}, args...))
	}

	if c.UseShortOptionHandling {
		args = translateShortOptions(args)
	}

	if !c.SkipArgReorder {
		args = reorderArgs(args)
	}

	err = set.Parse(args)
	if err != nil {
		return nil, err
	}

	err = normalizeFlags(c.Flags, set)
	if err != nil {
		return nil, err
	}

	return set, nil
}

// reorderArgs moves all flags before arguments as this is what flag expects
func reorderArgs(args []string) []string {
	var nonflags, flags []string

	readFlagValue := false
	for i, arg := range args {
		if arg == "--" {
			nonflags = append(nonflags, args[i:]...)
			break
		}

		if readFlagValue && !strings.HasPrefix(arg, "-") && !strings.HasPrefix(arg, "--") {
			readFlagValue = false
			flags = append(flags, arg)
			continue
		}
		readFlagValue = false

		if arg != "-" && strings.HasPrefix(arg, "-") {
			flags = append(flags, arg)

			readFlagValue = !strings.Contains(arg, "=")
		} else {
			nonflags = append(nonflags, arg)
		}
	}

	return append(flags, nonflags...)
}

func translateShortOptions(flagArgs Args) []string {
	// separate combined flags
	var flagArgsSeparated []string
	for _, flagArg := range flagArgs {
		if strings.HasPrefix(flagArg, "-") && strings.HasPrefix(flagArg, "--") == false && len(flagArg) > 2 {
			for _, flagChar := range flagArg[1:] {
				flagArgsSeparated = append(flagArgsSeparated, "-"+string(flagChar))
			}
		} else {
			flagArgsSeparated = append(flagArgsSeparated, flagArg)
		}
	}
	return flagArgsSeparated
}

// Names returns the names including short names and aliases.
func (c Command) Names() []string {
	names := []string{c.Name}

	if c.ShortName != "" {
		names = append(names, c.ShortName)
	}

	return append(names, c.Aliases...)
}

// HasName returns true if Command.Name or Command.ShortName matches given name
func (c Command) HasName(name string) bool {
	for _, n := range c.Names() {
		if n == name {
			return true
		}
	}
	return false
}

// TODO: Why are we recreating the entire object? Can we not just save copy the object entirely
// instead of recreating it attribute by attribute?
func (c Command) startCLI(ctx *Context) error {
	app := New(&CLI{
		Metadata:              ctx.CLI.Metadata,
		Name:                  ctx.CLI.Name,
		Usage:                 c.Usage,
		Description:           c.Description,
		ArgsUsage:             c.ArgsUsage,
		CommandNotFound:       ctx.CLI.CommandNotFound,
		CustomCLIHelpTemplate: c.CustomHelpTemplate,
		Commands:              c.Subcommands,
		Flags:                 c.Flags,
		HideHelp:              c.HideHelp,
		Version:               ctx.CLI.Version,
		HideVersion:           ctx.CLI.HideVersion,
		CompiledOn:            ctx.CLI.CompiledOn,
		Writer:                ctx.CLI.Writer,
		ErrWriter:             ctx.CLI.ErrWriter,
		categories:            CommandCategories{},
		BashCompletion:        ctx.CLI.BashCompletion,
		OnUsageError:          c.OnUsageError,
		Before:                c.Before,
		After:                 c.After,
		Logger:                ctx.CLI.Logger,
	})
	for _, command := range c.Subcommands {
		app.categories = app.categories.AddCommand(command.Category, command)
	}

	sort.Sort(app.categories)

	if c.BashComplete != nil {
		app.BashComplete = c.BashComplete
	}

	if c.Action != nil {
		app.Action = c.Action
	} else {
		app.Action = helpSubcommand.Action
	}

	for index, cc := range app.Commands {
		app.Commands[index].commandNamePath = []string{c.Name, cc.Name}
	}

	return app.RunAsSubcommand(ctx)
}

// VisibleFlags returns a slice of the Flags with Hidden=false
func (c Command) VisibleFlags() []Flag {
	return visibleFlags(c.Flags)
}
