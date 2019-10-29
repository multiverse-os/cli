package cli

import (
	"strings"

	argument "github.com/multiverse-os/cli/framework/argument"
	token "github.com/multiverse-os/cli/framework/argument/token"
	//data "github.com/multiverse-os/cli/framework/data"
)

// TODO: Input needs validation should probably move everything but input logic
// back into CLI or break context off input. then let input mingle with datatype
// and provide user input validation. we want to validate every input into the
// software including input from trusted users.

// TODO: Maybe we should just call this process, and also include the PID, and
// helpers for kill, signal handling, child process spawning, resource
// information, and on-the-fly daemonization. This would enable the receiver to
// do a lot with the resulting context including completely manage the process.
// TODO: Should be able to provide configPath based on CLI.Name, as well as
// cache folder.
// TODO: CLI framewwork ABSOLUTELY should be providing service management, daemonziation, generation of systemd/sysv init scripts, managing configurations, data folder and tempory folder.
type Context struct {
	PID           int
	CLI           *CLI
	CWD           string
	Executable    string
	Command       *argument.Command
	Flags         map[string]*argument.Flag
	Params        argument.Params
	ArgumentChain *argument.Chain
	Args          []string
}

func (self *Context) AddCommand(command *argument.Command) {
	self.CLI.Log(DEBUG, DebugInfo("Context.addCommand()"), "Add command func() on context with", VarInfo(command.Arg))
	self.CLI.Log(DEBUG, DebugInfo("Context.addCommand()"), VarInfo(self.Command.Arg))
	// Since this can't be done for both flag and command due to import looping, rather not do this with command if its possible
	//self.CLI.Debug(DebugInfo("Context.addCommand()"), VarInfo("Context.Command.definition.Subcommands", fmt.Sprintf("%s", self.Command.Definition.Subcommands)))

	//if !data.IsZero(len(self.Command.Definition.Subcommands)) {
	//	self.CLI.Log(DEBUG, DebugInfo("Context.addCommand()"), "Context.Subcommands length is not zero, so looping")
	//	//for _, subcommand := range self.Command.definition.Subcommands {
	//	//	self.CLI.Log(Debug, DebugInfo("Context.addCommand()"), VarInfo(subcommand.Arg), " == ", VarInfo(command.Arg))
	//	//	// TODO: What is this? Because it would need to check against alias too if this is a validation. Should just add validation to the model.
	//	//	//if subcommand.Arg == command.Arg {
	//	//	//	//self.Command = newInputCommand(command, subcommand.Arg)
	//	//	//	//self.CommandPath = append(self.CommandPath, self.Command.Arg)
	//	//	//}
	//	//}
	//}
}

// TODO: Still lacks support for declaring non-bool flags without "=", to
// acheive this we will need declaration of flags. Should just work this way if
// not declared, and if the developer declares then we support it
//func (self *Flag) MatchLongFlag(str string) bool {
//	flagComponents := strings.Split(str, "=")
//	switch {
//	case str[:1] == longFlag:
//		if []byte(str[1:(len(self.Name)+2)]) == []byte(self.Name) {
//			fmt.Println("str:", str[(len(self.Name)+2):])
//			self.Value = str[(len(self.Name) + 3):]
//		}
//	case str[0] == shortFlag:
//		parsedFlag := flagComponents[0][1:]
//		if 2 <= len(parsedFlag) {
//			for _, flag := range parsedFlag {
//
//			}
//		}
//	default:
//		return false
//	}
//}
func (self *Context) parseFlag(flag string) (*argument.Flag, bool) {
	//parsed := argument.Flag{}
	if strings.HasPrefix(flag, token.Long.String()) {
		// Long Flag - convention is enforcing '=' on Long val
		//flagParts := strings.Split(flag[:1], valueDelimeter)
		//if data.IsGreaterThan(1, len(flagParts)) {
		//	parsed.Name = flagParts[0]
		//	parsed.Value = flagParts[1]
		//} else {
		//	parsed.Name = flag[2:]
		//}
		//if flag, ok := self.CLI.command.Flag(parsed.Arg); ok {
		//	parsed.Name = flag.Name
		//	parsed.Type = flag.Type
		//	return parsed, true
		//}
	} else {
		// Short Flag (or Alias) (ex. ls -a)
		//for index, alias := range flag[1:] {
		//	// Stacked Short Flags (ex. `la -lah`)
		//	if flag, ok := self.CLI.command.Flag(string(alias)); ok {
		//		parsed.Arg = flag.Arg
		//		if data.IsGreaterThan(len(flag[1:]), index) {
		//			flagParts := strings.Split(string(alias), valueDelimeter)
		//			if data.IsGreaterThan(1, len(flagParts)) {
		//				parsed.Value = flagParts[1]
		//			}
		//		}
		//		return parsed, true
		//	}
		//}
	}
	return nil, false
}
