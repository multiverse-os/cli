package cli

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

// TODO: replace `Action: interface{}` with `Action: ActionFunc` once some kind
// of deprecation period has passed, maybe?
// Execute this function if a usage error occurs.
// TODO: Why do we have 'Usage' AND 'UsageText' seems like we should be merging this in some way. Also is this diff than description?
type Command struct {
	Name            string
	ShortOption     string
	Aliases         []string
	Usage           string
	UsageText       string
	Description     string
	ArgsUsage       string
	Category        string
	CommandCategory *CommandCategory
	Action          interface{}
	Subcommands     Commands
	Flags           []Flag
	SkipFlagParsing bool
	SkipArgReorder  bool
	// Why is there hide help for a single command? when would help displayed for each command ever be desirable?
	HideHelp           bool
	Hidden             bool
	commandNamePath    []string
	CustomHelpTemplate string
	BashComplete       BashCompleteFunc
	Before             BeforeFunc
	After              AfterFunc
	OnUsageError       OnUsageErrorFunc
}

type Commands []Command
type CommandsByName []Command

// TODO: What the fuck is this, the name is terrible
func (self Command) FullName() string {
	if self.commandNamePath == nil {
		return self.Name
	}
	return strings.Join(self.commandNamePath, " ")
}

func (self Command) HasSubcommands() bool {
	return (len(self.Subcommands) > 0)
}

func (self Command) Run(ctx *Context) (err error) {
	if self.HasSubcommands {
		return self.startCLI(ctx)
	}
	if !self.HideHelp {
		self.Flags = append(self.Flags, HelpFlag)
	}
	set, err := self.parseFlags(ctx.Args().Tail())
	context := NewContext(ctx.CLI, set, ctx)
	context.Command = self
	if checkCommandCompletions(context, self.Name) {
		return nil
	}
	if err != nil {
		if self.OnUsageError != nil {
			err := self.OnUsageError(context, err, false)
			context.CLI.handleExitCoder(context, err)
			return err
		}
		fmt.Fprintln(context.CLI.Writer, "Incorrect Usage:", err.Error())
		fmt.Fprintln(context.CLI.Writer)
		ShowCommandHelp(context, self.Name)
		return err
	}
	if checkCommandHelp(context, self.Name) {
		return nil
	}
	if self.After != nil {
		defer func() {
			afterErr := self.After(context)
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
	if self.Before != nil {
		err = self.Before(context)
		if err != nil {
			ShowCommandHelp(context, self.Name)
			context.CLI.handleExitCoder(context, err)
			return err
		}
	}
	if self.Action == nil {
		self.Action = helpSubcommand.Action
	}
	err = HandleAction(self.Action, context)
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
	cmd := &CLI{
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
	}
	for _, command := range c.Subcommands {
		cmd.categories = cmd.categories.AddCommand(command.Category, command)
	}

	sort.Sort(cmd.categories)

	if c.BashComplete != nil {
		cmd.BashComplete = c.BashComplete
	}

	if c.Action != nil {
		cmd.Action = c.Action
	} else {
		cmd.Action = helpSubcommand.Action
	}

	for index, cc := range cmd.Commands {
		cmd.Commands[index].commandNamePath = []string{c.Name, cc.Name}
	}

	return cmd.RunAsSubcommand(ctx)
}

func (c Command) VisibleFlags() (flags []Flag) {
	for _, flag := range c.Flags {
		if !flag.Hidden {
			flags = append(flags, flag)
		}
	}
	return flags
}
