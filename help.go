package cli

import (
	"fmt"
	"io"
	"strings"
	"text/tabwriter"
	"text/template"

	text "github.com/multiverse-os/cli-framework/text"
)

type helpPrinter func(w io.Writer, templ string, data interface{})

var HelpPrinter helpPrinter = printHelp

// TODO: Shouldnt this be in the RenderHelpText func ?
//if !self.HideHelp && checkHelp(context) {
//	ShowCLIHelp(context)
//	return nil
//}

//if !self.HideVersion && checkVersion(context) {
//	PrintVersion(context)
//	return nil
//}

// TODO: Lets not use global variables, it just doesnt feel right

// TODO: Why are commands VisibleCategories? This section is practically unreadable and very hard to customize
// TODO: All lines should be checked for length of 80 and broken into new line if so with the correct tab spacing prefixing it
// TODO: Use table library code to improve the structure of this and do better alignment of values
func (self *CLI) PrintHelp() {
	if self.Description != "" {
		fmt.Println(text.Strong(self.Description))
	}
	if self.Usage != "" {
		if self.NoANSIFormatting {
			fmt.Println("Usage")
		} else {
			fmt.Println(text.Strong("Usage"))
		}
		fmt.Print(text.Repeat(" ", 4))

		if self.NoANSIFormatting {
			fmt.Print(self.Name)
		} else {
			fmt.Print(text.Header(self.Name))
		}
		if self.HasVisibleFlags() {
			fmt.Print(" [options]")
		} else {
		}
		if self.HasVisibleCommands() {
			fmt.Print(" command [command options]")
		}
		if self.ArgsUsage != "" {
			fmt.Println(" " + self.ArgsUsage)
		} else {
			fmt.Println("[arguments...]")
		}
		if self.HasVisibleFlags() {
			fmt.Println("\n" + text.Strong("Options"))
			for _, flag := range self.VisibleFlags() {
				fmt.Println("flag: ", flag)
			}
		}
	}
}

var CLIHelpTemplate = `{{range $index, $option := .VisibleFlags}}{{if $index}}{{"\n"}}{{end}}{{"\t\t"}}{{$option}}{{end}}{{"\n"}}{{if .VisibleCategories}}{{"\n"}}` +
	fmt.Sprintf(text.STRONG) + `Commands` + fmt.Sprintf(text.RESET) + `{{range .VisibleCategories}}{{if .Name}}{{"\n"}}{{.Name}}:{{end}}{{end}}{{range .VisibleCommands}}{{"\n\t "}}` + fmt.Sprintf(text.H1) + ` {{join .Names ", "}}` + fmt.Sprintf(text.RESET) + `{{"\t"}}{{.Usage}}{{end}}{{end}}{{"\n"}}`

var CommandHelpTemplate = fmt.Sprintf(text.H1) + `{{.Name}}` + fmt.Sprintf(text.RESET) + ` - {{.Usage}}{{"\n"}}` + fmt.Sprintf(text.H1) + `Usage` + fmt.Sprintf(text.RESET) +
	`{{"\n"}}{{if .UsageText}}{{.UsageText}}{{else}}{{.Name}}{{if .VisibleFlags}} [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}{{if .Category}}

` + fmt.Sprintf(text.H2) + `Category` + fmt.Sprintf(text.RESET) + `
   {{.Category}}{{end}}{{if .Description}}

` + fmt.Sprintf(text.H2) + `Description` + fmt.Sprintf(text.RESET) + `
   {{.Description}}{{end}}{{if .VisibleFlags}}

` + fmt.Sprintf(text.H2) + `Options` + fmt.Sprintf(text.RESET) + `
   {{range .VisibleFlags}}{{.}}{{end}}{{end}}
`

var SubcommandHelpTemplate = `Name
   ` + fmt.Sprintf(text.H1) + `{{.HelpName}}` + fmt.Sprintf(text.RESET) + ` - {{if .Description}}{{.Description}}{{else}}{{.Usage}}{{end}}

` + fmt.Sprintf(text.H2) + `Usage` + fmt.Sprintf(text.RESET) + `
   {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} command{{if .VisibleFlags}} [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}

` + fmt.Sprintf(text.H2) + `Commands` + fmt.Sprintf(text.RESET) + `{{range .VisibleCategories}}{{if .Name}}
   {{.Name}}:{{end}}{{range .VisibleCommands}}{{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{end}}{{if .VisibleFlags}}
` + fmt.Sprintf(text.H2) + `Options` + fmt.Sprintf(text.RESET) + `
   {{range .VisibleFlags}}{{.}}{{end}}{{end}}
`

