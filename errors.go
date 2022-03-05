package cli

import (
	"errors"
)

var (
	errInvalidActionType      = errors.New("invalid default action")
	errIndexOutOfRange        = errors.New("index out of range")
	errFailedNameAssignment   = errors.New("failed to assign 'Name' attribute")
  errInvalidArgumentLength  = errors.New("maximum argument length is 32")
  errInvalidFlagShortLength = errors.New("maximum flag short length is 1")
  errInvalidArgumentFormat  = errors.New("invalid argument format")
)
