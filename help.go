package cli

import (
	"strings"

	template "github.com/multiverse-os/cli/framework/template"
	color "github.com/multiverse-os/cli/framework/terminal/ansi/color"
	style "github.com/multiverse-os/cli/framework/terminal/ansi/style"
	text "github.com/multiverse-os/cli/framework/text"
	banner "github.com/multiverse-os/cli/framework/text/banner"
)

type helpType int

const (
	applicationHelp helpType = iota
	commandHelp
)

// TODO: Since this is being generated from a template, to avoid wasting time,
// and ensuring the documentation is consistent, this should be output to a
// documentation that can be referrenced from the README.
func (self *CLI) renderCommandHelp(command Command) error {
	return self.renderHelpTemplate(commandHelp, command)
}
func (self *CLI) renderApplicationHelp() error {
	return self.renderHelpTemplate(applicationHelp, Command{})
}

func (self *CLI) renderHelpTemplate(renderType helpType, command Command) (err error) {
	helpOptions := map[string]string{
		"header":            self.asciiHeader("big"),
		"usage":             color.SkyBlue(style.Bold("Usage")),
		"availableCommands": color.SkyBlue(style.Bold("Available Commands")),
		"availableFlags":    color.SkyBlue(style.Bold("Flags")),
	}
	switch renderType {
	case applicationHelp:
		err = template.OutputStdOut(self.help(), helpOptions)
	case commandHelp:
		err = template.OutputStdOut(self.commandHelp(command), helpOptions)
	}
	if err != nil {
		return err
	} else {
		return nil
	}
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

func commandUsage(cmd Command) string {
	return "    " + style.Bold(cmd.usage()) + strings.Repeat(" ", (18-len(cmd.usage()))) + style.Dim(cmd.Description) + "\n"
}

func flagUsage(flag Flag) (output string) {
	if !IsNil(flag.DefaultValue) {
		if !IsBlank(flag.DefaultValue.(string)) {
			output = " [Default: " + flag.DefaultValue.(string) + "]"
		}
	}
	return "    " + style.Bold(flag.usage()) + strings.Repeat(" ", (18-len(flag.usage()))) + style.Dim(flag.Description) + output + "\n"
}

// TODO: Create the below variant as an option and store these options in their
// own subpackages just like with spinners and loaders in text library.
// TODO: Check if a template is speicifed otherwise show default. Or rather
// probably cache?
// TODO: The amount of code duplication sucks. It also doesn't support
// templating easily. Really need a better way.
///////////////////////////////////////////////////////////////////////////////
func (self *CLI) help() (t string) {
	t += "{{.header}}"
	t += "  {{.usage}}:\n"
	t += "    " + color.Fuchsia(style.Bold(self.Name)) + "  " + style.Dim("[command]") + "\n\n"
	t += "  {{.availableCommands}}:\n"
	for _, command := range self.Commands {
		t += commandUsage(command)
	}
	t += "\n"
	t += "  {{.availableFlags}}:\n"
	for _, flag := range self.Flags {
		t += flagUsage(flag)
	}
	t += "\n"

	return t
}

// TODO: If default value is provided, should indicate this
func (self *CLI) commandHelp(command Command) (t string) {
	t += "{{.header}}"
	t += "  {{.usage}}:\n"
	t += "    " + style.Bold(color.Fuchsia(self.Name)+" "+color.SkyBlue(command.Name)) + " " + style.Dim("[command]") + "\n\n"
	t += "  {{.availableCommands}}:\n"
	for _, command := range command.visibleSubcommands() {
		t += commandUsage(command)
	}
	t += "\n"
	t += "  {{.availableFlags}}:\n"
	for _, flag := range self.Flags {
		t += flagUsage(flag)
	}
	t += "\n"

	return t
}
