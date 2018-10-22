package term

import (
	"errors"
	"os"
	"syscall"
	"unsafe"
)

var ErrGetWinsizeFailed = errors.New("term: syscall.TIOCGWINSZ failed")

func GetTerminalWindowSize(file *os.File) (*Size, error) {
	var dimensions [4]uint16
	if _, _, err := syscall.Syscall6(syscall.SYS_IOCTL, file.Fd(), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&dimensions)), 0, 0, 0); err != 0 {
		return nil, err
	}

	return &Size{
		Lines:   int(dimensions[0]),
		Columns: int(dimensions[1]),
	}, nil
}
