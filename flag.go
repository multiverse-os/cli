package cli

import (
	"strconv"
	"strings"

	data "github.com/multiverse-os/cli/data"
)


type FlagType int

const (
	Short FlagType = iota + 1
	Long
	NotAvailable
)

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

type FlagLevel uint8

const (
	GlobalFlag FlagLevel = iota
	CommandFlag
)

// TODO: Be able to define the file extension that would be selected for when generating an autocomplete file
//        -
//       Could eventually build out more functionality to support autocomplete
type Flag struct {
	Command     *Command
	Level       FlagLevel
	Name        string
	Alias       string
	Description string
	Options     []string
	Hidden      bool
	Default     string
	Value       string
	DataType    data.Type
}

type flags []*Flag 

func Flags(definedFlags ...Flag) (flagPointers flags) { 
  for _, flag := range definedFlags {
    flagPointers = append(flagPointers, &flag)
  }
  return flagPointers
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

func (self flags) Scope(scope FlagLevel) (f flags) {
  for _, flag := range self {
    if flag.Level == scope {
      f = append(f, flag)
    }
  }
  return f
}


func HasFlagPrefix(flag string) (FlagType, bool) {
	if strings.HasPrefix(flag, Long.String()) {
		return Long, true
	} else if strings.HasPrefix(flag, Short.String()) {
		return Short, true
	} else {
		return NotAvailable, false
	}
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

func (self Flag) String() string { return self.Value }

func (self Flag) Int() int {
	intValue, err := strconv.Atoi(self.Value)
	if err != nil {
		return 0
	} else {
		return intValue
	}
}

func (self Flag) Bool() bool { return data.IsTrue(self.Value) }
