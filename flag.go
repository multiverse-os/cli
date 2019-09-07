package cli

type Flags []Flag

type FlagType int

// TODO: These will provide for better help text, but will also allow for fine
// grain validation without needing to specify beyond FlagType.
// Even if special URLFlag, and such are decided to be a good idea, this will be
// far, far smaller than the original concept which relied on entire structs and
// files for each individual type and only allowed for a few, and maintaince or
// additions were nightmare scenerio.
const (
	BoolFlag FlagType = iota
	IntFlag
	StringFlag
	PathFlag
	FilenameFlag
	// TODO: Decide if this level of distincition will be necessary or would be
	// better handled with a further validaiton.
	//URLFlag FlagType = 5
	//IPFlag FlagType = 6
)

// TODO: This file is terrible, we can just use an interface and do a switchcase top determine type
// this will make a 700 line file maybe 100 lines
// TODO: Maybe just base what has - and -- on size (thats the original method)
type Flag struct {
	Name    string // Primary name
	Aliases []string
	Type    FlagType
	Usage   string
	Hidden  bool
	Value   interface{}
	// TODO: Previously this was being used to indicate where it was loaded from.
	// This may turn out to be a good idea, but for now it will remain commented
	// out.
	//EnvVar      string
	//Description string
	//Destination bool
	//Duration    bool
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
