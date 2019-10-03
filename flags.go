package cli

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

func (flag Flag) Names() []string { return append([]string{flag.Name}, flag.Aliases...) }
func (self Flag) Visible() bool   { return !self.Hidden }

func (self Flag) String() (output string) {
	if len(self.Aliases) > 0 {
		if len(self.Aliases[0]) >= 2 {
			output += "--" + self.Aliases[0]
		} else {
			output += "-" + self.Aliases[0]
		}
	}
	output += ", --" + self.Name
	return output
}

func (self Flag) Is(name string) bool {
	for _, flagName := range self.Names() {
		if flagName == name {
			return true
		}
	}
	return false
}

func defaultFlags() []Flag {
	return []Flag{
		Flag{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "Print version",
			Hidden:  false,
		},
		Flag{
			Name:    "help",
			Aliases: []string{"h"},
			Usage:   "Print help text",
			Hidden:  false,
		},
	}
}

func defaultCommandFlags() []Flag {
	return []Flag{
		Flag{
			Name:    "help",
			Aliases: []string{"h"},
			Usage:   "Print help text",
			Hidden:  false,
		},
	}
}
