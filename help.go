package cli

import (
	"fmt"
	"io"
	"os"
	"strings"
	"text/tabwriter"
	"text/template"

	text "github.com/multiverse-os/cli-framework/text"
	color "github.com/multiverse-os/cli-framework/text/color"
)

var CLIHelpTemplate = fmt.Sprintf(color.Subheader) + `{{.Name}} ` + fmt.Sprintf(color.Strong) + `v{{.Version}}
` + fmt.Sprintf(color.Reset) + text.Repeat("=", 80) + `{{if .Description}}
{{.Description}}
{{end}}
` + fmt.Sprintf(color.Strong) + `Usage:` + fmt.Sprintf(color.Reset) + `
   {{if .UsageText}}{{.UsageText}}{{else}}` + fmt.Sprintf(color.Subheader) + `{{.HelpName}} ` + fmt.Sprintf(color.Reset) + `{{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}
` + fmt.Sprintf(color.Strong) + `Global Options:` + fmt.Sprintf(color.Reset) + `
   {{range $index, $option := .VisibleFlags}}{{if $index}}
   {{end}}{{$option}}{{end}}
` + fmt.Sprintf(color.Strong) + `Commands:` + fmt.Sprintf(color.Reset) + `{{range .VisibleCategories}}{{if .Name}}
   {{.Name}}:{{end}}{{range .VisibleCommands}}
    ` + fmt.Sprintf(color.Subheader) + ` {{join .Names ", "}}` + fmt.Sprintf(color.Reset) + `{{"\t"}}{{.Usage}}{{end}}{{end}}
`

var CommandHelpTemplate = fmt.Sprintf(color.Subheader) + `{{.HelpName}}` + fmt.Sprintf(color.Reset) + ` - {{.Usage}}

` + fmt.Sprintf(color.Header) + `Usage:` + fmt.Sprintf(color.Reset) + `
   {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}}{{if .VisibleFlags}} [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}{{if .Category}}

` + fmt.Sprintf(color.Header) + `Category:` + fmt.Sprintf(color.Reset) + `
   {{.Category}}{{end}}{{if .Description}}

` + fmt.Sprintf(color.Header) + `Description:` + fmt.Sprintf(color.Reset) + `
   {{.Description}}{{end}}{{if .VisibleFlags}}

` + fmt.Sprintf(color.Header) + `Options:` + fmt.Sprintf(color.Reset) + `
   {{range .VisibleFlags}}{{.}}{{end}}{{end}}
`

var SubcommandHelpTemplate = `Name:
   ` + fmt.Sprintf(color.Subheader) + `{{.HelpName}}` + fmt.Sprintf(color.Reset) + ` - {{if .Description}}{{.Description}}{{else}}{{.Usage}}{{end}}

` + fmt.Sprintf(color.Header) + `Usage:` + fmt.Sprintf(color.Reset) + `
   {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} command{{if .VisibleFlags}} [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}

` + fmt.Sprintf(color.Header) + `Commands:` + fmt.Sprintf(color.Reset) + `{{range .VisibleCategories}}{{if .Name}}
   {{.Name}}:{{end}}{{range .VisibleCommands}}{{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{end}}{{if .VisibleFlags}}
` + fmt.Sprintf(color.Header) + `Options:` + fmt.Sprintf(color.Reset) + `
   {{range .VisibleFlags}}{{.}}{{end}}{{end}}
`

var helpCommand = Command{
	Name:      "help",
	Aliases:   []string{"h"},
	Usage:     "Shows a list of commands or help for one command",
	ArgsUsage: "[command]",
	Action: func(c *Context) error {
		args := c.Args()
		if args.Present() {
			return ShowCommandHelp(c, args.First())
		}

		ShowCLIHelp(c)
		return nil
	},
}

var helpSubcommand = Command{
	Name:      "help",
	Aliases:   []string{"h"},
	Usage:     "Shows a list of commands or help for one command",
	ArgsUsage: "[command]",
	Action: func(c *Context) error {
		args := c.Args()
		if args.Present() {
			return ShowCommandHelp(c, args.First())
		}
		return ShowSubcommandHelp(c)
	},
}

// Prints help for the CLI or Command
type helpPrinter func(w io.Writer, templ string, data interface{})

