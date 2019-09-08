package cli

import (
	"fmt"

	template "github.com/multiverse-os/cli/template"
	text "github.com/multiverse-os/cli/text"
	color "github.com/multiverse-os/cli/text/ansi/color"
	style "github.com/multiverse-os/cli/text/ansi/style"
)

func (self *CLI) header() string {
	return " " + (color.White(self.Name) + "    v" + style.Dim(self.Version.String()))
}

func (self *CLI) renderUI() error {
	output, err := template.OutputString(defaultTemplate(), map[string]string{
		"title": self.header(),
		"line":  text.Repeat("=", len(self.header())),
	})
	if err != nil {
		return err
	}
	fmt.Println("output:", output)
	return nil
}

func defaultTemplate() string {
	return `
{{.title}}
{{.line}}
{{.description}}
`
}
