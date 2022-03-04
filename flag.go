package cli

import (
	"strings"

	data "github.com/multiverse-os/cli/data"
)


type FlagType int

const (
	Short FlagType = iota + 1
	Long
	NotAvailable
)

func (self FlagType) TrimPrefix(flagArgument string) string { 
  return flagArgument[self.Length():]
}

// NOTE: DEV Function
func (self FlagType) Name() string {
	switch self {
	case Short:
		return "short"
	case Long:
		return "long"
	default: // NotAvailable
		return "n/a"
	}
}

func (self FlagType) Is(flagType FlagType) bool { return self == flagType }
func (self FlagType) Length() int               { return int(self) }
func (self FlagType) String() string            { return strings.Repeat("-", self.Length()) }

type FlagSeparator int

const (
	Whitespace FlagSeparator = iota
	Equal
)

func (self FlagSeparator) Is(flagSeparator FlagSeparator) bool { return self == flagSeparator }

func (self FlagSeparator) String() string {
	if self.Is(Equal) {
		return "="
	} else {
		return " "
	}
}


// TODO: Be able to define the file extension that would be selected for when generating an autocomplete file
//        -
//       Could eventually build out more functionality to support autocomplete
type Flag struct {
	Command     *Command
	Level       Level
	Name        string
	Alias       string
	Description string
	Options     []string
	Hidden      bool
	Default     string

  Param       *Param
}

// NOTE: Can be thought of as equivilent to New() but flags are a special
//       sub-type that do not exist without a command. 
func (self *Flag) Copy() (newFlag *Flag) {
  newFlag = self
  return newFlag
}

func (self *Flag) Set(value string) *Flag {
  // TODO: Validate against param's validation
  self.Param.Value = value
  return self
}

func ValidateFlag(flag Flag) error {
  if 32 < len(flag.Name) {
    return errInvalidArgumentLength
  }
  for _, flagRune := range flag.Name {
    // NOTE: 
    // a = 97
    // z = 122
    // - = 45
    if (97 <= flagRune && flagRune <= 122) || flagRune == 45 {
      return errInvalidArgumentFormat
    }
  }
  return nil
}

func (self Flag) IsValid() bool { return ValidateFlag(self) != nil }

type flags []*Flag 

func Flags(flags ...Flag) (flagPointers flags) { 
  for index, _ := range flags {
    flagPointers = append(flagPointers, &flags[index])
  }
  return flagPointers
}

func (self flags) Add(flag *Flag) flags {
  return append(self, flag)
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

func (self flags) Hidden() (hiddenFlags flags) {
  for _, flag := range self {
    if flag.Hidden {
      hiddenFlags = append(hiddenFlags, flag)
    }
  }
  return hiddenFlags
}

func (self flags) Level(level Level) (f flags) {
  for _, flag := range self {
    if flag.Level == level {
      f = append(f, flag)
    }
  }
  return f
}

func (self flags) Count() int { return len(self) }
func (self flags) IsZero() bool { return self.Count() == 0 }

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
 	return NotAvailable, false
}

// TODO: Added to ToLower here where it should ahve beent the whole time; so as
// a consequence of bad programming before we need to remove ToLowers where find
// them elsewhere 
func (self Flag) is(name string) bool { 
  name = strings.ToLower(name)
  return self.Name == name || self.Alias == name 
}

func (self Flag) flagNames() (output string) { return output }

func (self Flag) help() string {
	usage := Long.String() + self.Name
	if data.NotBlank(self.Alias) {
		usage += ", " + Short.String() + self.Alias
	}
	var defaultValue string
	if len(self.Default) != 0 {
		defaultValue = " [≅ " + self.Default + "]"
	}
	return strings.Repeat(" ", 4) + usage + strings.Repeat(" ", 18-len(usage)) + self.Description + defaultValue + "\n"
}


// TODO: I like these and they are similar to the idea had earlier for a active
// record analogue
func (self Flag) Type() ArgumentType { return FlagArgument }

func (self Flag) String() string { return self.Param.String() }
func (self Flag) Int() int { return self.Param.Int() }
func (self Flag) Bool() bool { return self.Param.Bool() }
