package cli

import (
	"fmt"
)

var (
	errInvalidActionType    = fmt.Errorf("invalid default action")
	errIndexOutOfRange      = fmt.Errorf("index out of range")
	errFailedNameAssignment = fmt.Errorf("failed to assign 'Name' attribute")
)
