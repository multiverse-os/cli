package cli

import (
	"strings"

	template "github.com/multiverse-os/cli/template"
	text "github.com/multiverse-os/cli/text"
	color "github.com/multiverse-os/cli/text/ansi/color"
	style "github.com/multiverse-os/cli/text/ansi/style"
	banner "github.com/multiverse-os/cli/text/banner"
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
		"header":            self.header(true),
		"usageDescription":  self.Usage,
		"usage":             color.SkyBlue(style.Bold("Usage")),
		"availableCommands": color.SkyBlue(style.Bold("Available Commands")),
		"availableFlags":    color.SkyBlue(style.Bold("Flags")),
	}
	switch renderType {
	case applicationHelp:
		err = template.OutputStdOut(templates.DefaultHelp(self.Name, self.visibleCommands(), self.visibleFlags()), helpOptions)
	case commandHelp:
		err = template.OutputStdOut(templates.DefaultHelp(self.Name, command.visibleSubcommands(), command.visibleCommandFlags()), helpOptions)
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
func (self *CLI) header(showVersion bool) string {
	// TODO: Calling the banner.New() MUST be moved into separate template file
	// because its kinda bullky since it calls in a bunch of fonts currently. And
	// ideally we want to avoid calling it in if we don't use it. To do that we
	// move it out of the package and call that package if this is != 0
	// This new file could handle all the various templates. Which should also
	// include version output
	if len(self.HelpHeader) == 0 {
		banner := banner.New("calvins", " "+self.Name)
		var version string
		if showVersion {
			version = text.Brackets(self.Version.String())
		}
		return style.Bold(color.SkyBlue(banner.String()[:len(banner.String())-1])) + version + "\n"
	} else {
		return self.HelpHeader
	}
}

// TODO: Create the below variant as an option and store these options in their
// own subpackages just like with spinners and loaders in text library.
///////////////////////////////////////////////////////////////////////////////
func defaultHelp(name string, commands []Command, flags []Flag) (t string) {
	t += "{{.header}}"
	t += "  {{.usage}}:\n"
	t += "    " + color.Fuchsia(style.Bold(name)) + "  " + style.Dim("[command]") + "\n\n"
	if len(commands) > 0 {
		t += "  {{.availableCommands}}:\n"
		for _, command := range commands {
			t += "    " + style.Bold(command.NameHelpString()) + strings.Repeat(" ", (18-len(command.NameHelpString()))) + style.Dim(command.Usage) + "\n"
		}
		t += "\n"
	}
	t += "  {{.availableFlags}}:\n"
	for _, flag := range flags {
		t += "    " + style.Bold(flag.NameHelpString()) + strings.Repeat(" ", (18-len(flag.NameHelpString()))) + style.Dim(flag.Usage) + "\n"
	}
	t += "\n"

	return t
}
