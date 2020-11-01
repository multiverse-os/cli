package cli

import (
	"strconv"
	"strings"

	data "github.com/multiverse-os/cli/data"
)

type Params struct {
	Position int
	Type     data.Type
	Value    []string
}

// TODO: Add ability to output URL, and Path types, since these would be very
// common and the ability to validate them would be nice. For example, being
// able to check if a file exists easily.

func (self Params) Strings() []string { return self.Value }

func (self Params) String() string { return strings.Join(self.Value, " ") }

func (self Params) Int() int {
	intValue, err := strconv.Atoi(self.Value[0])
	if err != nil {
		return 0
	} else {
		return intValue
	}
}

func (self Params) Bool() bool {
	for _, trueString := range data.TrueStrings {
		if trueString == self.Value[0] {
			return true
		}
	}
	return false
}
