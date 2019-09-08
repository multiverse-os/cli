package cli

import (
	template "github.com/multiverse-os/cli/template"
	text "github.com/multiverse-os/cli/text"
	color "github.com/multiverse-os/cli/text/ansi/color"
	style "github.com/multiverse-os/cli/text/ansi/style"
	banner "github.com/multiverse-os/cli/text/ascii/banner"
)

func (self *CLI) renderUI() error {
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

// Available Banners
///////////////////////////////////////////////////////////////////////////////
// BigFont(text string) Banner
// ChunkyFont(text string) Banner
// CyberLargeFont(text string) Banner
// CyberMediumFont(text string) Banner
// DoomFont(text string) Banner
// EliteFont(text string) Banner
// Isometric3Font(text string) Banner
// Isometric4Font(text string) Banner
// Larry3DFont(text string) Banner
// LettersFont(text string) Banner
// NancyJFont(text string) Banner
// RectanglesFont(text string) Banner
// ReliefFont(text string) Banner
// SmallFont(text string) Banner
// Smisome1Font(text string) Banner
// StandardFont(text string) Banner
// TicksFont(text string) Banner
// TicksSlantFont(text string) Banner

func (self *CLI) header() string {
	banner := banner.RectanglesFont(self.Name)
	version := text.Brackets(color.White("v") + style.Dim(self.Version.String()))
	return style.Bold(color.SkyBlue(banner.StringWithPrefix("  ")[:len(banner.String())+6])) + version

	//title := color.White(style.Bold(self.Name)) + "    " + text.Brackets(self.Version.StringWithANSI())
	//return title + "\n" + style.Dim(text.Repeat("=", len(title)))
}