// TODO: There is literally no reason we need a special function for this
//func ShowCLIHelpAndExit(c *Context, exitCode int) {
//	ShowCLIHelp(c)
//	os.Exit(exitCode)
//}

// TODO: We need to move to a programmatic way instead of just chaining together
// strings and using linebreaks and spaces to do TUI visual. It would be much
// much better to do this via functions, then we can easily modify it, easily
// add if statements, etc.
func ShowCLIHelp(c *Context) {
	c.CLI.PrintBanner()
	c.CLI.PrintHelp()
	HelpPrinter(c.CLI.Writer, CLIHelpTemplate, c.CLI)
}

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

// TODO: There is no reason to have ShowCommandHelp and ShowCommandHelpAndExit
// as separate functions. This is prime example of unnecessary bloat in this
// codebase.
//func ShowCommandHelpAndExit(c *Context, command string, code int) {
//	ShowCommandHelp(c, command)
//	os.Exit(code)
//}

func ShowCommandHelp(ctx *Context, command string) error {
	if command == "" {
		HelpPrinter(ctx.CLI.Writer, SubcommandHelpTemplate, ctx.CLI)
		return nil
	}

	// TODO: Thhis is real gross, totally should not be doing this kind of stuff
	//for _, c := range ctx.CLI.Commands {
	//if c.HasName(command) {
	//	HelpPrinter(ctx.CLI.Writer, CommandHelpTemplate, c)
	//	return nil
	//}
	//}

	// TODO: No
	//if ctx.CLI.CommandNotFound == nil {
	//	return NewExitError(fmt.Sprintf("No help topic for '%v'", command), 3)
	//}

	//ctx.CLI.CommandNotFound(ctx, command)
	return nil
}

// ShowSubcommandHelp prints help for the given subcommand
func ShowSubcommandHelp(c *Context) error {
	return ShowCommandHelp(c, c.Command.Name)
}

func PrintVersion(c *Context) {
	fmt.Fprintf(c.CLI.Writer, "%v version %v\n", c.CLI.Name, c.CLI.Version.String())
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

// TODO: Why are we doing a check version? It should be present
//func checkVersion(c *Context) bool {
//	found := false
//	if VersionFlag.GetName() != "" {
//		eachName(VersionFlag.GetName(), func(name string) {
//			if c.GlobalBool(name) || c.Bool(name) {
//				found = true
//			}
//		})
//	}
//	return found
//}

// TODO: Again why are we checking if help exists? it should be present by
// default, there is almost noc ases we would not have help text for a
// command-line interface
//func checkHelp(c *Context) bool {
//	found := false
//	if HelpFlag.GetName() != "" {
//		eachName(HelpFlag.GetName(), func(name string) {
//			if c.GlobalBool(name) || c.Bool(name) {
//				found = true
//			}
//		})
//	}
//	return found
//}

// TODO: why?
//func checkCommandHelp(c *Context, name string) bool {
//	if c.Bool("h") || c.Bool("help") {
//		ShowCommandHelp(c, name)
//		return true
//	}
//	return false
//}

// TODO: Again why? Why not just store this flag info in the
// CLI and use it to run help. You should not need this.
// need it because you just ran help? well return nil isntead
// of checkiung this and then returning nil
//func checkSubcommandHelp(c *Context) bool {
//	if c.Bool("h") || c.Bool("help") {
//		ShowSubcommandHelp(c)
//		return true
//	}
//	return false
//}

// TODO: COmmand completion should be done with radix trees

//func checkShellCompleteFlag(a *CLI, arguments []string) (bool, []string) {
//	if !a.BashCompletion {
//		return false, arguments
//	}
//	pos := len(arguments) - 1
//	lastArg := arguments[pos]
//	if lastArg != "--"+BashCompletionFlag.GetName() {
//		return false, arguments
//	}
//	return true, arguments[:pos]
//}
//
//func checkCompletions(c *Context) bool {
//	if !c.shellComplete {
//		return false
//	}
//	if args := c.Args(); args.Present() {
//		name := args.First()
//		if cmd := c.CLI.Command(name); cmd != nil {
//			// let the command handle the completion
//			return false
//		}
//	}
//	ShowCompletions(c)
//	return true
//}
//
//func checkCommandCompletions(c *Context, name string) bool {
//	if !c.shellComplete {
//		return false
//	}
//
//	ShowCommandCompletions(c, name)
//	return true
//}
