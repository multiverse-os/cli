package uidevice

import (
	"github.com/multiverse-os/cli/framework/uidevice/keyboard"
	"github.com/multiverse-os/cli/framework/uidevice/mouse"
)

// Polling represents the update speed for detecting input.
type Polling uint8

// Polling mode constants, Eco will check the input ten times per second.
// Normal will check thirty times per seconds and Game will uodate at sixty
// times per second. Default is Eco which is suggested for standard console
// or user input that does not require a high update interval
const (
	Eco    Polling = iota // 10x/sec
	Normal                // 30x/sec
	Game                  // 60x/sec
)

// InputState holds all current input states
type InputState struct {
	Keyboard keyboard.State
	Mouse    mouse.State
}

var (
	// DefaultKeyboard allows immediate access to default keyboard input
	DefaultKeyboard = NewKeyboard()
	// Defaultmouse allows immediate access to default mouse input
	DefaultMouse = NewMouse()
)

func NewKeyboard() *keyboard.Watcher {
	return keyboard.NewWatcher()
}

func NewMouse() *mouse.Watcher {
	return mouse.NewWatcher()
}
