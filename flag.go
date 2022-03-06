package cli

import (
  "fmt"
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

  // Prob #1 priorirty
  // TODO: Before adding a check if param exists and creating a if condition to
  //       either assign or create param and assign-- we should check if the
  //       defaults are being correctly defined (and only once)
  //if self.P
  fmt.Println("self.Param is nil?", self.Param.Value)
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

func (self flags) Reversed() (reversedFlags flags) {
  for i := self.Count() - 1; i >= 0; i-- {
    reversedFlags = append(reversedFlags, self[i])
  }
  return reversedFlags
  // TODO: Using the principle of this function we could add a Prepend and make
  // add Append but we will only be prepending and avoiding looping through and
  // resorting it more than once
}

// TODO: This required changing IsValid to return the error, and this must be
// done for param and command.
func (self flags) Add(flag *Flag) (flags, error) {
  // TODO: Add should prepend, by looping thruogh and assigning to a new loop
  // that is initiated with our new flag
  // TODO: Doesn't correctly handle this; default can be blank, and we need a
  // default value. Also if value already exists, we don't overwrite it!
  flag.Param = &Param{Value: flag.Default}
  err := ValidateFlag(*flag)
  if err != nil {
    return append(self, flag), err
  }else{
    return self, err
  }
}

func (self flags) Name(name string) *Flag {
  for _, flag := range self {
    fmt.Println("checking flag with name:", flag.Name)
    fmt.Println("                   alias:", flag.Alias)
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

func (self flags) SetDefaults() flags {
  for _, flag := range self {
    flag.SetDefault()
  }
  return self
}
