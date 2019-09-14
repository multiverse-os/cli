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

func defaultFlags() []Flag {
	return []Flag{
		Flag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "Print version",
			Hidden:  true,
		},
		Flag{
			Name:    "help",
			Aliases: []string{"h"},
			Usage:   "Print help text",
			Hidden:  true,
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

func flagPrefix(name string) string {
	if len(name) >= 2 {
		return "--"
	} else {
		return "-"
	}
}

// TODO: Are hooks really necessary? Maybe it would be better to just implement
// a middleware like functionality and push this even closer to being more like
// web development to make it easier to comphrehend and extend
// TODO: Why do we have 'Usage' AND 'UsageText' seems like we should be merging this in some way. Also is this diff than description?
type Command struct {
	Hidden        bool
	Category      int
	Name          string
	Aliases       []string
	ParentCommand *Command
	Subcommands   map[string]Command
	Flags         map[string]Flag
	Usage         string
	Action        func(c *Context) error
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

func (self Command) InitSubcommands() []Command {
	return []Command{
		Command{
			Name:          "help",
			Aliases:       []string{"h"},
			Usage:         "List of available commands or details for a specified command",
			ParentCommand: &self,
			Action: func(c *Context) error {
				// TODO: Build out a template for command help which displays the
				// subcommands instead of the top level global commands
				//c.CLI.renderCommandHelp()
				return nil
			},
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

func (self Command) isEmpty() bool        { return len(self.Name) == 0 }
func (self Command) Names() []string      { return append([]string{self.Name}, self.Aliases...) }
func (self Command) HasSubcommands() bool { return len(self.Subcommands) > 0 }

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
			ok, flag := self.isFlag(flagName)
			if ok {
				flag.Value = flagValue
				context.Flags[flag.Name] = flag
			}
		} else {
			if context.Command.isEmpty() {
				ok, command := self.isCommand(argument)
				if ok {
					context.Command = command
				}
			} else {
				ok, subcommand := self.isSubcommand(context.Command, argument)
				if ok {
					context.Subcommand = subcommand
				}
			}
		}
	}
	return context
}
