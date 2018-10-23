package cli

import (
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"

	log "github.com/multiverse-os/cli-framework/log"
)

// TODO: replace `Action: interface{}` with `Action: ActionFunc` once some kind
// of deprecation period has passed, maybe? // Execute this function if a usage error occurs.
// TODO: Why do we have 'Usage' AND 'UsageText' seems like we should be merging this in some way. Also is this diff than description?
// TODO: Short option? Its alias, flags only require special short option becausxe they parse with 1 '-' instead of '--'
type Command struct {
	Name            string
	Aliases         []string
	Category        string
	CommandCategory *CommandCategory
	Usage           string
	UsageText       string
	Description     string
	ArgsUsage       string
	ParentCommand   Command
	Subcommands     map[string]Command
	Flags           map[string]Flag
	//SkipFlagParsing bool
	SkipArgReorder  bool
	Hidden          bool
	commandNamePath []string
	CustomHelpText  string

	Action interface{}
	Before BeforeFunc
	After  AfterFunc

	BashComplete BashCompleteFunc

	OnUsageError OnUsageErrorFunc
}

type Commands []Command
type CommandsByName []Command

// TODO: Why not instead of using globals we make a function that initializes the
// CLI struct commands array with help command

// TODO: Why do we have shortOption if we are just going to use alias
func InitCommands() (commands Commands) {
	return append(commands, Command{
		Name:      "help",
		Aliases:   []string{"h"},
		Usage:     "List of available commands or details for a specified command",
		ArgsUsage: "[command]",
		Hidden:    true,
		Action: func(c *Context) error {
			args := c.Args()
			if args.Present() {
				return ShowCommandHelp(c, args.First())
			}
			ShowCLIHelp(c)
			return nil
		},
	})
}

func InitSubcommands(subcommands Commands) {
	return append(subcommands, Command{
		Name:      "help",
		Aliases:   []string{"h"},
		Usage:     "List of available commands or details for a specified command",
		ArgsUsage: "[command]",
		Action: func(c *Context) error {
			args := c.Args()
			if args.Present() {
				return ShowCommandHelp(c, args.First())
			}
			return ShowSubcommandHelp(c)
		},
	})
}

// TODO: Currently only works with 3 levels, but program could support infinite levels. Should
// rebuild this function to support recursive level of nested commands and subcommands
func (self Command) Breadcrumbs() Commands {
	if self.ParentCommand == nil {
		return []Command{self}
	} else {
		if self.ParentCommand.ParentCommand == nil {
			return []Command{self.ParentCommand, self}
		} else {
			return []Command{self.ParentCommand.ParentCommand, self.ParentCommand, self}
		}
	}
}

func (self Command) HasSubcommands() bool {
	return (len(self.Subcommands) > 0)
}

func (self Command) Run(ctx *Context) (err error) {
	// TODO: So what is not explained here, is we are nesting the CLI function essentially, and running a new instance
	// every time we use a command or subcommand. This is an interesting design, but could be implemented better than
	// it is and maybe explain a bit for other developers to make sense of it quicker
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
		// TODO: What is going in here?
		fmt.Fprintln(context.CLI.Writer, "Incorrect Usage:", err.Error())
		fmt.Fprintln(context.CLI.Writer)
		ShowCommandHelp(context, self.Name)
		return err
	}

	// TODO: Break After, Before and Action into their own functions to clean up this parent function
	// and make it very clear how the hooks are working.

	// TODO: Interesting that After is ran before Before and deferred, seems clever but not sure
	// if it actually provides expected functionality
	if self.After != nil {
		defer func() {
			afterErr := self.After(context)
			if afterErr != nil {
				log.Fatal(afterErr)
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
	// TODO: This is right, so why are tehre like 5 functions and tons of extra memory dedicated to knowing
	// if the help command or help subcommand is being displayed. if tehre is no default action, we render help
	// if we parse help, we render help.
	// TODO: no need to handle action, just print help and exit
	if self.Action == nil {
		self.Action = helpSubcommand.Action
		err = HandleAction(self.Action, context)
		if err != nil {
			context.CLI.handleExitCoder(context, err)
		}
	}
	return err
}

func (c *Command) parseFlags(args Args) (*flag.FlagSet, error) {
	set, err := flagSet(c.Name, c.Flags)
	if err != nil {
		return nil, err
	}
	set.SetOutput(ioutil.Discard)
	//if c.SkipFlagParsing {
	//	return set, set.Parse(append([]string{"--"}, args...))
	//}
	// TODO: Think we just use aliases
	//if c.UseShortOptionHandling {
	//  args = translateShortOptions(args)
	//}
	// TODO: Parse and handle result in a switchase
	err = set.Parse(args)
	if err != nil {
		return nil, err
	}
	//err = normalizeFlags(c.Flags, set)
	//if err != nil {
	//	return nil, err
	//}
	return set, nil
}

// TODO: Removing reorder flags to before args func because why? if we are aprsing
// that there are flags in any location we can then load the data into the command
// or CLI struct for execution, but we are not reprinting the command with better
// form, so why reorder? thats a lot of wasted resources for nothing

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

func (c Command) Names() []string {
	names := []string{c.Name}
	if c.ShortName != "" {
		names = append(names, c.ShortName)
	}
	return append(names, c.Aliases...)
}

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
