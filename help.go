package cli

import (
	"strings"

	template "github.com/multiverse-os/cli/template"
	text "github.com/multiverse-os/cli/text"
	color "github.com/multiverse-os/cli/text/ansi/color"
	style "github.com/multiverse-os/cli/text/ansi/style"
	banner "github.com/multiverse-os/cli/text/banner"
)

// TODO: Since this is being generated from a template, to avoid wasting time,
// and ensuring the documentation is consistent, this should be output to a
// documentation that can be referrenced from the README.

func (self *CLI) renderCommandHelp(command Command) error {
	err := template.OutputStdOut(defaultHelpTemplate(self.Name, command.visibleSubcommands(), command.visibleFlags()), map[string]string{
		"header":            self.header(true),
		"description":       self.Description,
		"usage":             color.SkyBlue(style.Bold("Usage")),
		"availableCommands": color.SkyBlue(style.Bold("Available Commands")),
		"availableFlags":    color.SkyBlue(style.Bold("Flags")),
	})
	if err != nil {
		return err
	}
	return nil
}

func (self *CLI) renderHelp() error {
	err := template.OutputStdOut(defaultHelpTemplate(self.Name, self.visibleCommands(), self.visibleFlags()), map[string]string{
		"header":            self.header(true),
		"description":       self.Description,
		"usage":             color.SkyBlue(style.Bold("Usage")),
		"availableCommands": color.SkyBlue(style.Bold("Available Commands")),
		"availableFlags":    color.SkyBlue(style.Bold("Flags")),
	})
	if err != nil {
		return err
	}
	return nil
}

// TODO: Create the below variant as an option and store these options in their
// own subpackages just like with spinners and loaders in text library.
///////////////////////////////////////////////////////////////////////////////
func defaultHelpTemplate(name string, commands []Command, flags []Flag) (t string) {
	t += "{{.header}}\n"
	t += "  {{.usage}}:\n"
	t += "    " + color.Fuchsia(style.Bold(name)) + "  " + style.Dim("[command]") + "\n\n"
	t += "  {{.availableCommands}}:\n"
	for _, command := range commands {
		t += "    " + style.Bold(command.Name) + strings.Repeat(" ", (18-len(command.Name))) + style.Dim(command.Usage) + "\n"
	}
	t += "\n"
	t += "  {{.availableFlags}}:\n"
	for _, flag := range flags {
		t += "    " + style.Bold(flag.Alias()+", --"+flag.Name) + strings.Repeat(" ", (12-len(flag.Name))) + style.Dim(flag.Usage) + "\n"
	}
	t += "\n"

	return t
}

// Available Banners Fonts
///////////////////////////////////////////////////////////////////////////////
// Big, Chunky, CyberLarge, CyberMedium, Doom, Elite, Isometric3, Isometric4
// Larry3D, Letters, NancyJ, Rectangles, Relief, Small, Smisome1, Standard
// Ticks, TicksSlant, calvins
func (self *CLI) header(showVersion bool) string {
	banner := banner.New("calvins", " "+self.Name)
	var version string
	if showVersion {
		version = text.Brackets(self.Version.String())
	}
	return style.Bold(color.SkyBlue(banner.String()[:len(banner.String())-1])) + version + "\n"
}
