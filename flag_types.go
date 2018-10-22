package cli

import (
	"flag"
	"strconv"
	"time"
)

type BoolFlag struct {
	Name        string
	Usage       string
	EnvVar      string
	FilePath    string
	Hidden      bool
	Destination *bool
}

func (f BoolFlag) String() string {
	return FlagStringer(f)
}

func (f BoolFlag) GetName() string {
	return f.Name
}

func (c *Context) Bool(name string) bool {
	return lookupBool(name, c.flagSet)
}

func (c *Context) GlobalBool(name string) bool {
	if fs := lookupGlobalFlagSet(name, c); fs != nil {
		return lookupBool(name, fs)
	}
	return false
}

func lookupBool(name string, set *flag.FlagSet) bool {
	f := set.Lookup(name)
	if f != nil {
		parsed, err := strconv.ParseBool(f.Value.String())
		if err != nil {
			return false
		}
		return parsed
	}
	return false
}

type BoolTFlag struct {
	Name        string
	Usage       string
	EnvVar      string
	FilePath    string
	Hidden      bool
	Destination *bool
}

func (f BoolTFlag) String() string {
	return FlagStringer(f)
}

func (f BoolTFlag) GetName() string {
	return f.Name
}

func (c *Context) BoolT(name string) bool {
	return lookupBoolT(name, c.flagSet)
}

func (c *Context) GlobalBoolT(name string) bool {
	if fs := lookupGlobalFlagSet(name, c); fs != nil {
		return lookupBoolT(name, fs)
	}
	return false
}

func lookupBoolT(name string, set *flag.FlagSet) bool {
	f := set.Lookup(name)
	if f != nil {
		parsed, err := strconv.ParseBool(f.Value.String())
		if err != nil {
			return false
		}
		return parsed
	}
	return false
}

type DurationFlag struct {
	Name        string
	Usage       string
	EnvVar      string
	FilePath    string
	Hidden      bool
	Value       time.Duration
	Destination *time.Duration
}

func (f DurationFlag) String() string {
	return FlagStringer(f)
}

func (f DurationFlag) GetName() string {
	return f.Name
}

func (c *Context) Duration(name string) time.Duration {
	return lookupDuration(name, c.flagSet)
}

func (c *Context) GlobalDuration(name string) time.Duration {
	if fs := lookupGlobalFlagSet(name, c); fs != nil {
		return lookupDuration(name, fs)
	}
	return 0
}

func lookupDuration(name string, set *flag.FlagSet) time.Duration {
	f := set.Lookup(name)
	if f != nil {
		parsed, err := time.ParseDuration(f.Value.String())
		if err != nil {
			return 0
		}
		return parsed
	}
	return 0
}

type Float64Flag struct {
	Name        string
	Usage       string
	EnvVar      string
	FilePath    string
	Hidden      bool
	Value       float64
	Destination *float64
}

func (f Float64Flag) String() string {
	return FlagStringer(f)
}

func (f Float64Flag) GetName() string {
	return f.Name
}

func (c *Context) Float64(name string) float64 {
	return lookupFloat64(name, c.flagSet)
}

func (c *Context) GlobalFloat64(name string) float64 {
	if fs := lookupGlobalFlagSet(name, c); fs != nil {
		return lookupFloat64(name, fs)
	}
	return 0
}

func lookupFloat64(name string, set *flag.FlagSet) float64 {
	f := set.Lookup(name)
	if f != nil {
		parsed, err := strconv.ParseFloat(f.Value.String(), 64)
		if err != nil {
			return 0
		}
		return parsed
	}
	return 0
}

type GenericFlag struct {
	Name     string
	Usage    string
	EnvVar   string
	FilePath string
	Hidden   bool
	Value    Generic
}

func (f GenericFlag) String() string {
	return FlagStringer(f)
}

func (f GenericFlag) GetName() string {
	return f.Name
}

func (c *Context) Generic(name string) interface{} {
	return lookupGeneric(name, c.flagSet)
}

func (c *Context) GlobalGeneric(name string) interface{} {
	if fs := lookupGlobalFlagSet(name, c); fs != nil {
		return lookupGeneric(name, fs)
	}
	return nil
}

func lookupGeneric(name string, set *flag.FlagSet) interface{} {
	f := set.Lookup(name)
	if f != nil {
		parsed, err := f.Value, error(nil)
		if err != nil {
			return nil
		}
		return parsed
	}
	return nil
}

type Int64Flag struct {
	Name        string
	Usage       string
	EnvVar      string
	FilePath    string
	Hidden      bool
	Value       int64
	Destination *int64
}

func (f Int64Flag) String() string {
	return FlagStringer(f)
}

func (f Int64Flag) GetName() string {
	return f.Name
}

func (c *Context) Int64(name string) int64 {
	return lookupInt64(name, c.flagSet)
}

func (c *Context) GlobalInt64(name string) int64 {
	if fs := lookupGlobalFlagSet(name, c); fs != nil {
		return lookupInt64(name, fs)
	}
	return 0
}

func lookupInt64(name string, set *flag.FlagSet) int64 {
	f := set.Lookup(name)
	if f != nil {
		parsed, err := strconv.ParseInt(f.Value.String(), 0, 64)
		if err != nil {
			return 0
		}
		return parsed
	}
	return 0
}

type IntFlag struct {
	Name        string
	Usage       string
	EnvVar      string
	FilePath    string
	Hidden      bool
	Value       int
	Destination *int
}

func (f IntFlag) String() string {
	return FlagStringer(f)
}

func (f IntFlag) GetName() string {
	return f.Name
}

func (c *Context) Int(name string) int {
	return lookupInt(name, c.flagSet)
}

