package cli

// TODO: This file is terrible, we can just use an interface and do a switchcase top determine type
// this will make a 700 line file maybe 100 lines
type Flag struct {
	Name        string
	Alias       string
	Description string
	Usage       string
	EnvVar      string
	FilePath    string
	Hidden      bool
	Destination bool
	Duration    bool
	Value       interface{}
}

// TODO: How about we don't use globals?
//var VersionFlag Flag = BoolFlag{
//	Name:   "version",
//	Alias:  "v",
//	Usage:  "Print version",
//	Hidden: true,
//}
//
//var HelpFlag Flag = BoolFlag{
//	Name:   "help",
//	Alias:  "h",
//	Usage:  "Print help text",
//	Hidden: true,
//}
