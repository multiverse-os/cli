package terminal

import (
	"syscall"
	"unsafe"

	"github.com/multiverse-os/cli/os"
)

type Terminal struct {
	Title  string
	Cursor cursor.Cursor // TODO: Really want to support multiseat

	// TODO: At least an active buffer, and shaddow buffer, but possibly several
	//       so one can be an overlay, etc.

	// TODO: Require, Ticks, Subscribers, information about the logged in user
	//       etc
	User             os.User
	WorkingDirectory os.Path

	Dimensions terminalDimensions
}

type terminalDimensions struct {
	CharacterWidth  uint16
	CharacterHeight uint16
	PixelWidth      uint16
	PixelHeight     uint16
}

func New() Terminal {
	return Terminal{
		Dimensions: terminalDimensions{
			CharacterWidth:  0,
			CharacterHeight: 0,
		},
	}
}

func Width() uint {
	dimensions := &terminalDimensions{}
	data, _, err := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(dimensions)),
	)

	if int(data) == -1 {
		panic(err)
	}
	return uint(dimensions.CharacterWidth)
}
