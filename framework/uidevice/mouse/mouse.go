package mouse

// State represents a single mouse state, such as Up or Down.
type State uint8

// Mouse button state constants, Down implies the button is currently pressed
// down, and up implies it is not. The InvalidState is declared to help users
// detect uninitialized variables.
const (
	InvalidState State = iota
	Down
	Up
)

// Button represents a single mouse button.
type Button uint8

// Mouse button constants for buttons one through eight. The Invalid button is
// declared to help users detect uninitialized variables.
const (
	Invalid Button = iota
	One
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
)

// Left, Right, Middle and Wheel are simply aliases. Their true names are mouse
// button One, Two, and Three (for both Middle and Wheel, respectively).
const (
	Left   = One
	Right  = Two
	Wheel  = Three
	Middle = Three
)
