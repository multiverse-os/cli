package cli

import (
	template "github.com/multiverse-os/cli/template"
	text "github.com/multiverse-os/cli/text"
	color "github.com/multiverse-os/cli/text/ansi/color"
	style "github.com/multiverse-os/cli/text/ansi/style"
	banner "github.com/multiverse-os/cli/text/ascii/banner"
)

// TODO: Since this is being generated from a template, to avoid wasting time,
// and ensuring the documentation is consistent, this should be output to a
// documentation that can be referrenced from the README.

func (self *CLI) renderHelp() error {
	err := template.OutputStdOut(defaultTemplate(), map[string]string{
		"header":            self.header(),
		"name":              self.Name,
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

func defaultTemplate() string {
	return `{{.header}}
  {{.usage}}:
    ` + color.Fuchsia(style.Bold(`{{.name}}`)) + ` ` + style.Dim(`[command]`) + `
  
  {{.availableCommands}}:
    ` + style.Bold(`help`) + `       ` + style.Dim(`Display help text, specify a command for in depth command help`) + `
    ` + style.Bold(`version`) + `    ` + style.Dim(`Display version, and compiler information`) + `
  
  {{.availableFlags}}:
    ` + style.Bold(`-h`) + `, ` + style.Bold(`--help`) + `      ` + style.Dim(`help for {{.name}}`) + `
        ` + style.Bold(`--version`) + `   ` + style.Dim(`version for {{.name}}`) + `

`
}

// Available Banners Fonts
///////////////////////////////////////////////////////////////////////////////
// Big, Chunky, CyberLarge, CyberMedium, Doom, Elite, Isometric3, Isometric4
// Larry3D, Letters, NancyJ, Rectangles, Relief, Small, Smisome1, Standard
// Ticks, TicksSlant

func (self *CLI) header() string {
	banner := banner.RectanglesFont(self.Name)
	version := text.Brackets(color.White("v") + style.Dim(self.Version.String()))
	return style.Bold(color.SkyBlue(banner.StringWithPrefix("  ")[:len(banner.String())+6])) + version
	//title := color.White(style.Bold(self.Name)) + "    " + text.Brackets(self.Version.StringWithANSI())
	//return title + "\n" + style.Dim(text.Repeat("=", len(title)))
}
