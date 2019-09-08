package cli

import (
	"errors"
	"io"
	"os"
)

var (
	errInvalidActionType = errors.New("[cli] invalid default action")
	errIndexOutOfRange   = errors.New("[cli] index out of range")
)

var OsExiter = os.Exit
var ErrWriter io.Writer = os.Stderr
