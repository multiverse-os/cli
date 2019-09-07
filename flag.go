package cli

type Flags []Flag

type FlagType int

const (
	BoolFlag FlagType = iota
	IntFlag
	StringFlag
	PathFlag
	FilenameFlag
)

type Flag struct {
	Name    string // Primary name
	Aliases []string
	Type    FlagType
	Usage   string
	Hidden  bool
	Value   interface{}
}

func (flag Flag) Names() []string {
	return append([]string{flag.Name}, flag.Aliases...)
}

func prefixFor(name string) string {
	if len(name) == 1 {
		return "-"
	} else {
		return "--"
	}
}

// TODO: How about we don't use globals?
var VersionFlag Flag = Flag{
	Name:    "version",
	Aliases: []string{"v"},
	Usage:   "Print version",
	Hidden:  true,
}

var HelpFlag Flag = Flag{
	Name:    "help",
	Aliases: []string{"h"},
	Usage:   "Print help text",
	Hidden:  true,
}
