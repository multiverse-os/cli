package cli

import (
	"flag"
	"io/ioutil"

	log "github.com/multiverse-os/cli-framework/log"
)

// TODO: Why do we have 'Usage' AND 'UsageText' seems like we should be merging this in some way. Also is this diff than description?
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

func InitCommands(command Command) (commands Commands) {
	return append(commands, Command{
		Name:          "help",
		Aliases:       []string{"h"},
		Usage:         "List of available commands or details for a specified command",
		ArgsUsage:     "[command]",
		ParentCommand: command,
		Subcommands:   InitSubcommands(),
		Hidden:        true,
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

func InitSubcommands() (subcommands Commands) {
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

func (c Command) VisibleFlags() (flags []Flag) {
	for _, flag := range c.Flags {
		if !flag.Hidden {
			flags = append(flags, flag)
		}
	}
	return flags
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
	// TODO: Structure needs to evolve into parse -> execute using switch case
	set, err := self.parseFlags(ctx.Args().Tail())

	//
	// TODO: Why are we creating a new context? why not just use existing one?
	////////////////////////////////////////////////////////////////////////////
	//context := NewContext(ctx.CLI, set, ctx)
	//// TODO: What? Why would we bother doing that?
	//context.Command = self
	//if err != nil {
	//	if self.OnUsageError != nil {
	//		err := self.OnUsageError(context, err, false)
	//		context.CLI.handleExitCoder(context, err)
	//		return err
	//	}
	//	// TODO: What is going in here?
	//	fmt.Fprintln(context.CLI.Writer, "Incorrect Usage:", err.Error())
	//	fmt.Fprintln(context.CLI.Writer)
	//	ShowCommandHelp(context, self.Name)
	//	return err
	//}
	// TODO: Break After, Before and Action into their own functions to clean up this parent function
	// and make it very clear how the hooks are working.
	// TODO: Interesting that After is ran before Before and deferred, seems clever but not sure
	// if it actually provides expected functionality
	// TODO: This is right, so why are tehre like 5 functions and tons of extra memory dedicated to knowing
	// if the help command or help subcommand is being displayed. if tehre is no default action, we render help
	// if we parse help, we render help.
	// TODO: no need to handle action, just print help and exit
	defer self.ExecuteAfterAction()
	self.ExecuteBeforeAction()
	self.ExecuteAction()
	return err
}

func (self *Command) ExecuteBeforeAction() {
	if self.BeforeAction != nil {
		err = self.BeforeAction(context)
		if err != nil {
			ShowCommandHelp(context, self.Name)
			context.CLI.handleExitCoder(context, err)
			return err
		}
	}
}

func (self *Command) ExecuteAfterAction() {
	if self.After != nil {
		afterErr := self.After(context)
		if afterErr != nil {
			log.Fatal(afterErr)
			if err != nil {
				err = NewMultiError(err, afterErr)
			} else {
				err = afterErr
			}
		}
	}
}

func (self *Command) ExecuteAction() {
	if self.Action == nil {
		self.Action = helpSubcommand.Action
		err = HandleAction(self.Action, context)
		if err != nil {
			context.CLI.handleExitCoder(context, err)
		}
	}
}

func (c *Command) parseFlags(args Args) (*flag.FlagSet, error) {
	// TODO ?
	set.SetOutput(ioutil.Discard)
	// TODO: We dont skip flag parsing, we can just skip executing
	// TODO: Parse and handle result in a switchase
	err = set.Parse(args)
	if err != nil {
		return nil, err
	}
	return set, nil
}

func (self Command) Names() []string {
	return append([]string{self.Name}, self.Aliases...)
}

func (self Command) HasName(name string) bool {
	if len(self.Aliases) == 0 {
		return (self.Name == name)
	} else {
		for _, commandName := range self.Names() {
			if commandName == name {
				return true
			}
		}
		return false
	}
}
