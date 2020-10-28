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
	Parent      *Command // TODO: Considier reducing this by replacing this dot.path command.subcommand
	Subcommands []Command
	Flags       []Flag
	Action      Action
}

// TODO Consider doing murmur has for comparisons, to speed up everything.

// TODO: Be able to walk over the command tree to output all items, for putting into a map, for hashing and creating ids, etc
// this was already written just look a few commits back

// TODO: Support printing command tree
// TODO: These may be able to go back to being private, were made public when we experimented with having their own package
func (self Command) is(name string) bool { return self.Name == name || self.Alias == name }

func (self Command) visibleSubcommands() (subcommands []Command) {
	for _, subcommand := range self.Subcommands {
		if !subcommand.Hidden {
			subcommands = append(subcommands, subcommand)
		}
	}
	return subcommands
}

func (self Command) visibleFlags() (flags []Flag) {
	for _, flag := range self.Flags {
		if !flag.Hidden {
			flags = append(flags, flag)
		}
	}
	return flags
}

func (self Command) usage() (output string) {
	if len(self.Alias) != 0 {
		output += ", " + self.Alias
	}
	return self.Name + output
}

func (self Command) path() []string {
	route := []string{self.Name}
	for parent := self.Parent; parent != nil; parent = parent.Parent {
		route = append(route, parent.Name)
	}
	return route
}

//
// Public Methods
///////////////////////////////////////////////////////////////////////////////
func Commands(commands ...Command) []Command { return commands }

func (self Command) Subcommand(name string) (Command, bool) {
	for _, subcommand := range self.Subcommands {
		if subcommand.is(name) {
			return subcommand, true
		}
	}
	return Command{}, false
}

// NOTE: Must cascade through parents recursively if the flag is missing. If no
// commands in the path have the flag defined, then it is ignored.
func (self Command) Flag(arg string) (*Command, *Flag, bool) {
	arg = strings.ToLower(arg)
	for _, flag := range self.Flags {
		if flag.Name == arg || flag.Alias == arg {
			return &self, &flag, true
		}
		//if self.ParentCommand
	}
	return nil, nil, false
}

func (self Command) Path() []string {
	route := []string{self.Name}
	for parent := self.Parent; parent != nil; parent = parent.Parent {
		route = append(route, parent.Name)
	}
	return route
}

// TODO: We could try a radix tree that is loaded with the commands. Iterating
// through each row edge first, and assigning values using
// command1.command2.command3. Then we take our path and join with . and do a
// prefix search. We can try that later and see if the preformance gain is worth
// the extra overhead but this is not terrible, its technically a bread-first
// search
// NOTE: Public to allow essentially re-running the application without needing to
// start a new process
// TODO: THIS IS THE SLOWEST FUNCTION, THIS IS OUR BOTTLENECK
func (self Command) Route(path []string) (Command, bool) {
	if len(path) == 1 && self.Name == path[0] {
		return self, true
	} else {
		command := self
		for _, name := range path[1:] {
			if subcommand, ok := command.Subcommand(name); ok {
				command = subcommand
			} else {
				return Command{}, false
			}
		}
		return command, true
	}
	return Command{}, false
}
