package cli

import (
	"strings"

	style "github.com/multiverse-os/cli/framework/terminal/ansi/style"
)

type Command struct {
	Hidden      bool
	Category    int
	Name        string
	Aliases     []string
	Subcommands []Command
	Flags       []Flag
	Description string
	Action      func(c *Context) error
}

func (self Command) Help() string {
	return "    " +
		style.Bold(self.Usage()) +
		strings.Repeat(" ", (18-len(self.Usage()))) +
		style.Dim(self.Description) +
		"\n"
}

func (self Command) Visible() bool   { return !self.Hidden }
func (self Command) Empty() bool     { return len(self.Name) == 0 }
func (self Command) NotEmpty() bool  { return !self.Empty() }
func (self Command) Names() []string { return append([]string{self.Name}, self.Aliases...) }

func (self Command) Usage() (output string) {
	output += self.Name
	if len(self.Aliases) > 0 {
		if len(self.Aliases[0]) > 1 {
			output += ", " + self.Aliases[0]
		}
	}
	return output
}

func (self Command) Is(name string) bool {
	for _, commandName := range self.Names() {
		if commandName == name {
			return true
		}
	}
	return false
}

func (self Command) visibleFlags() (flags []Flag) {
	for _, flag := range self.Flags {
		if flag.Visible() {
			flags = append(flags, flag)
		}
	}
	return append(flags, defaultCommandFlags()...)
}

func (self Command) visibleSubcommands() (commands []Command) {
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
			Hidden:      true,
			Name:        "help",
			Aliases:     []string{"h"},
			Description: "List of available commands or details for a specified command",
			Action: func(c *Context) error {
				c.CLI.RenderApplicationHelp()
				return nil
			},
		},
		Command{
			Hidden:      true,
			Name:        "version",
			Aliases:     []string{"v"},
			Description: "Display the version number, and other compile details",
			Action: func(c *Context) error {
				c.CLI.RenderVersion()
				return nil
			},
		},
	}
}

func defaultSubcommands() []Command {
	return []Command{
		Command{
			Name:        "help",
			Aliases:     []string{"h"},
			Description: "List of available commands or details for a specified command",
		},
	}
}
