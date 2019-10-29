package argument

//
// Input Commands
///////////////////////////////////////////////////////////////////////////////
type Command struct {
	Chain       *Chain
	Position    int
	Flags       map[string]Flag
	Parent      *Command
	Subcommands []string // Should we just have a pointer to Command Definition and pull the names from there?
	Arg         string
}

// TODO: This is WRONG, it MUST be the command object or at least not define the parent until we can confirm this is a child
func AddCommand(arg string) Command {
	return Command{
		Arg:         arg,
		Flags:       make(map[string]Flag),
		Subcommands: []string{},
	}
}

// TODO: IF this is going to stay, it needs to be used in call cases []*inputFlag is used
func (self Command) AddFlags(flags Flags) {
	// TODO: We should have the ability to do this task with a validation inline
	for _, flag := range flags {
		self.Flags[flag.Value()] = flag
	}
}

func (self Command) Path() []string {
	route := []string{self.Arg}
	for parent := self.Parent; parent != nil; parent = parent.Parent {
		route = append(route, parent.Arg)
	}
	return route
}

///////////////////////////////////////////////////////////////////////////////
func (self Command) Name() string { return self.Arg }

///////////////////////////////////////////////////////////////////////////////
func (self Command) NextArgument() (Argument, bool) { return self.Chain.NextArgument(self.Position) }

func (self Command) TrailingFlags() Flags { return self.Chain.TrailingFlags(self.Position) }
