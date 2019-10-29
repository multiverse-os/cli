package cli

import (
	"strings"

	data "github.com/multiverse-os/cli/framework/data"
	template "github.com/multiverse-os/cli/framework/template"
	color "github.com/multiverse-os/cli/framework/terminal/ansi/color"
	style "github.com/multiverse-os/cli/framework/terminal/ansi/style"
	text "github.com/multiverse-os/cli/framework/text"
	banner "github.com/multiverse-os/cli/framework/text/banner"
)

// TODO: Since this is being generated from a template, to avoid wasting time,
// and ensuring the documentation is consistent, this should be output to a
// documentation that can be referrenced from the README.
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
	return (style.Bold(color.SkyBlue(banner.String())) + text.Brackets(self.Version.String()) + "\n")
}

func (self *CLI) simpleHeader() string {
	return style.Bold(color.SkyBlue(self.Name)) + text.Brackets(self.Version.String()) + "\n"
}

func flagUsage(flag Flag) (output string) {
	if data.NotNil(flag.Default) {
		if data.NotBlank(flag.Default) {
			output = " [≅ " + flag.Default + "]"
		}
	}
	return "    " + style.Bold(flag.Usage()) + strings.Repeat(" ", (18-len(flag.Usage()))) + style.Dim(flag.Description) + output + "\n"
}

// TODO: Create the below variant as an option and store these options in their
// own subpackages just like with spinners and loaders in text library.
// TODO: Check if a template is speicifed otherwise show default. Or rather
// probably cache?
// TODO: The amount of code duplication sucks. It also doesn't support
// templating easily. Really need a better way.
///////////////////////////////////////////////////////////////////////////////
// TODO: With the new structure where the CLI itself is a command, we can now easily merge these two helps
// TODO: If default value is provided, should indicate this
func (self *CLI) helpTemplate(command Command) (t string) {
	path := command.Path()
	t += "\n{{.header}}"
	t += "  {{.usage}}\n"
	if len(path) == 0 {
		t += "    " + style.Bold(color.Fuchsia(self.Name)) + " " + color.Silver(style.Dim("[parameters]")) + "\n\n"
	} else if len(path) == 1 {
		t += "    " + style.Bold(color.Fuchsia(self.Name)) + " " + color.SkyBlue(style.Dim("[command]")) + " " + color.Silver(style.Dim("[parameters]")) + "\n\n"
	} else {
		t += "    " + style.Bold(color.Fuchsia(self.Name)) + " " + strings.Join(command.Path()[1:], " ") + " " + color.SkyBlue(style.Dim("[subcommand]")) + color.Silver(style.Dim("[parameters]")) + "\n\n"
	}
	t += "  {{.availableCommands}}\n"
	for index, subcommand := range command.VisibleSubcommands() {
		t += "    " + style.Bold(subcommand.Usage()) + strings.Repeat(" ", (18-len(subcommand.Usage()))) + style.Dim(subcommand.Description)
		if index != len(command.VisibleSubcommands())-1 {
			t += "\n"
		}
	}
	t += "\n\n"
	t += "  {{.availableFlags}}\n"
	for _, flag := range self.Flags {
		t += flagUsage(flag)
	}
	t += "\n"

	return t
}
