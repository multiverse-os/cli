package cli

import (
	"errors"
)

var (
	errInvalidActionType    = errors.New("invalid default action")
	errIndexOutOfRange      = errors.New("index out of range")
	errFailedNameAssignment = errors.New("failed to assign 'Name' attribute")
)
