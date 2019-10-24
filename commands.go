package cli

import (
	"strings"
)

type Command struct {
	Category    int
	Name        string
	Alias       string
	Description string
	Hidden      bool
	path        []string
	parent      *Command // TODO: Considier reducing this by replacing this dot.path command.subcommand
	Subcommands []Command
	Flags       []Flag
	Action      func(c *Context) error
}

func (self Command) is(name string) bool { return self.Name == name || self.Alias == name }
func (self Command) visible() bool       { return !self.Hidden }

func (self Command) usage() (output string) {
	if !IsBlank(self.Alias) {
		output += ", " + self.Alias
	}
	return self.Name + output
}

//
// Input Commands
///////////////////////////////////////////////////////////////////////////////
// TODO: This is the object we are giving to the resulting action defined in the
// command.
// TODO: Should this be all private?
type inputCommand struct {
	Name        string
	Flags       map[string]*inputFlag
	Parent      *inputCommand
	Subcommand  *inputCommand
	Subcommands []string // Should we just have a pointer to Command Definition and pull the names from there?
}

// TODO: This is WRONG, it MUST be the command object or at least not define the
// parent until we can confirm this is a child
func newInputCommand(name string) *inputCommand {
	return &inputCommand{
		Name:        name,
		Flags:       make(map[string]*inputFlag),
		Subcommands: []string{},
	}
}

// TODO: IF this is going to stay, it needs to be used in call cases
// []*inputFlag is used
func (self *inputCommand) addFlags(flags *inputFlags) {
	for _, flag := range *flags {
		self.Flags[flag.Name] = flag
	}
}

func (self *inputCommand) addCommandTree(commands []Command) {
	for _, command := range commands {
		inputCmd := &inputCommand{Name: command.Name, Parent: self}
		self.Subcommands = append(self.Subcommands, inputCmd)
		if IsGreaterThan(0, len(command.Subcommands)) {
			inputCmd.addCommandTree(command.Subcommands)
		}
	}
}

// InputCommands //////////////////////////////////////////////////////////////
type inputCommands []*inputCommand

func newCommandGroup() *inputCommands { return &inputCommands{} }

func (self *inputCommand) add(flag *inputCommand) { *self = append(*self, flag) }
func (self *inputCommand) reset()                 { self = &inputCommands{} }
func (self *inputCommand) isEmpty() bool          { return IsZero(len(*self)) }

//
// Public Methods
///////////////////////////////////////////////////////////////////////////////
func Commands(commands ...Command) []Command { return commands }

func (self *Command) Subcommand(name string) (Command, bool) {
	for _, subcommand := range self.Subcommands {
		if subcommand.Name == name || subcommand.Alias == name {
			return subcommand, true
		}
	}
	return Command{}, false
}

// NOTE: Must cascade through parents recursively if the flag is missing. If no
// commands in the path have the flag defined, then it is ignored.
func (self *Command) Flag(name string) (*Flag, bool) {
	name = strings.ToLower(name)
	for _, flag := range self.Flags {
		if flag.Name == name || flag.Alias == name {
			return &flag, true
		}
		//if self.ParentCommand
	}
	return nil, false
}

// TODO: We could try a radix tree that is loaded with the commands. Iterating
// through each row edge first, and assigning values using
// command1.command2.command3. Then we take our path and join with . and do a
// prefix search. We can try that later and see if the preformance gain is worth
// the extra overhead but this is not terrible, its technically a bread-first
// search
// NOTE: Public to allow essentially re-running the application without needing to
// start a new process
// TODO: Structurally this function still sucks
func (self *Command) Route(path []string) (*Command, bool) {
	if len(path) <= 1 && self.Name == path[0] {
		return self, true
	} else {
		command := self
		for _, name := range path[1:] {
			if subcommand, ok := command.Subcommand(name); ok {
				command = &subcommand
			} else {
				return nil, false
			}
		}
		return command, true
	}
	return nil, false
}
