package cli

import (
	"fmt"
	"strings"
	"time"

	data "github.com/multiverse-os/cli/data"
	loading "github.com/multiverse-os/loading"
	squares "github.com/multiverse-os/loading/bars/squares"
	moon "github.com/multiverse-os/loading/spinners/moon"
)

///////////////////////////////////////////////////////////////////////////////
// Ontology of a command-line interface
///////////////////////////////////////////////////////////////////////////////
//
//            global flag    command flag             parameters (params)
//              __|___         __|__             __________|____________
//             /      \       /     \           /                       \
//     app-cli --flag=2 open -f thing template /path/to/file /path/to-file
//     \_____/          \__/           \____/
//        |              |                |
//   application       command        subcommand
//
///////////////////////////////////////////////////////////////////////////////
// Alpha Release

// TODO: Expand range of the tests so it test more possible conditions to
// guarantee it works when changes are made

// TODO: change receiver variable names on methods from self to the convention

// TODO: Rewrite the README.md

// TODO: Add ability to access Banner/Spinner (and others) text user interface
// (TUI) tools from actions.
//          context.CLI.Spinner()

// TODO: Ability to use ansii via CLI.Screen.Clear(), or CLI.Text.Blue("test")

type App struct {
	Name        string
	Description string
	Version     Version
	Debug       bool
	Outputs     Outputs
	GlobalFlags flags
	Commands    commands
	Actions     Actions
}

type CLI struct {
	Name        string
	Version     Version
	Build       Build
	Debug       bool
	Context     *Context
	Outputs     Outputs
	Actions     Actions
	MinimumArgs int    // TODO: Not yet implemented
	Locale      string // TODO: Not yet implented
}

func (c CLI) Log(output ...string)   { c.Outputs.Log(DEBUG, output...) }
func (c CLI) Warn(output ...string)  { c.Outputs.Log(WARN, output...) }
func (c CLI) Error(output ...string) { c.Outputs.Log(ERROR, output...) }
func (c CLI) Fatal(output ...string) { c.Outputs.Log(FATAL, output...) }

type loaderType int

const (
	Bar loaderType = iota
	Spinner
)

//func (self loaderType) String() string {
//  switch self {
//  case Spinner:
//    return "spinner"
//  case Bar:
//    return "bar"
//  default: // UndefinedLoaderType
//    return ""
//  }
//}
//
//func MarshalLoaderType(lType string) loaderType {
//  switch lType {
//  case Spinner.String():
//    return Spinner
//  default: // Bar
//    return Bar
//  }
//}

func (c CLI) LoadingBar() *loading.Bar {
	return loading.ToBar(c.Loader(Bar))
}

// TODO: It would be nice to be able to pass the animation to this spinner or
// loading bar via this and through ToSpinner()
func (c CLI) Spinner() *loading.Spinner {
	return loading.ToSpinner(c.Loader(Spinner))
}

func (c CLI) Loader(loader loaderType) loading.Loader {
	switch loader {
	case Spinner:
		return loading.NewSpinner(moon.Animation)
	case Bar:
		return loading.NewBar(squares.Animation).Length(80)
	default:
		return nil
	}
	//loadingBar.Status(color.Green("Completed!")).Complete()
}

// TODO: Submodule problem need to resolve to get this working, but tis
// advisable to eventually get this
//func (c CLI) Box(message string) string {
//  return text.Box(message)
//}

