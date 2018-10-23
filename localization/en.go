package localization

var en_US = Locale{
	// Example Localized Text
	//////////////////////////////////////////////////////////////////////////////
	//"key": LocalizedText{
	//	Data: map[string]string{
	//		"VarName":  "5",
	//		"OtherVar": "mega",
	//	},
	//	Message: "This {{.VarName}} is that {{.OtherVar}} thing",
	//},
	// TODO: %v need to be replaced with data key/value
	// en_US Localized Text
	//////////////////////////////////////////////////////////////////////////////
	"help":                      LocalizedText{Message: "help"},
	"help_v":                    LocalizedText{Message: "v"},
	"help_version":              LocalizedText{Message: "version"},
	"help_args_usage":           LocalizedText{Message: "[command]"},
	"help_command_not_found":    LocalizedText{Message: "No help topic for '%v'"},
	"help_name":                 LocalizedText{Message: "Name"},
	"help_category":             LocalizedText{Message: "Category"},
	"help_command":              LocalizedText{Message: "Command"},
	"help_commands":             LocalizedText{Message: "Commands"},
	"help_command_options":      LocalizedText{Message: "command options"},
	"help_arguments":            LocalizedText{Message: "arguments..."},
	"help_options":              LocalizedText{Message: "Options"},
	"help_usage":                LocalizedText{Message: "Usage"},
	"help_usage_text":           LocalizedText{Message: "List of available commands or details for a specified command"},
	"error_invalid_action_type": LocalizedText{Message: "ERROR invalid Action type. Must be `func(*Context`)` or `func(*Context) error)."},
	"error_duplicate_flag_name": LocalizedText{Message: "Cannot use two forms of the same flag. %v"},
	"error_index_out_of_range":  Localized_Text{Message: "index out of range"},
}
