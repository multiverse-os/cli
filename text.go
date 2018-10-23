package cli

// TODO: text.go is a first step in centralizing all the strings in the codebase to a single file so they can be merged
// into the locales/*.go system and then the switch can be made to support localization

// Error text
var (
	errInvalidActionType = NewExitError("ERROR invalid Action type. Must be `func(*Context`)` or `func(*Context) error).", 2)
	errDuplicateFlagName = "Cannot use two forms of the same flag %v %v"
	errIndexOutOfRange   = "index out of range"
)

//
// Text to be set in this file
///////////////////////////////////////////////////////////////////////////////
//return errors.New("Cannot use two forms of the same flag: " + name + " " + ff.Name)
//return errors.New("index out of range")
//defaultVal = fmt.Sprintf(" (default: %s)", strings.Join(defaultVals, ", "))

//return errors.New("index out of range")

//fmt.Fprintln(context.CLI.Writer, "Incorrect Usage:", err.Error())

//Usage:  "Print version",
//Usage:  "Print help text",

//fmt.Fprintf(self.Writer, "%s %s\n\n", "Incorrect Usage.", err.Error())

//return fmt.Errorf("could not parse %s as bool value for flag %s: %s", envVal, f.Name, err)
//return fmt.Errorf("could not parse %s as value for flag %s: %s", fileEnvVal, f.Name, err)
//return fmt.Errorf("could not parse %s as string value for flag %s: %s", envVal, f.Name, err)
//return fmt.Errorf("could not parse %s as int slice value for flag %s: %s", envVal, f.Name, err)
//return fmt.Errorf("could not parse %s as int64 slice value for flag %s: %s", envVal, f.Name, err)
//return fmt.Errorf("could not parse %s as uint value for flag %s: %s", envVal, f.Name, err)
//return fmt.Errorf("could not parse %s as uint64 value for flag %s: %s", envVal, f.Name, err)
//return fmt.Errorf("could not parse %s as duration for flag %s: %s", envVal, f.Name, err)

//fmt.Fprintf(self.Writer, "%s %s\n\n", "Incorrect Usage.", err.Error())

//Name:   "version, v",
//Usage:  "Print version",

//Name:   "help, h",
//Usage:  "Print help text",
