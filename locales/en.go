package localization

// TODO: For a usable translation system, we need a way to insert in subject, because
// for example in other languages the ordering would not be "no help topics found for {SUBJECT}"
// so being able to move around the {SUBJECT} is a very important aspect of localization support

var en_US = map[string]string{
	"help":                   "help",
	"help_args_usage":        "[command]",
	"help_command_not_found": "No help topics found for",
	"help_name":              "Name",
	"help_category":          "Category",
	"help_command":           "Command",
	"help_commands":          "Commands",
	"help_command_options":   "command options",
	"help_arguments":         "arguments...",
	"help_options":           "Options",
	"help_usage":             "Usage",
	"help_usage_text":        "List of available commands or details for a specified command",
}
