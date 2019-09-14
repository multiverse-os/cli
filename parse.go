package cli

import (
	"fmt"
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

func (self Command) HasSubcommands() bool { return len(self.Subcommands) > 0 }
func (self Command) Names() []string      { return append([]string{self.Name}, self.Aliases...) }

func (self *CLI) parse(arguments []string) *Context {
	var skipArgument bool

	context := &Context{
		CLI:        self,
		Flags:      []*Flag{},
		Command:    &Command{},
		Subcommand: &Command{},
	}

	// TODO: Decide if flags before command should be global or if flags will in
	// general get ran by globals then command flags regardless of placement
	for index, argument := range arguments {
		fmt.Println("argument:", argument)
		if skipArgument {
			skipArgument = false
			continue
		}
		if string(argument[0]) == "-" || argument[:2] == "--" {
			argument = strings.ReplaceAll(argument, "-", "")
			if strings.Contains(argument, "=") {
				flag := strings.Split(argument, "=")
				context.Flags = append(context.Flags, &Flag{Name: flag[0], Value: flag[1]})
			} else {
				skipArgument = true
				if len(arguments) > (index + 1) {
					context.Flags = append(context.Flags, &Flag{Name: argument, Value: arguments[index+1]})
				}
			}
		} else {
			if context.Command.isEmpty() {
				context.Command = &Command{Name: argument}
			} else {
				context.Subcommand = &Command{Name: argument}
			}
		}
	}
	return context
}

func (self *Command) isEmpty() bool { return len(self.Name) == 0 }
