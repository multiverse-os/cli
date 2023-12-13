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

func (ft FlagType) TrimPrefix(flag string) string { return flag[int(ft):] }

func (ft FlagType) String() string {
	return strings.Repeat("-", int(ft))
}

func HasFlagPrefix(flag string) (FlagType, bool) {
	// NOTE: It is unnecessary to do the len(flag) != 0 check since arguments by
	// definition to be parsed by the OS must be not blank.
	if flag[0] == 45 {
		if strings.HasPrefix(flag, Long.String()) {
			return Long, true
		} else {
			return Short, true
		}
	}
	return UndefinedFlagType, false
}

///////////////////////////////////////////////////////////////////////////////

// Default boolean should be true, and if we dont assign it will be false, so we
// need to assign true somewhere; ALSO match boolean and default value
type Flag struct {
	Command     *Command
	Name        string
	Alias       string
	Description string
	Category    string
	Hidden      bool
	Required    bool
	Default     string
	Boolean     bool
	Action      Action
	Param       *Param
}

// TODO: This is should be BOTH setting the default, AND its
//
//	not even being used!
func ValidateFlag(flag Flag) error {
	// TODO: Validate param
	if 32 < len(flag.Name) {
		return ErrInvalidArgumentLength
	}
	if flag.Required && len(flag.Param.value) == 0 {
		return ErrArgumentRequired
	}
	// TODO: Validate format - we are just concerned about Linux POSIX
	//for _, flagRune := range flag.Name {
	//  // NOTE: a = 97; z = 122; - = 45
	//  if unicode.IsLetter(flagRune) || flagRune == 45 {
	//    fmt.Println("flagRune:", rune(flagRune))
	//    return ErrInvalidFlagFormat
	//  }
	//}
	return nil
}

func (f Flag) IsValid() bool { return ValidateFlag(f) != nil }

func (f Flag) is(name string) bool {
	return (len(f.Name) == len(name) && f.Name == name) ||
		(len(f.Alias) == len(name) && f.Alias == name)
}

func (f Flag) HasCategory() bool { return len(f.Category) != 0 }

func (f *Flag) String() string {
	if f != nil && f.Param != nil {
		return f.Param.value
	} else {
		return "0"
	}
}

func (f *Flag) Int() int         { return f.Param.Int() }
func (f *Flag) Bool() bool       { return f.Param.Bool() }
func (f *Flag) Float64() float64 { return f.Param.Float64() }

// TODO
// Path / File
// URL

func (f *Flag) Set(newValue string) *Flag {
	// TODO: Validate against param's validation (or create a param set that does
	// the validation and use that function preferably)
	f.Param = &Param{
		value: newValue,
	}
	return f
}

func (f *Flag) SetDefault() *Flag {
	if f.Param == nil && len(f.Param.value) == 0 && len(f.Default) != 0 {
		f.Param = &Param{value: f.Default}
	}
	return f
}

func (f *Flag) SetTrue() *Flag  { return f.Set("1") }
func (f *Flag) SetFalse() *Flag { return f.Set("0") }

func (f *Flag) Toggle() *Flag {
	if data.IsTrue(f.Param.value) {
		return f.SetFalse()
	} else {
		return f.SetTrue()
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

func (fs flags) Arguments() (flagArguments arguments) {
	for _, flag := range fs {
		flagArguments = append(flagArguments, Argument(flag))
	}
	return flagArguments
}

// TODO: We add three 1 for help 1 for version 1 for our name
func (fs *flags) Add(flag *Flag) {
	// TODO: Probably verify here???
	*fs = append(append(flags{}, flag), *fs...)
}

func (fs flags) Name(name string) *Flag {
	for _, flag := range fs {
		if flag.is(name) {
			return flag
		}
	}
	return nil
}

func (fs flags) Category(name string) (flagsInCategory flags) {
	for _, flag := range fs {
		// TODO: I hate string comparisons, maybe length check before
		if len(flag.Category) == len(name) && flag.Category == name {
			flagsInCategory = append(flagsInCategory, flag)
		}
	}
	return flagsInCategory
}

func (fs flags) HasFlag(name string) bool {
	return fs.Name(name) != nil
}

func (fs flags) Exists(name string) bool { return fs.HasFlag(name) }

func (fs flags) Visible() (visibleFlags flags) {
	for _, flag := range fs {
		if !flag.Hidden {
			visibleFlags = append(visibleFlags, flag)
		}
	}
	return visibleFlags
}

func (fs flags) Categories() (categories []string) {
	for _, flag := range fs {
		if flag.HasCategory() {
			var categoryExists bool
			for _, category := range categories {
				if len(category) == len(flag.Category) &&
					category == flag.Category {
					categoryExists = true
					break
				}
			}

			if !categoryExists {
				categories = append(categories, flag.Category)
			}
		}
	}
	return categories
}

func (fs flags) Validate() (errs []error) {
	for _, flag := range fs {
		if err := ValidateFlag(flag); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func (fs flags) Count() int   { return len(fs) }
func (fs flags) IsZero() bool { return fs.Count() == 0 }

func (fs flags) First() *Flag {
	if 0 < fs.Count() {
		return fs[0]
	}
	return nil
}

func (fs flags) Last() *Flag {
	if 0 < fs.Count() {
		return fs[len(fs)-1]
	}
	return nil
}

func (fs flags) Reverse() (reversedFlags flags) {
	for reversedIndex := fs.Count() - 1; reversedIndex >= 0; reversedIndex-- {
		reversedFlags = append(reversedFlags, fs[reversedIndex])
	}
	return reversedFlags
}

func (fs flags) SetDefaults() flags {
	for _, flag := range fs {
		flag.SetDefault()
	}
	return fs
}

// TODO: This will fix some issues, and make context.Flags make more sense, but
// will result in pretty large changes to the Parse() function
func (fs flags) Assigned() (assignedFlags flags) {
	for _, flag := range fs {
		// TODO: May need to just check param as it may never be initialized
		if len(flag.Param.value) != 0 {
			assignedFlags = append(assignedFlags, flag)
		}
	}
	return assignedFlags
}

func (fs flags) Unassigned() (unassignedFlags flags) {
	for _, flag := range fs {
		if len(flag.Param.value) == 0 {
			unassignedFlags = append(unassignedFlags, flag)
		}
	}
	return unassignedFlags
}
