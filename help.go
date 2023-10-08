package cli

import (
	"fmt"
	"strings"

	banner "github.com/multiverse-os/banner"

	data "github.com/multiverse-os/cli/data"
	ansi "github.com/multiverse-os/cli/terminal/ansi"
	template "github.com/multiverse-os/cli/terminal/template"
)

// TODO: Do a lot of cleanup; like there is no more expecting subcommand concept
// and I already did some work to make it more functional

// TODO: Consider a help or template object, then we easily assign things like
// indentation, figlet font or not, and make everything here a method. then it
// will be super easy to customize the output. be able to pass a go template or
// definte the various aspects.
// TODO: Make render help fit the type for Action so that it can be assigned to
// the help commands action for greater simplicity and less hard-coding.

// TODO: Can we rebuild this with a matrix grid like? https://github.com/sh0rez/asciimatrix/blob/master/asciimatrix.go
// I bet we can get 100x cleaner code, end up with smaller code, and maximum
// customability easily

func RenderDefaultHelpTemplate(context *Context) error {
	// NOTE: This is important for localization
	helpOptions := map[string]string{
		"header":      context.bannerHeader("chunky"),
		"description": context.Commands.First().Description,
		"usage":       "Usage",
		"subcommands": "Subcommands",
		"global":      "Global",
		"flags":       "Flags",
		"options":     "options",
		"command":     "command",
		"subcommand":  "subcommand",
		"params":      "parameters",
	}
	return template.StdOut(context.defaultHelpTemplate(), helpOptions)
}

func RenderDefaultVersionTemplate(context *Context) error {
	err := template.StdOut(context.defaultVersionTemplate(), map[string]string{
		"header":  ansi.Bold(ansi.SkyBlue(context.Commands.First().Name)),
		"version": context.CLI.Version.ColorString(),
	})
	//"build": table.New(BuildInformation{
	//	Source:     "n/a",
	//	Commit:     "n/a",
	//	Signature:  "n/a",
	//	CompiledAt: "n/a",
	//}).String(),
	if data.NotNil(err) {
		return err
	}
	return nil
}

func (self Context) defaultVersionTemplate() string {
	return "{{.header}}" + ansi.SkyBlue(ansi.Light(" version ")) + "{{.version}}\n"
}

// Available Banners Fonts
// /////////////////////////////////////////////////////////////////////////////
// Big, Chunky, CyberLarge, CyberMedium, Doom, Elite, Isometric3, Isometric4
// Larry3D, Letters, NancyJ, Rectangles, Relief, Small, Smisome1, Standard
// Ticks, TicksSlant, calvins
// TODO: Should probably make an enumerator
func (self Context) bannerHeader(font string) string {
	return banner.New(" "+self.Commands.First().Name).Font(font).String() +
		self.CLI.Version.String()
}

func (self Context) simpleHeader() string {
	return self.Commands.First().Name + "[v" + self.CLI.Version.String() + "]"
}

// TODO: Would be preferable to define a template and use it than have a static
//
//	template like this. This could be the default fallback.
func (self Context) defaultHelpTemplate() (t string) {
	t += "\n{{.header}}\n  {{.description}}\n\n  "
	t += ansi.Bold("{{.usage}}")
	// TODO: Usage needs to be fixed, after we minimized it a bit
	t += ansi.Light(" " + strings.Join(self.Commands.Names(), " "))
	t += ansi.Gray(" [{{.options}}]")
	if !self.Commands.Last().Subcommands.IsZero() {
		t += ansi.Gray(" [{{.subcommand}}]")
	}
	t += ansi.Gray(" [{{.params}}]\n\n")

	if !self.Commands.Last().Subcommands.IsZero() {
		t += ansi.Bold("  {{.subcommands}}\n")
		for _, subcommand := range self.Commands.Last().Subcommands.Reverse().Visible() {
			t += commandUsage(*subcommand)
		}
		t += "\n"
	}

	// TODO: This method will not ever print command flags, and so this has been
	// broken fundamentally
	if len(self.Commands.First().Flags) != 0 {
		t += ansi.Bold("  {{.flags}}\n")
		if !self.Commands.Last().Flags.IsZero() && self.Commands.Count() != 1 {
			t += ansi.Bold("   Command options\n")
			for _, flag := range self.Commands.Last().Flags.Visible().Reverse() {
				if !flag.HasCategory() {
					t += flagHelp(*flag)
				}
			}
			t += "\n"
		}
		t += ansi.Bold("   Global options\n")
		for _, flag := range self.Commands.First().Flags.Visible().Reverse() {
			if !flag.HasCategory() {
				t += flagHelp(*flag)
			}
		}
		t += "\n"
		for _, category := range self.Commands.First().Flags.Categories() {
			t += ansi.Bold(fmt.Sprintf("   %v options\n", category))
			for _, flag := range self.Commands.First().Flags.Category(category).Visible() {
				t += flagHelp(*flag)
			}
			t += "\n"
		}

	}

	return t
}

///////////////////////////////////////////////////////////////////////////////

func commandUsage(command Command) string {
	usage := command.Name
	if len(command.Alias) != 0 {
		usage += ", " + command.Alias
	}
	return "    " + usage + strings.Repeat(" ", 18-len(usage)) +
		ansi.Gray(command.Description) + "\n"
}

func flagHelp(flag Flag) string {
	var usage string
	if len(flag.Alias) != 0 {
		usage += Short.String() + flag.Alias + ", "
	}
	usage += Long.String() + flag.Name
	var defaultValue string
	if len(flag.Default) != 0 {
		defaultValue = " [≅ " + flag.Default + "]"
	}
	return "    " + usage + strings.Repeat(" ", 18-len(usage)) +
		ansi.Gray(flag.Description) + ansi.Bold(defaultValue) + "\n"
}
