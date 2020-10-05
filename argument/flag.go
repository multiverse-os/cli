package argument

import (
	"errors"
	"os"
	"strconv"
	"strings"

	data "github.com/multiverse-os/cli/argument/data"
	token "github.com/multiverse-os/cli/argument/token"
)

// TODO: Using Position seems like a bad idea, especially if we want to be able to insert arguments placed in the wrong spot like in the params to be more intuitive
// we may want to just provide an ID then do a scan throught he chain for the ID
//
// Flag Input
///////////////////////////////////////////////////////////////////////////////
type Flag struct {
	Identifier token.Identifier
	Type       data.Type
	Stacked    bool
	Name       string
	Value      string
	Arg        string
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

func (self Flag) Valid(dataType data.Type) (bool, error) {
	switch dataType {
	case data.Bool:
		boolStrings := append(data.TrueStrings, data.FalseStrings...)
		for _, boolValue := range boolStrings {
			if boolValue == self.Value {
				return true, nil
			}
		}
		return false, errors.New("[error] could not parse valid boolean value")
	//case Int:
	case data.String:
		return true, nil
	//case Directory:
	case data.Filename:
		_, err := os.Stat(self.Value)
		return (err == nil), nil
	//case Filenames:
	//case URL:
	//case IPv4:
	//case IPv6:
	//case Port:
	default:
		return false, errors.New("[error] failed to parse data type")
	}
}

func (self Flag) String() string { return self.Value }

func (self Flag) Int() int {
	intValue, err := strconv.Atoi(self.Value)
	if err != nil {
		return 0
	} else {
		return intValue
	}
}

func (self Flag) Bool() bool {
	for _, trueString := range data.TrueStrings {
		if trueString == self.Value {
			return true
		}
	}
	return false
}
