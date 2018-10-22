package cli

// text.go is a first step in centralizing all the strings in the codebase to a single file so they can be merged
// into the locales/*.go system and then the switch can be made to support localization

// Error text
var (
	errInvalidActionType = NewExitError("ERROR invalid Action type. Must be `func(*Context`)` or `func(*Context) error).", 2)
)
