package cli

import (
	"fmt"
	"io"
	"strings"
	"text/tabwriter"
	"text/template"

	text "github.com/multiverse-os/cli/text"
	color "github.com/multiverse-os/cli/text/ansi/color"
	style "github.com/multiverse-os/cli/text/ansi/style"
)

// TODO: We should automatically generate things to be appended to README, or
// straight up docuentation to be added to each project just based on the
// ifnormation in here so we dont have to keep reepating the same things
func (self *CLI) PrintBanner() {
	title := color.White(self.Name) + "  " + style.Dim(self.Version.String())
	fmt.Printf(title + "\n" + text.Repeat("=", len(title)) + "\n")
}

type helpPrinter func(w io.Writer, templ string, data interface{})

var HelpPrinter helpPrinter = printHelp

func (self *CLI) PrintHelp() {
	if len(self.Description) != 0 {
		fmt.Println(style.Strong(self.Description))
	}
	if len(self.Usage) != 0 {
		if self.ANSI {
			fmt.Println("Usage")
		} else {
			fmt.Println(style.Strong("Usage"))
		}
		fmt.Print(text.Repeat(" ", 4))

		if self.ANSI {
			fmt.Print(self.Name)
		} else {
			fmt.Print(color.White(self.Name))
		}
		if self.visibleFlags {
			fmt.Print(" [options]")
		} else {
		}
		if self.visibleCommands {
			fmt.Print(" command [command options]")
		}
		if self.ArgsUsage != "" {
			fmt.Println(" " + self.ArgsUsage)
		} else {
			fmt.Println("[arguments...]")
		}
		if self.HasVisibleFlags() {
			fmt.Println("\n" + style.Strong("Options"))
			for _, flag := range self.VisibleFlags() {
				fmt.Println("flag: ", flag)
			}
		}
	}
}

// TODO: Migrate this to markdown file, to simplify modifciations,
// customziations and allow for drop in replacements. Then several varations
// could be supplied.
var CLIHelpTemplate = `{{range $index, $option := .VisibleFlags}}{{if $index}}{{"\n"}}{{end}}{{"\t\t"}}{{$option}}{{end}}{{"\n"}}{{if .VisibleCategories}}{{"\n"}}` +
	style.Bold(`Commands`) + `{{range .VisibleCategories}}{{if .Name}}{{"\n"}}{{.Name}}:{{end}}{{end}}{{range .VisibleCommands}}{{"\n\t "}}` + color.White(` {{join .Names ", "}}`) + `{{"\t"}}{{.Usage}}{{end}}{{end}}{{"\n"}}`

var CommandHelpTemplate = color.White(`{{.Name}}`) + ` - {{.Usage}}{{"\n"}}` + color.White(`Usage`) +
	`{{"\n"}}{{if .UsageText}}{{.UsageText}}{{else}}{{.Name}}{{if .VisibleFlags}} [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}{{if .Category}}

` + color.Silver(`Category`) + `
   {{.Category}}{{end}}{{if .Description}}

` + color.Silver(`Description`) + `
   {{.Description}}{{end}}{{if .VisibleFlags}}

` + color.Silver(`Options`) + `
   {{range .VisibleFlags}}{{.}}{{end}}{{end}}
`

var SubcommandHelpTemplate = `Name
   ` + color.White(`{{.HelpName}}`) + ` - {{if .Description}}{{.Description}}{{else}}{{.Usage}}{{end}}

` + color.Silver(`Usage`) + `
   {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} command{{if .VisibleFlags}} [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}

` + color.Silver(`Commands`) + `{{range .VisibleCategories}}{{if .Name}}
   {{.Name}}:{{end}}{{range .VisibleCommands}}{{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{end}}{{if .VisibleFlags}}
` + color.Silver(`Options`) + `
   {{range .VisibleFlags}}{{.}}{{end}}{{end}}
`

func (self *CLI) ShowCLIHelp() {
	self.PrintBanner()
	self.PrintHelp()
	HelpPrinter(self.Writer, CLIHelpTemplate, self)
}

func (self *CLI) DefaultCLIComplete() {
	for _, command := range self.Commands {
		if command.Hidden {
			continue
		}
		for _, name := range command.Names() {
			fmt.Fprintln(self.Writer, name)
		}
	}
}

func (self *CLI) ShowCommandHelp(command string) error {
	if len(command) == 0 {
		HelpPrinter(self.Writer, SubcommandHelpTemplate, self)
		return nil
	}
	return nil
}

func (self *CLI) PrintVersion() {
	fmt.Fprintf(self.Writer, "%v version %v\n", self.Name, self.Version.String())
}

func printHelp(out io.Writer, templ string, data interface{}) {
	funcMap := template.FuncMap{
		"join": strings.Join,
	}
	w := tabwriter.NewWriter(out, 1, 8, 2, ' ', 0)
	t := template.Must(template.New("help").Funcs(funcMap).Parse(templ))
	err := t.Execute(w, data)
	if err != nil {
		return
	}
	w.Flush()
}
