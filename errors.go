package cli

import (
	"errors"
)

var (
	errInvalidActionType    = errors.New("[cli] invalid default action")
	errIndexOutOfRange      = errors.New("[cli] index out of range")
	errFailedNameAssignment = errors.New("[cli] failed to assign 'Name' attribute")
)
