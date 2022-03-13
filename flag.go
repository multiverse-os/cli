package cli

import (
	"strings"

  data "github.com/multiverse-os/cli/data"
)
// TODO: Provide helpers/shortcuts for accessing flag.Param.Int() directly
// such as flag.Int()
type FlagType int

const (
  UndefinedFlagType FlagType = iota
	Short 
	Long
)

func (self FlagType) TrimPrefix(flag string) string { return flag[int(self):] }
func (self FlagType) String() string { return strings.Repeat("-", int(self)) }

func HasFlagPrefix(flag string) (FlagType, bool) {
  // NOTE: It is unnecessary to do the len(flag) != 0 check since arguments by
  // definition to be parsed by the OS must be not blank.
  if flag[0] == 45 {
	  if strings.HasPrefix(flag, Long.String()) {
	  	return Long, true
	  }else{
	  	return Short, true
	  }
  }
 	return UndefinedFlagType, false
}
///////////////////////////////////////////////////////////////////////////////

type Flag struct {
	Command     *Command
	Name        string
  Alias       string
	Description string
	Hidden      bool
	Default     string
  Param       *Param
}

func ValidateFlag(flag Flag) error {
  // TODO: Validate param
  if 32 < len(flag.Name) {
    return errInvalidArgumentLength
  }
  for _, flagRune := range flag.Name {
    // NOTE: a = 97; z = 122; - = 45
    if (97 <= flagRune && flagRune <= 122) || flagRune == 45 {
      return errInvalidArgumentFormat
    }
  }
  return nil
}

func (self Flag) IsValid() bool {  return ValidateFlag(self) != nil }

func (self Flag) is(name string) bool { 
  return self.Name == name || self.Alias == name
}

func (self Flag) String() string { return self.Param.value }
func (self Flag) Int() int { return self.Param.Int() }
func (self Flag) Bool() bool { return self.Param.Bool() }

func (self *Flag) To(newValue string) *Flag {
  // TODO: Validate against param's validation (or create a param set that does
  // the validation and use that function preferably)
  self.Param.value = newValue
  return self
}

func (self *Flag) ToDefault() *Flag {
  if len(self.Param.value) == 0 {
    if len(self.Default) != 0 {
      self.To(self.Default)
    }else{
      self.ToFalse()
    }
  }
  return self
}

func (self *Flag) ToTrue() *Flag { return self.To("1") }
func (self *Flag) ToFalse() *Flag { return self.To("0") }

func (self *Flag) ToggleBoolean() *Flag {
  if data.IsTrue(self.Param.value) {
    return self.ToFalse()
  }else{
    return self.ToTrue()
  }
}

///////////////////////////////////////////////////////////////////////////////

type flags []*Flag 

func Flags(flags ...Flag) (flagPointers flags) { 
  for index, _ := range flags {
    flags[index].Param = &Param{value: flags[index].Default}
    flagPointers = append(flagPointers, &flags[index])
  }
  return flagPointers
}

func (self flags) Arguments() (flagArguments arguments) {
  for _, flag := range self {
    flagArguments = append(flagArguments, Argument(flag))
  }
  return flagArguments
}

func (self flags) Add(flag *Flag) (updatedFlags flags) { 
  flag.Param = &Param{value: flag.Default}
  return append(append(updatedFlags, flag), self...)
}

func (self flags) Name(name string) *Flag {
  for _, flag := range self {
    if flag.is(name) {
      return flag
    }
  }
  return nil
}

func (self flags) Visible() (visibleFlags flags) {
  for _, flag := range self {
    if !flag.Hidden {
      visibleFlags = append(visibleFlags, flag)
    }
  }
  return visibleFlags
}

func (self flags) Validate() error {
  for _, flag := range self {
    if err := ValidateFlag(*flag); err != nil {
      return err 
    }
  }
  return nil
}

func (self flags) Count() int { return len(self) }
func (self flags) IsZero() bool { return self.Count() == 0 }

func (self flags) First() *Flag {
  if 0 < self.Count() {
    return self[0] 
  }
  return nil
}

func (self flags) Last() *Flag { 
  if 0 < self.Count() {
    return self[len(self)-1] 
  }
  return nil 
}

func (self flags) Reverse() (reversedFlags flags) {
  for reversedIndex := self.Count() - 1; reversedIndex >= 0; reversedIndex-- {
    reversedFlags = append(reversedFlags, self[reversedIndex])
  }
  return reversedFlags
}

func (self flags) ToDefaults() flags {
  for _, flag := range self {
    flag.ToDefault()
  }
  return self
}
