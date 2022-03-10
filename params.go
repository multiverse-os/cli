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

// NOTE: Length Limit in Linux
//         Path:     4096
//         Filename: 256
// TODO: *(?) Base the length on the datatype?*
func ValidateParam(param Param) error {
  if len(param.Value) < 4096 {
    return errInvalidArgumentLength
  }
  // TODO: Format validation should be based on data type
  //for _, paramRune := range param.Value {
  //  if !unicode.IsLetter(paramRune) {
  //    return errors.New(errInvalidParamFormat)
  //  }
  //}
  return nil
}

func (self Param) IsValid() bool { return ValidateParam(self) != nil }

// TODO: These should more heavily rely on existing code in data subpackage
func (self Param) String() string { return self.Value }
func (self Param) Bool() bool { return data.IsTrue(self.Value) }

func (self Param) Int() int {
  intValue, err := strconv.Atoi(self.Value[0:1])
	if err != nil {
		return 0
	} else {
		return intValue
	}
}

// TODO: Float 

// TODO: Path / Filename

// TODO: URL

// TODO: This all should be handled by data subpackage

///////////////////////////////////////////////////////////////////////////////
type params []*Param

func Params(params ...Param) (paramPointers params) { 
  for index, _ := range params {
    paramPointers = append(paramPointers, &params[index])
  }
  return paramPointers
}

func (self params) Arguments() (arguments arguments) {
  for _, param := range self {
    arguments = append(arguments, Argument(param))
  }
  return arguments
}

func (self params) Add(paramValue string) (params params) {
  //err := ValidateParam(*newParam)
  return append(append(params, &Param{Value: paramValue}), self...)
}

// TODO: Add ability to output URL, and Path types, since these would be very
//       common and the ability to validate them would be nice. For example,
//       being able to check if a file exists easily.

// TODO: Once the params have been loaded, begin loading params again; then
//       apply

func (self params) Count() int { return len(self) }

func (self params) First() *Param {
  if 0 < self.Count() {
    return self[0]
  }
  return nil
}

func (self params) Last() *Param { 
  if 0 < self.Count() {
    return self[self.Count()-1] 
  }
  return nil
}

func (self params) Reverse() (reversedParams params) {
  for index := self.Count() - 1; index >= 0; index-- {
    reversedParams = append(reversedParams, self[index])
  }
  return reversedParams
}

func (self params) IsZero() bool { return self.Count() == 0 }
func (self params) String() string { return strings.Join(self.Strings(), " ") }

func (self params) Strings() (paramStrings []string) { 
  for _, param := range self {
    paramStrings = append(paramStrings, param.Value)
  }
  return paramStrings
}
