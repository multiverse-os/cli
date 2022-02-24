package cli

import (
	"strconv"
	"strings"

	data "github.com/multiverse-os/cli/data"
)

type Param struct {
	DataType data.Type
	Value    string
}

type params []*Param

func (self *Param) Type() ArgumentType { return ParamArgument }

func Params(params ...Param) (paramPointers params) { 
  for _, param := range params {
    paramPointers = append(paramPointers, &param)
  }
  return paramPointers
}

// TODO: Add ability to output URL, and Path types, since these would be very
//       common and the ability to validate them would be nice. For example,
//       being able to check if a file exists easily.

// TODO: Once the params have been loaded, begin loading flags again; then
//       apply

func (self params) Count() int { return len(self) }
func (self params) Last() *Param { return self[self.Count()-1] }
func (self params) IsZero() bool { return self.Count() == 0 }

func (self params) Strings() (paramStrings []string) { 
  for _, param := range self {
    paramStrings = append(paramStrings, param.Value)
  }
  return paramStrings
}

func (self params) String() string { 
  return strings.Join(self.Strings(), " ")  
}

// TODO: Param.String() 
func (self Param) String() string {
  return self.Value
}

func (self Param) Int() int {
	intValue, err := strconv.Atoi(self.Value[0])
	if err != nil {
		return 0
	} else {
		return intValue
	}
}

func (self Param) Bool() bool {
	for _, trueString := range data.True.Strings() {
		if trueString == self.Value[0] {
			return true
		}
	}
	return false
}

// TODO: Float 

// TODO: Path

// TODO: URL

// TODO: 
