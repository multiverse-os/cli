package cli

import (
	"fmt"

	template "github.com/multiverse-os/cli/template"
	//text "github.com/multiverse-os/cli/text"
	//color "github.com/multiverse-os/cli/text/ansi/color"
	//style "github.com/multiverse-os/cli/text/ansi/style"
)

func (self *CLI) renderUI() error {
	output, err := template.OutputString(defaultTemplate(), map[string]string{
		"title": self.header(),
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
