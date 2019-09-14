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
	Action        interface{}
}

func (self Command) Is(name string) bool {
	for _, commandName := range self.Names() {
		if commandName == name {
			return true
		}
	}
	return false
}

func (self Command) Names() []string      { return append([]string{self.Name}, self.Aliases...) }
func (self Command) HasSubcommands() bool { return len(self.Subcommands) > 0 }

func (self *CLI) parse(arguments []string) *Context {
	var skipArgument bool

	context := &Context{
		CLI:        self,
		Flags:      []Flag{},
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
			var flagValue string
			if strings.Contains(argument, "=") {
				flagParts := strings.Split(argument, "=")
				flagName = flagParts[0]
				flagValue = flagParts[1]
			} else {
				skipArgument = true
				if len(arguments) > (index + 1) {
					flagName = argument
					flagValue = arguments[index+1]
				}
			}
			ok, flag := self.isFlag(flagName)
			if ok {
				flag.Value = flagValue
				context.Flags = append(context.Flags, flag)
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

func (self *Command) isEmpty() bool { return len(self.Name) == 0 }
