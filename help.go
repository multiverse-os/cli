package cli

import (
	//"strings"

	banner "github.com/multiverse-os/cli/framework/ascii/banner"
	template "github.com/multiverse-os/cli/framework/template"
	color "github.com/multiverse-os/cli/framework/terminal/ansi/color"
	style "github.com/multiverse-os/cli/framework/terminal/ansi/style"
	text "github.com/multiverse-os/cli/framework/text"
)

type helpType int

const (
	applicationHelp helpType = iota
	commandHelp
)

// TODO: Since this is being generated from a template, to avoid wasting time,
// and ensuring the documentation is consistent, this should be output to a
// documentation that can be referrenced from the README.
func (self *CLI) RenderCommandHelp(command Command) error {
	return self.RenderHelpTemplate(commandHelp, command)
}
func (self *CLI) RenderApplicationHelp() error {
	return self.RenderHelpTemplate(applicationHelp, Command{})
}

func (self *CLI) RenderHelpTemplate(renderType helpType, command Command) (err error) {
	helpOptions := map[string]string{
		"header":            self.asciiHeader("calvins"),
		"usageDescription":  self.Usage,
		"usage":             color.SkyBlue(style.Bold("Usage")),
		"availableCommands": color.SkyBlue(style.Bold("Available Commands")),
		"availableFlags":    color.SkyBlue(style.Bold("Flags")),
	}
	switch renderType {
	case applicationHelp:
		err = template.OutputStdOut(self.Help(), helpOptions)
	case commandHelp:
		err = template.OutputStdOut(self.CommandHelp(command), helpOptions)
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
	banner := banner.New(font, " "+self.Name)
	return style.Bold(color.SkyBlue(banner.String()[:len(banner.String())-1])) + text.Brackets(self.Version.String()) + "\n"
}

func (self *CLI) simpleHeader() string {
	return style.Bold(color.SkyBlue(self.Name)) + text.Brackets(self.Version.String()) + "\n"
}

// TODO: Create the below variant as an option and store these options in their
// own subpackages just like with spinners and loaders in text library.
// TODO: Check if a template is speicifed otherwise show default. Or rather
// probably cache?
///////////////////////////////////////////////////////////////////////////////
func (self *CLI) Help() (t string) {
	t += "{{.header}}"
	t += "  {{.usage}}:\n"
	t += "    " + color.Fuchsia(style.Bold(self.Name)) + "  " + style.Dim("[command]") + "\n\n"
	t += "  {{.availableCommands}}:\n"
	for _, command := range self.Commands {
		t += command.Help()
	}
	t += "\n"
	t += "  {{.availableFlags}}:\n"
	for _, flag := range self.Flags {
		t += flag.Help()
	}
	t += "\n"

	return t
}

func (self *CLI) CommandHelp(command Command) (t string) {
	t += "{{.header}}"
	t += "  {{.usage}}:\n"
	t += "    " + style.Bold(color.Fuchsia(self.Name)+" "+color.SkyBlue(command.Name)) + " " + style.Dim("[command]") + "\n\n"
	t += "  {{.availableCommands}}:\n"
	for _, command := range command.visibleSubcommands() {
		t += command.Help()
	}
	t += "\n"
	t += "  {{.availableFlags}}:\n"
	for _, flag := range self.Flags {
		t += flag.Help()
	}
	t += "\n"

	return t
}
