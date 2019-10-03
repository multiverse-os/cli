package cli

import (
	"strings"
)

func (self *CLI) parse(arguments []string) *Context {
	var skipArgument bool

	context := &Context{
		CLI:        self,
		Flags:      make(map[string]Flag),
		Command:    Command{},
		Subcommand: Command{},
		Args:       []string{},
	}

	// NOTE: **************************************************
	//         This parse method is temporary, a router based
	//         on a static double trie for at least handling
	//         the command, subcommand, and possibly 3 levels
	//         of commands. This will also provide a built in
	//         completer without needing to build special file
	//         for each shell.
	//        [shlexer] also this shell lexer needs to be merged
	//         in and reviewed against the mattn one to see if
	//         we can improve our rough draft
	//       **************************************************

	for index, argument := range arguments {
		if skipNextArgument {
			skipNextArgument = false
			continue
		}
		if string(argument[0]) == "-" || len(argument) > 2 && argument[:2] == "--" {
			argument = strings.ReplaceAll(argument, "-", "")
			var flagName string
			var flagValue interface{}
			if strings.Contains(argument, "=") {
				flagParts := strings.Split(argument, "=")
				flagName = flagParts[0]
				flagValue = flagParts[1]
			} else {
				skipNextArgument = true
				flagName = argument
				if len(arguments) > (index + 1) {
					flagValue = arguments[index+1]
				} else {
					flagValue = true
				}
			}
			if context.Command.NotEmpty() {
				ok, flag := self.isCommandFlag(context.Command, flagName)
				if ok {
					flag.Value = flagValue
					context.Flags[flag.Name] = flag
				}
			} else {
				ok, flag := self.isFlag(flagName)
				if ok {
					flag.Value = flagValue
					context.Flags[flag.Name] = flag
				}
			}
		} else {
			if context.Command.Empty() {
				ok, command := self.isCommand(argument)
				if ok {
					context.Command = command
					if len(arguments) > index+1 {
						ok, Subcommand := self.isSubcommand(context.Command, arguments[index+1])
						if ok {
							context.Subcommand = Subcommand
						} else {
							context.Args = arguments[index+1:]
						}
					}
					return context
				}
			}
		}
	}
	return context
}
