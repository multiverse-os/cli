package cli

import (
	"strings"

	template "github.com/multiverse-os/cli/template"

	banner "github.com/multiverse-os/banner"
)

func (self *Context) RenderHelpTemplate() (err error) {
	helpOptions := map[string]string{
		"header":            self.CLI.asciiHeader("big"),
		"usage":             "Usage",
		"availableCommands": "Commands",
		"availableFlags":    "Flags",
	}
	return template.StdOut(self.helpTemplate(self.Command.Parent), helpOptions)
}

// Available Banners Fonts
///////////////////////////////////////////////////////////////////////////////
// Big, Chunky, CyberLarge, CyberMedium, Doom, Elite, Isometric3, Isometric4
// Larry3D, Letters, NancyJ, Rectangles, Relief, Small, Smisome1, Standard
// Ticks, TicksSlant, calvins
func (self *CLI) asciiHeader(font string) string {
	banner := banner.New(" " + self.Name).Font(font)
	return banner.String() + self.Version.String() + "\n"
}

func (self *CLI) simpleHeader() string {
	return self.Name + "[v" + self.Version.String() + "]\n"
}

func (self *Context) expectingCommandsOrSubcommand() string {
	if self.HasCommands() {
		return " [command]"
	} else if !self.Command.Base() {
		return " [subcommand]"
	} else {
		return ""
	}
}

// TODO: Would be preferable to define a template and use it than have a static
//       template like this. This could be the default fallback.
func (self *Context) helpTemplate(command *Command) (t string) {
	t += "\n{{.header}}"
	t += "  {{.usage}}\n"
	t += "    " + self.CommandChain.PathExample() + self.expectingCommandsOrSubcommand() + " [parameters]" + "\n\n"
	t += "  {{.availableCommands}}\n"
	for index, subcommand := range command.visibleSubcommands() {
		t += "    " + subcommand.usage() + strings.Repeat(" ", (18-len(subcommand.usage()))) + subcommand.Description
		if index != len(command.visibleSubcommands())-1 {
			t += "\n"
		}
	}
	t += "\n\n"

	// TODO: Should the command flags be printed with global flags too?
	t += "  {{.availableFlags}}\n"
	for _, flag := range self.CommandChain.Last().Flags {
		var output string
		if len(flag.Default) != 0 {
			output = " [≅ " + flag.Default + "]"
		}
		t += "    " + flag.usage() + strings.Repeat(" ", (18-len(flag.usage()))) + flag.Description + output + "\n"
	}
	t += "\n"
	return t
}
