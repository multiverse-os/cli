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
//
//	Path:     4096
//	Filename: 256
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

func (p Param) IsValid() bool { return ValidateParam(p) != nil }

// TODO: These should more heavily rely on existing code in data subpackage
func (p Param) Value() string  { return p.value }
func (p Param) String() string { return p.value }
func (p Param) Bool() bool     { return data.IsTrue(p.value) }

func (p Param) Int() int {
	intValue, err := strconv.Atoi(p.value)
	if err != nil {
		return 0
	} else {
		return intValue
	}
}

// TODO: Float

// TODO: Path / Filename

// TODO: URL

// /////////////////////////////////////////////////////////////////////////////
type params []*Param

func Params(params ...Param) (paramPointers params) {
	for index, _ := range params {
		paramPointers = append(paramPointers, &params[index])
	}
	return paramPointers
}

func (ps params) Arguments() (newArguments arguments) {
	for _, param := range ps {
		newArguments = append(newArguments, param)
	}
	return newArguments
}

// TODO: COnvert to *params method I laid out in other data types that lets us
// do .Add without reassigning to itself
func (ps params) Add(param *Param) (updatedParams params) {
	return append(append(updatedParams, param), ps...)
}

// TODO: Add ability to output URL, and Path types, since these would be very
//
//	common and the ability to validate them would be nice. For example,
//	being able to check if a file exists easily.
func (ps params) Count() int { return len(ps) }

func (ps params) First() *Param {
	if 0 < ps.Count() {
		return ps[0]
	}
	return nil
}

func (ps params) Last() *Param {
	if 0 < ps.Count() {
		return ps[ps.Count()-1]
	}
	return nil
}

func (ps params) Reverse() (reversedParams params) {
	for reversedIndex := ps.Count() - 1; reversedIndex >= 0; reversedIndex-- {
		reversedParams = append(reversedParams, ps[reversedIndex])
	}
	return reversedParams
}

func (ps params) IsZero() bool   { return ps.Count() == 0 }
func (ps params) String() string { return strings.Join(ps.Strings(), " ") }

func (ps params) Strings() (paramStrings []string) {
	for _, param := range ps {
		paramStrings = append(paramStrings, param.value)
	}
	return paramStrings
}
