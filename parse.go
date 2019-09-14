package cli

import (
	"strings"
)

type FlagType int

const (
	BoolFlag FlagType = iota
	IntFlag
	StringFlag
	PathFlag
	FilenameFlag
)

type Flag struct {
	Name    string // Primary name
	Aliases []string
	Type    FlagType
	Usage   string
	Hidden  bool
	Value   interface{}
}

func (self Flag) Visible() bool { return !self.Hidden }

func (self Flag) Alias() string {
	if len(self.Aliases) > 0 {
		if len(self.Aliases[0]) >= 2 {
			return "--" + self.Aliases[0]
		} else {
			return "-" + self.Aliases[0]
		}
	} else {
		return ""
	}
}

func (self Command) visibleFlags() (flags []Flag) {
	for _, flag := range self.Flags {
		if flag.Visible() {
			flags = append(flags, flag)
		}
	}
	return append(flags, defaultCommandFlags()...)
}

func defaultFlags() []Flag {
	return []Flag{
		Flag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "Print version",
			Hidden:  false,
		},
		Flag{
			Name:    "help",
			Aliases: []string{"h"},
			Usage:   "Print help text",
			Hidden:  false,
		},
	}
}

func (self Flag) Is(name string) bool {
	for _, flagName := range self.Names() {
		if flagName == name {
			return true
		}
	}
	return false
}

func (flag Flag) Names() []string { return append([]string{flag.Name}, flag.Aliases...) }

// TODO: Are hooks really necessary? Maybe it would be better to just implement
// a middleware like functionality and push this even closer to being more like
// web development to make it easier to comphrehend and extend
// TODO: Why do we have 'Usage' AND 'UsageText' seems like we should be merging this in some way. Also is this diff than description?
type Command struct {
	Hidden      bool
	Category    int
	Name        string
	Aliases     []string
	Subcommands []Command
	Flags       []Flag
	Usage       string
	Action      func(c *Context) error
}

func (self Command) Visible() bool { return !self.Hidden }

func (self *Command) visibleSubcommands() (commands []Command) {
	for _, command := range self.Subcommands {
		if command.Visible() {
			commands = append(commands, command)
		}
	}
	return append(commands, defaultSubcommands()...)
}

func defaultCommands() []Command {
	return []Command{
		Command{
			Hidden:  true,
			Name:    "help",
			Aliases: []string{"h"},
			Usage:   "List of available commands or details for a specified command",
			Action: func(c *Context) error {
				c.CLI.renderHelp()
				return nil
			},
		},
		Command{
			Hidden:  true,
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "Display the version number, and other compile details",
			Action: func(c *Context) error {
				c.CLI.renderVersion()
				return nil
			},
		},
	}
}

func defaultCommandFlags() []Flag {
	return []Flag{
		Flag{
			Name:    "help",
			Aliases: []string{"h"},
			Usage:   "Print help text",
			Hidden:  false,
		},
	}
}

func defaultSubcommands() []Command {
	return []Command{
		Command{
			Name:    "help",
			Aliases: []string{"h"},
			Usage:   "List of available commands or details for a specified command",
		},
	}
}

func (self Command) Is(name string) bool {
	for _, commandName := range self.Names() {
		if commandName == name {
			return true
		}
	}
	return false
}

func (self Command) Empty() bool     { return len(self.Name) == 0 }
func (self Command) NotEmpty() bool  { return !self.Empty() }
func (self Command) Names() []string { return append([]string{self.Name}, self.Aliases...) }

func (self *CLI) parse(arguments []string) *Context {
	var skipArgument bool

	context := &Context{
		CLI:        self,
		Flags:      map[string]Flag{},
		Command:    Command{},
		Subcommand: Command{},
	}

	// TODO: Decide if flags before command should be global or if flags will in
	// general get ran by globals then command flags regardless of placement
	for index, argument := range arguments {
		if skipArgument {
			skipArgument = false
			continue
		}
		if string(argument[0]) == "-" || len(argument) > 2 && argument[:2] == "--" {
			argument = strings.ReplaceAll(argument, "-", "")
			var flagName string
			var flagValue interface{}
			if strings.Contains(argument, "=") {
				flagParts := strings.Split(argument, "=")
				flagName = flagParts[0]
				flagValue = flagParts[1]
			} else {
				skipArgument = true
				flagName = argument
				if len(arguments) > (index + 1) {
					flagValue = arguments[index+1]
				} else {
					flagValue = true
				}
			}
			if context.Command.NotEmpty() {
				ok, flag := self.isCommandFlag(context.Command, flagName)
				if ok {
					flag.Value = flagValue
					context.Flags[flag.Name] = flag
				}
			} else {
				ok, flag := self.isFlag(flagName)
				if ok {
					flag.Value = flagValue
					context.Flags[flag.Name] = flag
				}
			}
		} else {
			if context.Command.Empty() {
				ok, command := self.isCommand(argument)
				if ok {
					context.Command = command
				}
			} else {
				ok, Subcommand := self.isSubcommand(context.Command, argument)
				if ok {
					context.Subcommand = Subcommand
				}
			}
		}
	}
	return context
}
