package mouse

// Button represents a mouse button.
type Button int

// String implements fmt.Stringer()
func (b Button) String() string {
	if n, ok := buttonNames[b]; ok {
		return n
	}
	return "ButtonUnknown"
}

// buttonNames maps Button values to human readable names.
var buttonNames = map[Button]string{
	ButtonLeft:      "ButtonLeft",
	ButtonRight:     "ButtonRight",
	ButtonMiddle:    "ButtonMiddle",
	ButtonRelease:   "ButtonRelease",
	ButtonWheelUp:   "ButtonWheelUp",
	ButtonWheelDown: "ButtonWheelDown",
}

// Buttons recognized on the mouse.
const (
	buttonUnknown Button = iota
	ButtonLeft
	ButtonRight
	ButtonMiddle
	ButtonRelease
	ButtonWheelUp
	ButtonWheelDown
)
