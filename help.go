package cli

import (
	"strings"

	banner "github.com/multiverse-os/banner"
	template "github.com/multiverse-os/cli/terminal/template"
  data "github.com/multiverse-os/cli/data"
)

func (self *Context) RenderHelpTemplate(command *Command) error {
	helpOptions := map[string]string{
		"header":            self.CLI.asciiHeader("big"),
		"usage":             "Usage",
		"availableCommands": "Commands",
		"availableFlags":    "Flags",
	}
	return template.StdOut(self.helpTemplate(command), helpOptions)
}

// Available Banners Fonts
///////////////////////////////////////////////////////////////////////////////
// Big, Chunky, CyberLarge, CyberMedium, Doom, Elite, Isometric3, Isometric4
// Larry3D, Letters, NancyJ, Rectangles, Relief, Small, Smisome1, Standard
// Ticks, TicksSlant, calvins
func (self *CLI) asciiHeader(font string) string {
	banner := banner.New(" " + self.Context.Commands.First().Name).Font(font)
	return banner.String() +
         self.Version.String() + 
         "\n"
}

func (self *CLI) simpleHeader() string {
	return self.Context.Commands.First().Name + 
         "[v" + 
         self.Version.String() + 
         "]\n"
}

// TODO: Maybe default to just having command and then doing some sort of simple
// check to add sub? something easier than this possible?
func (self *Context) expectingCommandsOrSubcommand() string {
	if self.Commands.First().Subcommands.IsZero() {
		return " [command]"
	} else if 2 < self.Commands.Count() {
		return " [subcommand]"
	} else {
		return ""
	}
}

// TODO: Would be preferable to define a template and use it than have a static
//       template like this. This could be the default fallback.
func (self *Context) helpTemplate(command *Command) (t string) {
	t += "\n{{.header}}"
	t += Prefix() + "{{.usage}}\n"
	t += Tab() + 
       strings.ToLower(strings.Join(self.Commands.Names(), " ")) + 
       strings.ToLower(self.expectingCommandsOrSubcommand()) + 
       " [parameters]" + 
       "\n\n"
	t += Prefix() + 
       "{{.availableCommands}}\n"
	for _, subcommand := range command.Subcommands.Visible() {
		t += Tab() + 
         subcommand.usage() + 
        strings.Repeat(" ", (18-len(subcommand.usage()))) +
        subcommand.Description
	}
	t += "\n\n\n"

	// TODO: Should the command flags be printed with global flags too?
	for _, command := range self.Commands {
		if len(command.Flags) != 0 {
			if command.Base() {
				t += Prefix() +
             "{{.availableFlags}}\n"
			} else {
				t += Prefix() + 
             "Global {{.availableFlags}}\n"
			}
			for _, flag := range command.Flags {
				t += flagHelp(*flag)
			}
			t += "\n"
		}
	}

	return t
}


func flagHelp(flag Flag) string {
	usage := Long.String() + 
           flag.Name
	if data.NotBlank(flag.Alias) {
		usage += ", " +
             Short.String() +
             flag.Alias
	}
	var defaultValue string
	if len(flag.Default) != 0 {
		defaultValue = " [≅ " +
                  flag.Default +
                  "]"
	}
	return strings.Repeat(" ", 4) +
         usage +
         strings.Repeat(" ", 18-len(usage)) +
         flag.Description + defaultValue +
         "\n"
}

