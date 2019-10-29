package cli

import (
	"strings"

	template "github.com/multiverse-os/cli/framework/template"
	color "github.com/multiverse-os/cli/framework/terminal/ansi/color"
	style "github.com/multiverse-os/cli/framework/terminal/ansi/style"
	banner "github.com/multiverse-os/cli/framework/text/banner"
)

func (self *CLI) RenderHelpTemplate(command Command) (err error) {
	helpOptions := map[string]string{
		"header":            self.asciiHeader("big"),
		"usage":             color.SkyBlue(style.Bold("Usage")),
		"availableCommands": color.SkyBlue(style.Bold("Commands")),
		"availableFlags":    color.SkyBlue(style.Bold("Flags")),
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
	return style.Bold(color.SkyBlue(banner.String())) + self.Version.ColorString() + "\n"
}

func (self *CLI) simpleHeader() string {
	return style.Bold(color.SkyBlue(self.Name)) + "[v" + self.Version.ColorString() + "]\n"
}

func (self *CLI) helpTemplate(command Command) (t string) {
	path := command.path()
	t += "\n{{.header}}"
	t += "  {{.usage}}\n"
	if len(path) == 0 {
		t += "    " + style.Bold(color.Fuchsia(self.Name)) + " " + color.Silver(style.Dim("[parameters]")) + "\n\n"
	} else if len(path) == 1 {
		t += "    " + style.Bold(color.Fuchsia(self.Name)) + " " + color.SkyBlue(style.Dim("[command]")) + " " + color.Silver(style.Dim("[parameters]")) + "\n\n"
	} else {
		t += "    " + style.Bold(color.Fuchsia(self.Name)) + " " + strings.Join(command.path()[1:], " ") + " " + color.SkyBlue(style.Dim("[subcommand]")) + color.Silver(style.Dim("[parameters]")) + "\n\n"
	}
	t += "  {{.availableCommands}}\n"
	for index, subcommand := range command.visibleSubcommands() {
		t += "    " + style.Bold(subcommand.usage()) + strings.Repeat(" ", (18-len(subcommand.usage()))) + style.Dim(subcommand.Description)
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
		t += "    " + style.Bold(flag.usage()) + strings.Repeat(" ", (18-len(flag.usage()))) + style.Dim(flag.Description) + output + "\n"
		if index != len(command.visibleFlags())-1 {
			t += "\n"
		}
	}
	t += "\n"

	return t
}
