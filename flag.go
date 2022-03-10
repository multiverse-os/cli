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

func (self Flag) String() string { return self.Param.Value }
func (self Flag) Int() int { return self.Param.Int() }
func (self Flag) Bool() bool { return self.Param.Bool() }

func (self *Flag) Set(value string) *Flag {
  // TODO: Validate against param's validation (or create a param set that does
  // the validation and use that function preferably)
  self.Param.Value = value
  return self
}

func (self *Flag) SetDefault() *Flag {
  if len(self.Param.Value) == 0 {
    if len(self.Default) != 0 {
      self.Set(self.Default)
    }else{
      self.SetFalse()
    }
  }
  return self
}

// TODO: These should be replaced by a toggle, so if a bool is default
//       true, and the flag sets the variable to false, toggle will 
//       guarantee we have no edge case failrue. 
func (self *Flag) SetTrue() *Flag { return self.Set("1") }
func (self *Flag) SetFalse() *Flag { return self.Set("0") }

func (self *Flag) ToggleBoolean() *Flag {
  if data.IsTrue(self.Param.Value) {
    return self.SetFalse()
  }else{
    return self.SetTrue()
  }
}

///////////////////////////////////////////////////////////////////////////////

type flags []*Flag 

func Flags(flags ...Flag) (flagPointers flags) { 
  for index, _ := range flags {
    flags[index].Param = &Param{Value: flags[index].Default}
    flagPointers = append(flagPointers, &flags[index])
  }
  return flagPointers
}

func (self flags) Arguments() (arguments arguments) {
  for _, flag := range self {
    arguments = append(arguments, Argument(flag))
  }
  return arguments
}

func (self flags) Add(newFlag *Flag) (prepended flags) { 
  newFlag.Param = &Param{Value: newFlag.Default}
  // TODO: Previously it had validation, decide if its actually necessary
  //       because arguments for example doesnt have two return values; and we
  //       want consistency across the different add functions. if it does go
  //       in, we need it for all of them
  //err := ValidateFlag(*flag)
  return append(append(prepended, newFlag), self...)
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
  for index := self.Count() - 1; index >= 0; index-- {
    reversedFlags = append(reversedFlags, self[index])
  }
  return reversedFlags
}

func (self flags) SetDefaults() flags {
  for _, flag := range self {
    flag.SetDefault()
  }
  return self
}
