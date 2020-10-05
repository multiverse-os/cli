package cli

import (
	data "github.com/multiverse-os/cli/argument/data"
	token "github.com/multiverse-os/cli/argument/token"
)

// TODO: Be able to define the file extension that would be selected for when generating an autocomplete file
type Flag struct {
	Name        string
	Alias       string
	Description string
	Hidden      bool
	Default     string
	Type        data.Type
}

// TODO: Could probably speed up lookup and avoid this by putting flag in a
// lookup map twice, once with name and once with alias and just use a symbol
// TODO: Could probably be made private again since we had to move this back into the cli package for a sensible way of initializing and not requiring 5 imports
func (self Flag) is(name string) bool { return self.Name == name || self.Alias == name }

func (self Flag) usage() (output string) {
	output += token.Long.String() + self.Name
	if data.NotBlank(self.Alias) {
		output += ", " + token.Short.String() + self.Alias
	}
	return output
}

func Flags(flags ...Flag) []Flag { return flags }
