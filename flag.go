package cli

import (
	"strings"

	data "github.com/multiverse-os/cli/data"
	token "github.com/multiverse-os/cli/token"
)

// TODO: Be able to define the file extension that would be selected for when generating an autocomplete file
type Flag struct {
	Name        string
	Alias       string
	Description string
	Hidden      bool
	Default     string
	Value       string
	Type        data.Type
}

func HasFlagPrefix(flag string) (token.Identifier, bool) {
	if strings.HasPrefix(flag, token.Long.String()) &&
		data.IsGreaterThan(len(flag), token.Long.Length()) {
		return token.Long, true
	} else if strings.HasPrefix(flag, token.Short.String()) &&
		data.IsGreaterThan(len(flag), token.Short.Length()) {
		return token.Short, true
	} else {
		return token.NotAvailable, false
	}
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
