package cli

type ArgumentType int

const (
	FlagArgument ArgumentType = iota
	CommandArgument
)

type Argument interface {
	Names()
	Visible()
	Usage()
	Help()
}

type CLIAction int

// TODO: The idea here is to establish a router like model to make CLI
// application development more like web development, this will also let us
// easily do things like middleware which will speed up development, allow using
// libraries across a wider scope of applications and simplify the learning
// process
const (
	EmptyAction   CLIAction = iota
	CommandAction           // Command and possibly flags
	FlagAction              // Only flags specified
	VersionAction
	HelpAction
	HelpCommandAction // TODO: an this reasonably merge in with help the way we merged in subcommand?
)

// TODO: Decide if flags should be segregated into global flags and
// command flags
type Context struct {
	CLI      *CLI
	Command  Command
	Flags    map[string]interface{}
	Args     []string
	RawInput []string
}

///////////////////////////////////////////////////////////////////////////////

func (self *Context) hasCommand(name string) bool {
	return self.Command.Name == name
}

func (self *Context) hasCommandWithFlags(name string, flags ...string) bool {
	return self.hasCommand(name) &&
		self.hasFlags(flags...)
}

func (self *Context) hasVersionArgument() bool { return self.hasFlag("version") }

func (self *Context) hasHelpArgument() bool {
	// TODO: We don't need scope, we find that out by checking IF a command
	// exists. If not then the scope is

	// TODO:
	if self.Command.IsSubcommand() {

	} else if self.hasFlag("help") {
	} else {
		return false
	}
	// elsif self.Command.IsNotSubcommand (can only be here IF the command is
	// help, so can this condition even be met?
	// elsif self.Command.

	//return self.hasFlag("help")
}

func (self *Context) isEmpty(argType ArgumentType) bool {
	switch argType {
	case CommandArgument:
		return self.Command.Exists()
	case SubcommandArgument:
		return self.Subcommand.Exists()
	default:
		return false
	}
}

func (self *Context) hasNoArguments() bool { return (!self.Command.Exists() && Flags.Empty()) }
func (self *Context) hasNoFlags() bool     { return (len(self.Flags) == 0) }

func (self *Context) hasFlag(name string) bool {
	_, ok := self.Flags[name]
	return ok
}

func (self *Context) hasFlags(names ...string) bool {
	for _, name := range names {
		if !self.hasFlag(name) {
			return false
		}
	}
	return true
}

func (self *Context) hasCommand(name string) bool { return self.Command.is(name) }

////////////////////////////////////////////////////////////////////////////////
// TODO: I like this convience, but its only working with flags seems odd, it
// could be improved by allow naming of trailing arguments after a command so it
// could be used aswell. So a `./app open {Filename}` would get pulled up if
// flags turn up nothing. Since we have the FlagType (should be argument type
// turns out), we can just allow this type to be assigned during declaration and
// if that is defined, then it will check if the context helper lookup by type
// matches the type of the argument and wil return the argument if its not empty

func (self *Context) Filename(name string) string {
	flag, ok := self.Flags[name]
	if ok {
		return flag.Filename()
	} else {
		return "", errUnspecifiedFlag
	}
}

func (self *Context) Path(name string) string {
	flag, ok := self.Flags[name]
	if ok {
		return flag.Path()
	} else {
		return "", errUnspecifiedFlag
	}
}

func (self *Context) String(name string) string {
	flag, ok := self.Flags[name]
	if ok {
		return flag.String()
	} else {
		return "", errUnspecifiedFlag
	}
}

func (self *Context) Int(name string) int {
	flag, ok := self.Flags[name]
	if ok {
		return flag.Int()
	} else {
		return 0, errUnspecifiedFlag
	}
}

func (self *Context) Float(name string) float64 {
	flag, ok := self.Flags[nane]
	if ok {
		return flag.Float()
	} else {
		return float64(0.00), errUnspecifiedFlag
	}
}
