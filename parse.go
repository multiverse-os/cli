package cli

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"
)

func (self *CLI) Parse(arguments []string) (*Context, error) {
	defer self.benchmark(time.Now(), "benmarking argument parsing and action execution")
	cwd, executable := filepath.Split(arguments[0])

	context := &Context{
		CLI:          self,
		CWD:          cwd,
		Command:      &self.Command,
		Executable:   executable,
		CommandChain: &Chain{},
		Params:       Params{},
		Flags:        make(map[string]*Flag),
		Args:         arguments[1:],
	}
	context.CommandChain.AddCommand(&self.Command)

	parsedFlags := []Flag{}
	for index, argument := range context.Args {
		if flagType, ok := HasFlagPrefix(argument); ok {
			fmt.Println("index: [%v] \n", index)
			fmt.Println("argument:", argument)

			// TODO: Need to handle skipping next argument when next argument is used
			parsedFlags = append(parsedFlags, context.ParseFlag(flagType, argument, context.NextArgument(index)))

			//context.ParseFlag(index, flagType, &Flag{Name: argument})
		} else {
			if command, ok := context.Command.Subcommand(argument); ok {
				command.Parent = context.Command
				context.Command = &command
				context.CommandChain.AddCommand(context.Command)
			} else {
				for _, param := range context.Args[index:] {
					context.Params.Value = append(context.Params.Value, param)
				}
				break
			}
		}
	}

	// NOTE: Updating them in a batch as such will serve to avoid wasting resources.
	context.UpdateFlags(parsedFlags)

	if context.CommandChain.UnselectedCommand() {
		context.Command = &Command{
			Parent: context.Command,
			Name:   "help",
		}
	}

	self.Debug = context.HasFlag("debug")
	if context.Command.is("version") || context.HasGlobalFlag("version") {
		self.RenderVersionTemplate()
	} else if context.Command.is("help") || context.HasFlag("help") {
		context.RenderHelpTemplate()
	} else {
		context.Command.Action(context)
	}
	return context, nil
}

// TODO: MISSING ABILITY TO PARSE FLAGS THAT ARE USING "QUOTES TO SPACE TEXT".
// TODO: MISSING Flags of slice types can be passed multiple times (-f one -f two -f three)
// TODO: MISSING Collect ALL arguments trailing `--`
// TODO: MISSING ability to stack flag names of any size (right now assumes only
//       1 character size is allowed for short command names).
func (self *Context) ParseFlagNameAndValue(argument, nextArgument string) (string, string) {
	flagParts := strings.Split(StripFlagPrefix(argument), "=")

	if len(flagParts) == 2 {
		return strings.ToLower(flagParts[0]), flagParts[1]
	} else {
		// NOTE: Check if nextArgument is flag, flag is a boolean if nextArgument is
		//       either a flag or is a known command.
		if _, ok := HasFlagPrefix(nextArgument); ok {
			return strings.ToLower(flagParts[0]), "1"
		} else {
			for _, subcommand := range self.CommandChain.Reversed() {
				if subcommand.is(nextArgument) {
					return strings.ToLower(flagParts[0]), "1"
				}
			}
		}
	}
	return strings.ToLower(flagParts[0]), nextArgument
}

func (self *Context) ParseFlag(flagType FlagType, argument, nextArgument string) (parsedFlag Flag) {
	argument = strings.ToLower(argument)
	// NOTE: Next argument may be value for flag so it may not be lowercased by
	//       default.

	flagParts := strings.Split(StripFlagPrefix(argument), "=")
	fmt.Println("flagParts[0]:", flagParts[0])
	fmt.Println("len(flagParts[0]):", len(flagParts[0]))
	fmt.Println("1 <= len(flagParts[0]):", 1 <= len(flagParts[0]))

	//  TODO: We ONLY check for short to see if we have stacked flags.
	if flagType == Short {
	}
	// NOTE: Before attempting to parse as stacked short flags, attempt to parse
	//       as typo of a long flag.

	for _, command := range self.CommandChain.Reversed() {
		for _, flag := range command.Flags {
			// NOTE: With A_FLAG and NAME, and VALUE, drop out flag with
			if flag.is(flagParts[0]) {
				if len(flagParts[0]) == 2 {
					// NOTE: Two means that the value is already included, divided by
					//       an `=` sign.
					flag.Value = flagParts[1]
					return flag
				} else {
					//

				}

			}
		}
	}

	// TODO: Handle stacking short flag
	if len(flagParts[0]) != 1 {
		// NOTE: If a short tag is longer than 1 character
		for index, stackedFlag := range flagParts[0] {
			fmt.Println("stackedFlag:", stackedFlag)
			if index == len(flagParts[0]) {

			}
		}
		// TODO: Last item in stacked flags could be not boolean, check next
		// argumetn to decide

		flagParts = strings.Split(argument[2:len(argument)], "=")

		fmt.Println("flagParts:", flagParts)
		fmt.Println("len(flagParts):", len(flagParts))

		parsedFlag.Value = flagParts[1]
		parsedFlag.Name = flagParts[0]
		if _, ok := HasFlagPrefix(nextArgument); ok {
			// NOTE: Next argument is flag, so flag.Value is 1 (boolean)
			parsedFlag.Value = "1"
		}
		for _, subcommand := range self.Command.Subcommands {
			if nextArgument == subcommand.Name {
				// NOTE: Next argument is a command, so flag.Value is 1 (boolean)
				parsedFlag.Value = "1"
			}
		}

		if len(parsedFlag.Value) == 0 {
			fmt.Println("assigning nextArgument to parsedFlag.Value:", nextArgument)
			parsedFlag.Value = nextArgument
		}

		fmt.Println("parsedFlag.Name:", parsedFlag.Name)
		fmt.Println("parsedFlag.Value:", parsedFlag.Value)

	}
	return parsedFlag
}

func (self *Context) UpdateFlags(parsedFlags []Flag) {
	for _, parsedFlag := range parsedFlags {
		for _, command := range self.CommandChain.Reversed() {
			var flags []Flag
			for _, flag := range command.Flags {
				if flag.is(parsedFlag.Name) {
					fmt.Println("[ASSINGING PARSED FLAG] flag.is(parsedFlag.Name):", parsedFlag.Value)
					fmt.Println("assinging parsedFlag to flag.Value:", parsedFlag.Value)

					flag.Value = parsedFlag.Value
				}
				flags = append(flags, flag)
			}
			command.Flags = flags
		}
	}

}

// NOTE: These are here for dev reasons while parsing is being completed; once
// it is these can be moved into the appropriate files like flag.go
func StripFlagPrefix(flagName string) string { return strings.Replace(flagName, "-", "", -1) }

func FlagNameForType(flagType FlagType, argument string) (name string) {
	switch flagType {
	case Short:
		name = argument[1:len(argument)]
	case Long:
		name = argument[2:len(argument)]
	}
	return strings.ToLower(strings.Split(name, "=")[0])
}

func (self *Context) NextArgument(index int) string {
	if index+2 < len(self.Args) {
		return self.Args[index+1]
	}
	return ""
}
