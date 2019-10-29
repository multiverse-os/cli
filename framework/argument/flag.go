package argument

import (
	"strings"

	token "github.com/multiverse-os/cli/framework/argument/token"
	data "github.com/multiverse-os/cli/framework/data"
)

//
// Flag Input
///////////////////////////////////////////////////////////////////////////////
type Flag struct {
	Chain      *Chain
	Identifier token.Identifier
	Stacked    bool
	Position   int
	Type       data.Type
	Arg        string
}

func IsValidFlag(flag string) (token.Identifier, bool) {
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

func (self Flag) Name() string {
	flagParts := strings.Split(self.Value(), token.Short.String())
	return strings.Split(flagParts[len(flagParts)-1], token.Equal.String())[0]
}

func (self Flag) Value() string {
	if self.Separator(token.Equal) {
		return strings.Split(self.Arg, token.Equal.String())[0][1:]
	} else {
		// NOTE: Since the flag does not declare using an equal sign, we can assume
		// the next item is the value (the developer will know, and we will validate
		// it for any datatype. BUT if it is a valid flag, we are dealing with a Bool.
		nextValue, ok := self.NextArgument()
		if ok {
			if _, ok := IsValidFlag(nextValue.(Flag).Arg); ok {
				self.Type = data.Bool
				return data.BoolString(true)
			} else {
				if nextArgument, ok := self.NextArgument(); ok {
					return nextArgument.(Flag).Arg
				} else {
					return ""
				}
			}
		}
	}
	return data.Blank
}

func (self Flag) Separator(separatorToken token.Separator) bool {
	return strings.Contains(self.Arg, separatorToken.String())
}

///////////////////////////////////////////////////////////////////////////////
func (self Flag) NextArgument() (Argument, bool) {
	return self.Chain.NextArgument(self.Position)
}

// Flags //////////////////////////////////////////////////////////////////////
type Flags []Flag

func (self Flags) Add(flag Flag) { self = append(self, flag) }
func (self Flags) Reset()        { self = Flags{} }
func (self Flags) IsEmpty() bool { return data.IsZero(len(self)) }