// HelpPrinter is a function that writes the help output. If not set a default
// is used. The function signature is:
// func(w io.Writer, templ string, data interface{})
var HelpPrinter helpPrinter = printHelp

// ShowCLIHelpAndExit - Prints the list of subcommands for the cli and exits with exit code.
func ShowCLIHelpAndExit(c *Context, exitCode int) {
	ShowCLIHelp(c)
	os.Exit(exitCode)
}

// ShowCLIHelp is an action that displays the help.
func ShowCLIHelp(c *Context) {
	HelpPrinter(c.CLI.Writer, CLIHelpTemplate, c.CLI)
}

// DefaultCLIComplete prints the list of subcommands as the default cli completion method
func DefaultCLIComplete(c *Context) {
	for _, command := range c.CLI.Commands {
		if command.Hidden {
			continue
		}
		for _, name := range command.Names() {
			fmt.Fprintln(c.CLI.Writer, name)
		}
	}
}

// ShowCommandHelpAndExit - exits with code after showing help
func ShowCommandHelpAndExit(c *Context, command string, code int) {
	ShowCommandHelp(c, command)
	os.Exit(code)
}

// ShowCommandHelp prints help for the given command
func ShowCommandHelp(ctx *Context, command string) error {
	// show the subcommand help for a command with subcommands
	if command == "" {
		HelpPrinter(ctx.CLI.Writer, SubcommandHelpTemplate, ctx.CLI)
		return nil
	}

	for _, c := range ctx.CLI.Commands {
		if c.HasName(command) {
			HelpPrinter(ctx.CLI.Writer, CommandHelpTemplate, c)
			return nil
		}
	}

	if ctx.CLI.CommandNotFound == nil {
		return NewExitError(fmt.Sprintf("No help topic for '%v'", command), 3)
	}

	ctx.CLI.CommandNotFound(ctx, command)
	return nil
}

// ShowSubcommandHelp prints help for the given subcommand
func ShowSubcommandHelp(c *Context) error {
	return ShowCommandHelp(c, c.Command.Name)
}

func PrintVersion(c *Context) {
	fmt.Fprintf(c.CLI.Writer, "%v version %v\n", c.CLI.Name, c.CLI.Version.String())
}

// ShowCompletions prints the lists of commands within a given context
func ShowCompletions(c *Context) {
	a := c.CLI
	if a != nil && a.BashComplete != nil {
		a.BashComplete(c)
	}
}

// ShowCommandCompletions prints the custom completions for a given command
func ShowCommandCompletions(ctx *Context, command string) {
	c := ctx.CLI.Command(command)
	if c != nil && c.BashComplete != nil {
		c.BashComplete(ctx)
	}
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

func checkVersion(c *Context) bool {
	found := false
	if VersionFlag.GetName() != "" {
		eachName(VersionFlag.GetName(), func(name string) {
			if c.GlobalBool(name) || c.Bool(name) {
				found = true
			}
		})
	}
	return found
}

func checkHelp(c *Context) bool {
	found := false
	if HelpFlag.GetName() != "" {
		eachName(HelpFlag.GetName(), func(name string) {
			if c.GlobalBool(name) || c.Bool(name) {
				found = true
			}
		})
	}
	return found
}

func checkCommandHelp(c *Context, name string) bool {
	if c.Bool("h") || c.Bool("help") {
		ShowCommandHelp(c, name)
		return true
	}
	return false
}

func checkSubcommandHelp(c *Context) bool {
	if c.Bool("h") || c.Bool("help") {
		ShowSubcommandHelp(c)
		return true
	}
	return false
}

func checkShellCompleteFlag(a *CLI, arguments []string) (bool, []string) {
	if !a.EnableBashCompletion {
		return false, arguments
	}
	pos := len(arguments) - 1
	lastArg := arguments[pos]
	if lastArg != "--"+BashCompletionFlag.GetName() {
		return false, arguments
	}
	return true, arguments[:pos]
}

func checkCompletions(c *Context) bool {
	if !c.shellComplete {
		return false
	}
	if args := c.Args(); args.Present() {
		name := args.First()
		if cmd := c.CLI.Command(name); cmd != nil {
			// let the command handle the completion
			return false
		}
	}
	ShowCompletions(c)
	return true
}

func checkCommandCompletions(c *Context, name string) bool {
	if !c.shellComplete {
		return false
	}

	ShowCommandCompletions(c, name)
	return true
}
