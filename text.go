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

//Info("Shutdown initiated, gracefully closing outputs...")
//Info("Logger currently has no specified outputs, defaulting to ANSI-styled StdOut output.")
//FatalError(errors.New("Failed to initialize default log path: '" + logPath + "'"))
//FatalError(errors.New("Name attribute is required to initialize log file"))

//return errors.New("Cannot use two forms of the same flag: " + name + " " + ff.Name)
//return errors.New("index out of range")
//defaultVal = fmt.Sprintf(" (default: %s)", strings.Join(defaultVals, ", "))

//fmt.Fprintf(self.Writer, "%s %s\n\n", "Incorrect Usage.", err.Error())

//fmt.Fprintln(context.CLI.Writer, "Incorrect Usage:", err.Error())

//Usage:  "Print version",

//Usage:     "List of available commands or details for a specified command",

//log.Print(log.FATAL, "Failed to parse executable working directory in default 'Name' attribute assignment.")

//return errors.New("index out of range")

//fmt.Fprintln(context.CLI.Writer, "Incorrect Usage:", err.Error())

//Usage:  "Print version",
//Usage:  "Print help text",

//fmt.Fprintf(self.Writer, "%s %s\n\n", "Incorrect Usage.", err.Error())

//fmt.Fprintf(self.Writer, "%s %s\n\n", "Incorrect Usage.", err.Error())

//Name:   "version, v",
//Usage:  "Print version",

//Name:   "help, h",
//Usage:  "Print help text",
