package text

import (
	"syscall"
	"unsafe"
)

type Terminal struct {
	Size terminalSize
}

type terminalSize struct {
	CharacterWidth  uint16
	CharacterHeight uint16
	PixelWidth      uint16
	PixelHeight     uint16
}

// Public for now because the terminal object is just an idea or suggestion
func TerminalWidth() uint {
	tSize := &terminalSize{}
	retCode, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(tSize)),
	)

	if int(retCode) == -1 {
		panic(errno)
	}
	return uint(tSize.CharacterWidth)
}
