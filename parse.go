package cli

type Input struct {
	CLI     *CLI
	Command *Command
	Flags   []*Flag
}

func LoadInput(cli *CLI, command *Command, flags []*Flag) *Input {
	return &Input{CLI: cli, Command: command, Flags: flags}
}

type Flags []Flag
type Commands []Command

// TODO: Why do we have 'Usage' AND 'UsageText' seems like we should be merging this in some way. Also is this diff than description?
type Command struct {
	Hidden         bool
	Category       int
	Name           string
	Aliases        []string
	ParentCommand  *Command
	Subcommands    map[string]Command
	Flags          map[string]Flag
	Usage          string
	SkipArgReorder bool
	Action         interface{}
	Before         func()
	After          func()
}

func (self Command) HasSubcommands() bool { return (len(self.Subcommands) > 0) }
func (self Command) Run() (err error)     { return err }
func (self Command) Names() []string      { return append([]string{self.Name}, self.Aliases...) }

type FlagType int

func defaultCommands() []Command {
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

const (
	BoolFlag FlagType = iota
	IntFlag
	StringFlag
	PathFlag
	FilenameFlag
)

func flagPrefix(name string) string {
	if len(name) == 1 { // TODO: And two?
		return "-"
	} else {
		return "--"
	}
}

type Flag struct {
	Name    string // Primary name
	Aliases []string
	Type    FlagType
	Usage   string
	Hidden  bool
	Value   interface{}
}

func (flag Flag) Names() []string { return append([]string{flag.Name}, flag.Aliases...) }

var VersionFlag Flag = Flag{
	Name:    "version",
	Aliases: []string{"v"},
	Usage:   "Print version",
	Hidden:  true,
}

var HelpFlag Flag = Flag{
	Name:    "help",
	Aliases: []string{"h"},
	Usage:   "Print help text",
	Hidden:  true,
}
