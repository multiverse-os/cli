package cli

import (
	"errors"
)

var (
	ErrInvalidActionType      = errors.New("invalid default action")
	ErrIndexOutOfRange        = errors.New("index out of range")
	ErrFailedNameAssignment   = errors.New("failed to assign 'Name' attribute")
  ErrInvalidArgumentLength  = errors.New("maximum argument length is 32")
  ErrInvalidFlagShortLength = errors.New("maximum flag short length is 1")
  ErrInvalidArgumentFormat  = errors.New("invalid argument format")
)
