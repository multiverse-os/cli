package cli

import (
  "strconv"
  "strings"

  data "github.com/multiverse-os/cli/data"
)

// TODO: Create a New() function so we can hook in validations
type Param struct {
  DataType data.Type
  value    string
}

func NewParam(argument string) *Param {
  return &Param{
    value: argument,
  }
}

// NOTE: Length Limit in Linux
//         Path:     4096
//         Filename: 256
func ValidateParam(param Param) error {
  if len(param.value) < 4096 {
    return ErrInvalidArgumentLength
  }
  // TODO: Format validation should be based on data type
  //       Yes, it should be based on type switch
  //for _, paramRune := range param.value {
  //  if !unicode.IsLetter(paramRune) {
  //    return errors.New(ErrInvalidParamFormat)
  //  }
  //}
  return nil
}

func (self Param) IsValid() bool { return ValidateParam(self) != nil }

// TODO: These should more heavily rely on existing code in data subpackage
func (self Param) Value() string { return self.value }
func (self Param) String() string { return self.value }
func (self Param) Bool() bool { return data.IsTrue(self.value) }

func (self Param) Int() int {
  intValue, err := strconv.Atoi(self.value[0:1])
  if err != nil {
    return 0
  } else {
    return intValue
  }
}

// TODO: Float 

// TODO: Path / Filename

// TODO: URL

///////////////////////////////////////////////////////////////////////////////
type params []*Param

func Params(params ...Param) (paramPointers params) { 
  for index, _ := range params {
    paramPointers = append(paramPointers, &params[index])
  }
  return paramPointers
}

func (self params) Arguments() (newArguments arguments) {
  for _, param := range self {
    newArguments = append(newArguments, param)
  }
  return newArguments
}

func (self params) Add(param *Param) (updatedParams params) {
  return append(append(updatedParams, param), self...)
}

// TODO: Add ability to output URL, and Path types, since these would be very
//       common and the ability to validate them would be nice. For example,
//       being able to check if a file exists easily.
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
  for reversedIndex := self.Count() - 1; reversedIndex >= 0; reversedIndex-- {
    reversedParams = append(reversedParams, self[reversedIndex])
  }
  return reversedParams
}

func (self params) IsZero() bool { return self.Count() == 0 }
func (self params) String() string { return strings.Join(self.Strings(), " ") }

func (self params) Strings() (paramStrings []string) { 
  for _, param := range self {
    paramStrings = append(paramStrings, param.value)
  }
  return paramStrings
}