// TODO: Get rid of flag actions by simply catching version or help in a generic
// fallback that looks for these flags. This should also help resolve issues
// requiring hardcoding
func New(appDefinition ...App) (cli *CLI, errs []error) {
	// TODO: Clean this up so its not as ugly
	app := App{}
	if len(appDefinition) != 0 {
		app = appDefinition[0]
	}

	// Validation
	errs = append(errs, app.Commands.Validate()...)
	errs = append(errs, app.GlobalFlags.Validate()...)

	if len(errs) != 0 {
		fmt.Println("number of validation errors for flags and commands:", len(errs))
		return cli, errs
	}

	// NOTE: Sensical defaults to avoid error conditions, simplifying library use
	if data.IsBlank(app.Name) {
		app.Name = "app-cli"
	}
	if app.Version.undefined() {
		app.Version = Version{Major: 0, Minor: 1, Patch: 0}
	}
	if len(app.Outputs) == 0 {
		app.Outputs = append(app.Outputs, TerminalOutput())
	}

	// NOTE: If a fallback is not set, we render default help template.
	if app.Actions.Fallback == nil {
		app.Actions.Fallback = HelpCommand
	}

	cli = &CLI{
		Name:    app.Name,
		Version: app.Version,
		Outputs: app.Outputs,
		Actions: app.Actions,
		Build: Build{
			CompiledAt: time.Now(),
		},
	}

	// TODO: Why is Command, Flag
	// TODO: This is going to be troublesome come localization
	if !app.Commands.HasCommand("help") {
		app.Commands.Add(&Command{
			Name:        "help",
			Alias:       "h",
			Description: "outputs command and flag details",
			Action:      HelpCommand,
			Hidden:      true,
		})
	}

	if !app.Commands.HasCommand("version") {
		app.Commands.Add(&Command{
			Name:        "version",
			Alias:       "v",
			Description: "outputs version",
			Action:      RenderDefaultVersionTemplate,
			Hidden:      false,
		})
	}
	// NOTE: Application psuedo-command to store globals
	//       and simplify logic
	appCommand := Command{
		Name:        app.Name,
		Description: app.Description,
		Subcommands: app.Commands,
		Flags:       app.GlobalFlags,
		Hidden:      true,
		Action:      app.Actions.Fallback,
	}

	if !app.GlobalFlags.HasFlag("help") {
		hFlag := Flag{
			Command:     &appCommand,
			Name:        "help",
			Alias:       "h",
			Description: "outputs command and flag details",
			Hidden:      false,
			Action:      RenderDefaultHelpTemplate,
		}
		app.GlobalFlags.Add(hFlag)
	}

	if !app.GlobalFlags.HasFlag("version") {
		vFlag := Flag{
			Command:     &appCommand,
			Name:        "version",
			Alias:       "v",
			Description: "outputs version",
			Hidden:      true,
			Action:      RenderDefaultVersionTemplate,
		}
		app.GlobalFlags.Add(vFlag)
	}

	cli.Context = &Context{
		CLI:       cli,
		Process:   Process(),
		Commands:  Commands(appCommand),
		Params:    params{},
		Flags:     appCommand.Flags,
		Arguments: Arguments(appCommand),
		Actions:   actions{},
	}

	cli.Context.Command = cli.Context.Commands.First()

	return cli, errs
}

// TODO: We could use the BeforeAction hook to convert version and help flags
// into commands. Or rather convert trailing help commands into a flag? Don't
// fallback though on the concepts; just find better solutions

func (self *CLI) LastArgument() Argument {
	return self.Context.Arguments.Last()
}

func (self *CLI) FirstCommand() *Command {
	return self.Context.Commands.First()
}

func (self *CLI) LastCommand() *Command {
	return self.Context.Commands.Last()
}

func (self *CLI) IsLastArgumentCommand() bool {
	lastArgument := self.LastArgument()
	fmt.Printf("lastArgument(%v)\n", lastArgument)

	return false
}