func (c *Context) GlobalInt(name string) int {
	if fs := lookupGlobalFlagSet(name, c); fs != nil {
		return lookupInt(name, fs)
	}
	return 0
}

func lookupInt(name string, set *flag.FlagSet) int {
	f := set.Lookup(name)
	if f != nil {
		parsed, err := strconv.ParseInt(f.Value.String(), 0, 64)
		if err != nil {
			return 0
		}
		return int(parsed)
	}
	return 0
}

type IntSliceFlag struct {
	Name     string
	Usage    string
	EnvVar   string
	FilePath string
	Hidden   bool
	Value    *IntSlice
}

func (f IntSliceFlag) String() string {
	return FlagStringer(f)
}

func (f IntSliceFlag) GetName() string {
	return f.Name
}

func (c *Context) IntSlice(name string) []int {
	return lookupIntSlice(name, c.flagSet)
}

func (c *Context) GlobalIntSlice(name string) []int {
	if fs := lookupGlobalFlagSet(name, c); fs != nil {
		return lookupIntSlice(name, fs)
	}
	return nil
}

func lookupIntSlice(name string, set *flag.FlagSet) []int {
	f := set.Lookup(name)
	if f != nil {
		parsed, err := (f.Value.(*IntSlice)).Value(), error(nil)
		if err != nil {
			return nil
		}
		return parsed
	}
	return nil
}

type Int64SliceFlag struct {
	Name     string
	Usage    string
	EnvVar   string
	FilePath string
	Hidden   bool
	Value    *Int64Slice
}

func (f Int64SliceFlag) String() string {
	return FlagStringer(f)
}

func (f Int64SliceFlag) GetName() string {
	return f.Name
}

func (c *Context) Int64Slice(name string) []int64 {
	return lookupInt64Slice(name, c.flagSet)
}

func (c *Context) GlobalInt64Slice(name string) []int64 {
	if fs := lookupGlobalFlagSet(name, c); fs != nil {
		return lookupInt64Slice(name, fs)
	}
	return nil
}

func lookupInt64Slice(name string, set *flag.FlagSet) []int64 {
	f := set.Lookup(name)
	if f != nil {
		parsed, err := (f.Value.(*Int64Slice)).Value(), error(nil)
		if err != nil {
			return nil
		}
		return parsed
	}
	return nil
}

type StringFlag struct {
	Name        string
	Usage       string
	EnvVar      string
	FilePath    string
	Hidden      bool
	Value       string
	Destination *string
}

func (f StringFlag) String() string {
	return FlagStringer(f)
}

func (f StringFlag) GetName() string {
	return f.Name
}

func (c *Context) String(name string) string {
	return lookupString(name, c.flagSet)
}

func (c *Context) GlobalString(name string) string {
	if fs := lookupGlobalFlagSet(name, c); fs != nil {
		return lookupString(name, fs)
	}
	return ""
}

func lookupString(name string, set *flag.FlagSet) string {
	f := set.Lookup(name)
	if f != nil {
		parsed, err := f.Value.String(), error(nil)
		if err != nil {
			return ""
		}
		return parsed
	}
	return ""
}

type StringSliceFlag struct {
	Name     string
	Usage    string
	EnvVar   string
	FilePath string
	Hidden   bool
	Value    *StringSlice
}

func (f StringSliceFlag) String() string {
	return FlagStringer(f)
}

func (f StringSliceFlag) GetName() string {
	return f.Name
}

func (c *Context) StringSlice(name string) []string {
	return lookupStringSlice(name, c.flagSet)
}

func (c *Context) GlobalStringSlice(name string) []string {
	if fs := lookupGlobalFlagSet(name, c); fs != nil {
		return lookupStringSlice(name, fs)
	}
	return nil
}

func lookupStringSlice(name string, set *flag.FlagSet) []string {
	f := set.Lookup(name)
	if f != nil {
		parsed, err := (f.Value.(*StringSlice)).Value(), error(nil)
		if err != nil {
			return nil
		}
		return parsed
	}
	return nil
}

type Uint64Flag struct {
	Name        string
	Usage       string
	EnvVar      string
	FilePath    string
	Hidden      bool
	Value       uint64
	Destination *uint64
}

func (f Uint64Flag) String() string {
	return FlagStringer(f)
}

func (f Uint64Flag) GetName() string {
	return f.Name
}

func (c *Context) Uint64(name string) uint64 {
	return lookupUint64(name, c.flagSet)
}

func (c *Context) GlobalUint64(name string) uint64 {
	if fs := lookupGlobalFlagSet(name, c); fs != nil {
		return lookupUint64(name, fs)
	}
	return 0
}

func lookupUint64(name string, set *flag.FlagSet) uint64 {
	f := set.Lookup(name)
	if f != nil {
		parsed, err := strconv.ParseUint(f.Value.String(), 0, 64)
		if err != nil {
			return 0
		}
		return parsed
	}
	return 0
}

type UintFlag struct {
	Name        string
	Usage       string
	EnvVar      string
	FilePath    string
	Hidden      bool
	Value       uint
	Destination *uint
}

func (f UintFlag) String() string {
	return FlagStringer(f)
}

func (f UintFlag) GetName() string {
	return f.Name
}

func (c *Context) Uint(name string) uint {
	return lookupUint(name, c.flagSet)
}

func (c *Context) GlobalUint(name string) uint {
	if fs := lookupGlobalFlagSet(name, c); fs != nil {
		return lookupUint(name, fs)
	}
	return 0
}

func lookupUint(name string, set *flag.FlagSet) uint {
	f := set.Lookup(name)
	if f != nil {
		parsed, err := strconv.ParseUint(f.Value.String(), 0, 64)
		if err != nil {
			return 0
		}
		return uint(parsed)
	}
	return 0
}
