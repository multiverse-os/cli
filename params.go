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

func (self Param) IsValid() bool {
  return ValidateParam(self) != nil
}

// TODO: Param.String() 
func (self Param) String() string {
  return self.Value
}

func (self Param) Int() int {
  intValue, err := strconv.Atoi(self.Value[0:1])
	if err != nil {
		return 0
	} else {
		return intValue
	}
}

func (self Param) Bool() bool {
	for _, trueString := range data.True.Strings() {
    if trueString == self.Value[0:1] {
			return true
		}
	}
	return false
}

// TODO: Float 

// TODO: Path / Filename

// TODO: URL

// TODO: 

type params []*Param

func Params(params ...Param) (paramPointers params) { 
  for index, _ := range params {
    paramPointers = append(paramPointers, &params[index])
  }
  return paramPointers
}

func (self params) Add(param string) (params, error) { 
  newParam := &Param{Value: param}
  err := ValidateParam(*newParam)
  if err != nil {
    return append(self, newParam), err
  }else{
    return self, err
  }
}

// TODO: Add ability to output URL, and Path types, since these would be very
//       common and the ability to validate them would be nice. For example,
//       being able to check if a file exists easily.

// TODO: Once the params have been loaded, begin loading flags again; then
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
