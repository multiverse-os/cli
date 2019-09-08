package cli

import (
	template "github.com/multiverse-os/cli/template"
	text "github.com/multiverse-os/cli/text"
	color "github.com/multiverse-os/cli/text/ansi/color"
	style "github.com/multiverse-os/cli/text/ansi/style"
	//banner "github.com/multiverse-os/cli/text/ascii/banner"
)

func (self *CLI) renderUI() error {
	err := template.OutputStdOut(defaultTemplate(), map[string]string{
		"title":       self.header(),
		"description": self.Description,
	})
	if err != nil {
		return err
	}
	return nil
}

func defaultTemplate() string {
	return `{{.title}}
{{.description}}
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
	//banner := banner.RectanglesFont(self.Name)
	//version := text.Brackets(color.White("v") + style.Dim(self.Version.String()))
	//return banner.String()[:(len(banner.String())-(len(self.Version.String())+4))] +
	//	version + "\n" +
	//	style.Dim(text.Repeat("=", banner.Width))
	title := color.White(style.Bold(self.Name)) + "    " + text.Brackets(self.Version.StringWithANSI())
	return title + "\n" + style.Dim(text.Repeat("=", len(title)))
}
