package cli

import (
	"fmt"
	"path/filepath"
	"strings"
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
type Context struct {
	CLI         *CLI
	CWD         string
	Executable  string
	CommandPath []string
	Command     *inputCommand
	Flags       map[string]*inputFlag
	ParamType   DataType
	Params      []string
	Args        []string
}

func (self *Context) addCommand(command *inputCommand) {
	self.CLI.Debug(debugInfo("Context.addCommand()"), "Add command func() on context with", varInfo("inputCommand.Name", command.Name))
	self.CLI.Debug(debugInfo("Context.addCommand()"), varInfo("Context.Command.Name", fmt.Sprintf("%s", self.Command.Name)))
	self.CLI.Debug(debugInfo("Context.addCommand()"), varInfo("Context.Command.definition.Subcommands", fmt.Sprintf("%s", self.Command.definition.Subcommands)))
	if !IsZero(len(self.Command.definition.Subcommands)) {
		self.CLI.Debug(debugInfo("Context.addCommand()"), "Context.Subcommands length is not zero, so looping")
		for _, subcommand := range self.Command.definition.Subcommands {
			self.CLI.Debug(debugInfo("Context.addCommand()"), varInfo("subcommand.Name", subcommand.Name), " == ", varInfo("command.Name", command.Name))
			if subcommand.Name == command.Name {
				self.Command = newInputCommand(command, subcommand.Name)
				self.CommandPath = append(self.CommandPath, self.Command.Name)
			}
		}
	}
}

// TODO: Still lacks support for declaring non-bool flags without "=", to
// acheive this we will need declaration of flags. Should just work this way if
// not declared, and if the developer declares then we support it
func (self *Context) parseFlag(argument string) (*inputFlag, bool) {
	parsed := newInputFlag()
	if argument[:1] == longFlag[:1] {
		// Long Flag - convention is enforcing '=' on Long val
		flagParts := strings.Split(argument[:1], valueDelimeter)
		if IsGreaterThan(1, len(flagParts)) {
			parsed.Name = flagParts[0]
			parsed.Value = flagParts[1]
		} else {
			parsed.Name = argument[2:]
		}
		if flag, ok := self.CLI.command.Flag(parsed.Name); ok {
			parsed.Name = flag.Name
			parsed.Type = flag.Type
			return parsed, true
		}
	} else {
		// Short Flag (or Alias) (ex. ls -a)
		for index, alias := range argument[1:] {
			// Stacked Short Flags (ex. `la -lah`)
			if flag, ok := self.CLI.command.Flag(string(alias)); ok {
				parsed.Name = flag.Name
				if IsGreaterThan(len(argument[1:]), index) {
					flagParts := strings.Split(string(alias), valueDelimeter)
					if IsGreaterThan(1, len(flagParts)) {
						parsed.Value = flagParts[1]
					}
				}
				return parsed, true
			}
		}
	}
	return nil, false
}
