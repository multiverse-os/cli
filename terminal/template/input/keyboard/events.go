package keyboard

import (
	"fmt"
	"time"
)

// ButtonEvent represents an event when a keyboard button changes state (i.e.
// being pushed down when it was previously up, or being toggled on when it was
// previously off, etc)
//
// If Key == Invalid then the key may not be known, but it can still be
// uniquely identified and it's state watched via the Raw member (e.g. for
// special or non-US keys).
//
// The Raw member must uniquely identify the keyboard button whose state is
// changing, and must always be present regardless of whether or not Key ==
// Invalid. It could (but does not have to be) e.g. the scancode of the key.
type ButtonEvent struct {
	T     time.Time
	Key   Key
	State State
	Raw   uint64
}

// Time returns the time at which this event occured.
func (b ButtonEvent) Time() time.Time {
	return b.T
}

// String returns an string representation of this event.
func (b ButtonEvent) String() string {
	return fmt.Sprintf("ButtonEvent(Key=%v, State=%v, Raw=%v, Time=%v)", b.Key, b.State, b.Raw, b.T)
}

// Typed represents an event where some sort of user input has generated a
// string of text which should be considered as user input.
type Typed struct {
	T time.Time
	S string
}

// Time returns the time at which this event occured.
func (t Typed) Time() time.Time {
	return t.T
}

// String simply returns the user input string.
func (t Typed) String() string {
	return t.S
}
