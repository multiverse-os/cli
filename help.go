package cli

import (
	"fmt"
	"strings"

	template "github.com/multiverse-os/cli/template"

	banner "github.com/multiverse-os/banner"
)

func (self *CLI) RenderHelpTemplate(command *Command) (err error) {
	helpOptions := map[string]string{
		"header":            self.asciiHeader("big"),
		"usage":             "Usage",
		"availableCommands": "Commands",
		"availableFlags":    "Flags",
	}
	return template.OutputStdOut(self.helpTemplate(command), helpOptions)
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

// TODO: This is pretty slow think about how this can be sped up
func (self *CLI) helpTemplate(command *Command) (t string) {
	path := command.path()
	fmt.Println("path:", path)
	t += "\n{{.header}}"
	t += "  {{.usage}}\n"
	if len(path) == 0 {
		t += "    " + self.Name + " " + "[parameters]" + "\n\n"
	} else if len(path) == 1 {
		t += "    " + command.Name + " " + "[command]" + " " + "[parameters]" + "\n\n"
	} else {
		t += "    " + self.Name + " " + strings.Join(command.path()[1:], " ") + " " + "[subcommand]" + "[parameters]" + "\n\n"
	}
	t += "  {{.availableCommands}}\n"
	for index, subcommand := range command.visibleSubcommands() {
		t += "    " + subcommand.usage() + strings.Repeat(" ", (18-len(subcommand.usage()))) + subcommand.Description
		if index != len(command.visibleSubcommands())-1 {
			t += "\n"
		}
	}
	t += "\n\n"
	t += "  {{.availableFlags}}\n"
	for index, flag := range self.Flags {
		var output string
		if len(flag.Default) != 0 {
			output = " [≅ " + flag.Default + "]"
		}
		t += "    " + flag.usage() + strings.Repeat(" ", (18-len(flag.usage()))) + flag.Description + output + "\n"
		if index != len(command.visibleFlags())-1 {
			t += "\n"
		}
	}
	t += "\n"

	return t
}