func (self *CLI) Parse(arguments []string) *CLI {
	defer self.benchmark(time.Now(), "benmarking argument parsing")

	// NOTE
	// Skip one because we treat the application a command so it
	// can store the global flags. This model avoids a lot of
	// extra code
	for index, argument := range arguments[1:] {
		// Flag parse
		// But shouldn't flag parsing go from each command upwards?
		// Flag
		if flagType, ok := HasFlagPrefix(argument); ok {
			argument = flagType.TrimPrefix(argument)
			fmt.Printf("argument(%v)\n", argument)
			switch flagType {
			case Short:
				for index, shortFlag := range argument {
					// NOTE: Confirm we are not last && next argument is '=' (61) &&
					if len(argument) != index+1 && argument[index+1] == 61 {
						if flag := self.Context.Flag(string(shortFlag)); flag != nil {
							if flagParam := argument[index+2:]; len(flagParam) != 0 {
								flag.Set(flagParam)
							}
							self.Context.Arguments = self.Context.Arguments.Add(flag)
							break
						}
					} else {
						if flag := self.Context.Flag(string(shortFlag)); flag != nil {
							// NOTE: If the default value is not boolean or blank, no
							// assignment occurs to avoid input failures.

							//
							if data.IsBoolean(flag.Default) {
								flag.Toggle()
							} else if len(flag.Default) == 0 {
								flag.SetTrue()
							}

							self.Context.Arguments = self.Context.Arguments.Add(flag)
						}
					}
				}
			case Long:
				longFlagParts := strings.Split(argument, "=")
				fmt.Printf(
					"longFlagParts[0](%v) + len(%v)\n",
					longFlagParts[0],
					len(longFlagParts),
				)
				flag := self.Context.Flag(longFlagParts[0])
				//// TODO: Validate (which probably should be setting default)
				if flag != nil {
					switch len(longFlagParts) {
					case 2:
						flag.Set(longFlagParts[1])
					case 1:
						// NOTE
						// If we only use default to determine type we are ignoring a few edge
						// conditions that will inevitably make this difficult to use
						//   An Boolean: true type attribute would solve it
						// HasNext() check
						// TODO: +2 because index starts at 0 len() starts at 1
						if data.IsBoolean(flag.Default) {
							// NOTE IS BOOLEAN
							flag.Toggle()
						} else if index+2 <= len(arguments[1:]) {
							// NOTE HAS NEXT
							// NOTE
							// We don't need to know if it is a valid flag, just that it is a
							// flag
							if _, ok := HasFlagPrefix(arguments[index+2]); ok {
								fmt.Printf("Next argument is flag\n")
								flag.SetTrue()
							}

							if self.LastCommand().Subcommands.HasCommand(arguments[index+2]) {
								fmt.Printf("Next Argument is subcommand of previous command\n")
								// NOTE
								// In this condition we are talking about a boolean
								flag.SetTrue()
							}

							flag.Set(arguments[index+2])
						} else {
							// NOTE
							// NOT boolean && No Next
							// Should be using default or setting to false
							// else here should be assumed boolean
							flag.SetTrue()
						}

					}
					fmt.Printf("flag(%v) before adding...\n", flag)
					self.Context.Arguments.Add(flag)
				} else {
					// TODO: There are conditions that land here
					//       ones with default that not boolean
					//       and has two parts
					//       and doesnt have next argument
				}
			}

			// TODO: Ummm i dont like how this could trigger even if above is
			//       THIS HAS TO ONLY BE CHECKED NOT ON FIRST COMMAND
			//       THIS IS FOR SUBCOMMANDS ONLY!
			//       Commands need to be size of 2
			//fmt.Printf("len(self.Context.Commands)(%v)\n", len(self.Context.Commands))
			//fmt.Printf("default, checking for 'help'...\n")
			//if (len(argument) == 4 && argument == "help") ||
			//	(len(argument) == 1 && argument == "h") {
			//	if 2 <= len(self.Context.Commands) {
			//		helpFlag := self.LastCommand().Flag("help")
			//		if helpFlag == nil {
			//			lastCommand := self.LastCommand()
			//			fmt.Printf("lastCommand(%v)\n", lastCommand)

			//		}
			//	} else {
			//		// TODO: Trigger
			//	}
			//}
			// TODO: SO we are not catching HELP flags for each subcommand
			//       if
		} else {
			if command := self.Context.Command.Subcommand(argument); command != nil {
				// Command parse
				command.Parent = self.FirstCommand()

				self.Context.Commands.Add(command)
				// TODO: don't we do add here? otherwise what was the poiint?
				self.Context.Flags = append(self.Context.Flags, command.Flags...)

				self.Context.Arguments = self.Context.Arguments.Add(self.FirstCommand())

				self.Context.Command = self.FirstCommand()

				// TODO: Should not be
				//} else if (len(argument) == 4 && argument == "help") ||
				//	(len(argument) == 1 && argument == "h") {
				// TODO: Because using help on a subcommand doesnt parse because help is
				// global. And thats how it should work. Version doesn't need this.
				// But I really hate this hardcoding
				//helpCommand := self.LastCommand().Subcommand("help")
				//if helpCommand != nil {
				//	// TODO: Why is this the parent? What if we are dealing with
				//	//       a subcommand of a subcommand? we would want the subcommand
				//	//       not the first command
				//	helpCommand.Parent = self.FirstCommand()

				//	self.Context.Commands.Add(*helpCommand)
				//	self.Context.Flags = append(self.Context.Flags, helpCommand.Flags...)

				//	self.Context.Arguments = self.Context.Arguments.Add(
				//		self.FirstCommand(),
				//	)

				//	// TODO: Wait wtf whattt we are setting parent to the same thing
				//	//       as we are setting the fucking command this makes no fucking
				//	//       sense
				//	self.Context.Command = self.FirstCommand()
				//	break
				//}
			} else {
				// Params parse
				fmt.Printf("params(%v)\n", argument)
				// TODO: SO THIS IS WHERE OUR ISSUE IS RIGHT NOW
				//       THIS PARAM MUST BE SET TO LAST FLAG!

				//flag := self.Context.Arguments.PreviousIfFlag()
				//// THIS RETURNS WRONG TYPE OF OBJECT, WE EXPECT
				//// what we had above
				//fmt.Printf("previous if flag: (%v)\n", flag)
				//if flag != nil {
				//	if flag.Param.value == flag.Default {
				//		flag.Param = NewParam(argument)
				//	} else {
				//		flag = nil
				//	}
				//}
				//if flag == nil {
				self.Context.Params = self.Context.Params.Add(
					NewParam(argument),
				)
				self.Context.Arguments = self.Context.Arguments.Add(
					self.Context.Params.First(),
				)
			}
		}
	}
	// End of parse

	// for the purpose of making it easier to use
	// to access in this function in the reverse order.
	//self.Context.Arguments = Reverse(self.Context.Arguments)
	//self.Context.Commands = ToCommands(Reverse(self.Context.Commands.Arguments()))
	//self.Context.Params = ToParams(Reverse(self.Context.Params.Arguments()))

	//return self

	fmt.Printf("self.Context.Arguments ... len(%v)\n", len(self.Context.Arguments))
	return self
}

func (cli *CLI) Execute() {
	cli.Context.Actions.Add(cli.Actions.OnStart)

	// TODO: No what should happen is we go through the global flags and run those
	// actions first
	// TODO: This all fell apart when we had to hard-code 'help' flag AND
	// requiring a skip action assignment for commands
	var skipCommandAction bool
	for _, command := range cli.Context.Commands {
		for _, flag := range command.Flags {
			if flag.Action != nil && flag.Param != nil && data.IsTrue(flag.Param.value) {
				cli.Context.Actions = append(cli.Context.Actions, flag.Action)
				skipCommandAction = true
				break
			}
		}
	}

	if !skipCommandAction {
		if 0 < len(cli.Context.Commands) {
			command := cli.Context.Commands.First()
			if command.Action != nil {
				cli.Context.Actions = append(cli.Context.Actions, command.Action)
			}
		}
	}

	cli.Context.Actions.Add(cli.Actions.OnExit)

	// NOTE: Before handing the developer using the library the context we put
	// them in the expected left to right order, despite it being easier for us
	defer cli.benchmark(time.Now(), "benmarking action execution")
	for _, action := range cli.Context.Actions {
		action(cli.Context)
	}
}
