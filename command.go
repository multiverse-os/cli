package cli

// TODO: Why do we have 'Usage' AND 'UsageText' seems like we should be merging this in some way. Also is this diff than description?
type Command struct {
	Name    string
	Aliases []string

	Category        string
	CommandCategory *CommandCategory
	ParentCommand   *Command
	Subcommands     map[string]Command
	Flags           map[string]Flag

	// A short description of the usage of this command
	Usage string
	// Custom text to show on USAGE section of help
	//UsageText     string
	// A longer explanation of how the command works
	//ArgsUsage     string
	// A short description of the arguments of this command
	//Description   string

	//SkipFlagParsing bool
	SkipArgReorder  bool
	Hidden          bool
	commandNamePath []string
	CustomHelpText  string

	Action       interface{}
	Before       func()
	After        func()
	BashComplete func()
	OnUsageError func()
}

type Commands []Command

func (self Command) Names() []string {
	return append([]string{self.Name}, self.Aliases...)
}

func defaultCommands() []Command {
	// TODO: This inits a slice of commands, moving towards either radix tree or
	// just map
	return []Command{
		Command{
			Hidden:  true,
			Name:    "help",
			Aliases: []string{"h"},
			Usage:   "List of available commands or details for a specified command",
			//ArgsUsage: "[command]",
			//Subcommands: InitSubcommands(),
			Action: func() error {
				// TODO: Args need to be loaded into context so its accessible
				//args := c.Args()
				//if args.Present() {
				//	return ShowCommandHelp(c, args.First())
				//}
				//ShowCLIHelp(c)
				return nil
			},
		},
	}
}

func (self Command) InitSubcommands() []Command {
	return []Command{
		Command{
			Name:    "help",
			Aliases: []string{"h"},
			Usage:   "List of available commands or details for a specified command",
			//ArgsUsage:     "[command]",
			ParentCommand: &self,
			Action: func() error {
				// TODO: Fix this because this is all leading to massive bloat
				//args := c.Args()
				//if args.Present() {
				//	return ShowCommandHelp(c, args.First())
				//}
				//return ShowSubcommandHelp(c)
				return nil
			},
		},
	}
}

func (c Command) VisibleFlags() (flags []Flag) {
	for _, flag := range c.Flags {
		if !flag.Hidden {
			flags = append(flags, flag)
		}
	}
	return flags
}

func (self Command) HasSubcommands() bool {
	return (len(self.Subcommands) > 0)
}

func (self Command) Run() (err error) {
	// TODO: So what is not explained here, is we are nesting the CLI function essentially, and running a new instance
	// every time we use a command or subcommand. This is an interesting design, but could be implemented better than
	// it is and maybe explain a bit for other developers to make sense of it quicker

	// TODO: This nested context is really just leading to bloat and lots of extra
	// memory usage
	//if self.HasSubcommands() {
	//return self.startCLI(ctx)
	//}

	// TODO: Structure needs to evolve into parse -> execute using switch case
	// TODO: Fix parseflags, it should be in general parsing
	//set, err := self.parseFlags(ctx.Args().Tail())

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
	//defer self.ExecuteAfterAction()
	//self.ExecuteBeforeAction()
	//self.ExecuteAction()
	return err
}

// TODO: This should be far more generic, have a generic hook system we can just
// push actions into instead of having special before/after actions for
// commands, subcommands, flags, regular actions, etc
//func (self *Command) ExecuteBeforeAction() {
//	if self.BeforeAction != nil {
//		err = self.BeforeAction(context)
//		if err != nil {
//			ShowCommandHelp(context, self.Name)
//			context.CLI.handleExitCoder(context, err)
//			return err
//		}
//	}
//}

//func (self *Command) ExecuteAfterAction() {
//	if self.After != nil {
//		afterErr := self.After(context)
//		if afterErr != nil {
//			log.Fatal(afterErr)
//			if err != nil {
//				err = NewMultiError(err, afterErr)
//			} else {
//				err = afterErr
//			}
//		}
//	}
//}

//func (self *Command) ExecuteAction() {
//	if self.Action == nil {
//		self.Action = helpSubcommand.Action
//		err = HandleAction(self.Action, context)
//		if err != nil {
//			context.CLI.handleExitCoder(context, err)
//		}
//	}
//}

//func (c *Command) parseFlags(args Args) (*flag.FlagSet, error) {
//	// TODO ?
//	set.SetOutput(ioutil.Discard)
//	// TODO: We dont skip flag parsing, we can just skip executing
//	// TODO: Parse and handle result in a switchase
//	err = set.Parse(args)
//	if err != nil {
//		return nil, err
//	}
//	return set, nil
//}

// TODO: Use this to then do map to command pointers then use all the various
// names to map to the same pointers then we can use this to do a ultra simple
// and quite fast lookup

//func (self Command) HasName(name string) bool {
//	if len(self.Aliases) == 0 {
//		return (self.Name == name)
//	} else {
//		for _, commandName := range self.Names() {
//			if commandName == name {
//				return true
//			}
//		}
//		return false
//	}
//}
