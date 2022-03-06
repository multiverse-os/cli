package cli

import (
	"strings"
)

// TODO: When flag is blank, assume its boolean 

// TODO: Create a VersionFlag, HelpFlag, and DebugFlag all hidden by default and
// added by default to the global flags. 

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

type Flag struct {
	Command     *Command
	Name        string
  Alias       string
	Description string
	Hidden      bool
	Default     string
  Param       *Param
}

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

func (self *Flag) SetTrue() *Flag { return self.Set("1") }

func ValidateFlag(flag Flag) error {
  // TODO: Validate param
  if 32 < len(flag.Name) {
    return errInvalidArgumentLength
  }
  if len(flag.Alias) != 1 {
    return errInvalidFlagShortLength
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
}

// TODO: This required changing IsValid to return the error, and this must be
// done for param and command.
func (self flags) Add(flag *Flag) (flags, error) {
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
    return self[len(self)+1] 
  }
  return nil 
}
