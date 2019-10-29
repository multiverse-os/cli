package keyboard

// Key represents a single button on the keyboard.
// Printable characters are set to their ASCII/Unicode rune value.
// Non-printable (control) characters are equal to one of the constants defined
// below.
type Key rune

// String implements fmt.Stringer()
func (b Key) String() string {
	if n, ok := buttonNames[b]; ok {
		return n
	} else if b >= 0 {
		return string(b)
	}
	return "KeyUnknown"
}

// buttonNames maps Key values to human readable names.
var buttonNames = map[Key]string{
	KeyF1:         "KeyF1",
	KeyF2:         "KeyF2",
	KeyF3:         "KeyF3",
	KeyF4:         "KeyF4",
	KeyF5:         "KeyF5",
	KeyF6:         "KeyF6",
	KeyF7:         "KeyF7",
	KeyF8:         "KeyF8",
	KeyF9:         "KeyF9",
	KeyF10:        "KeyF10",
	KeyF11:        "KeyF11",
	KeyF12:        "KeyF12",
	KeyInsert:     "KeyInsert",
	KeyDelete:     "KeyDelete",
	KeyHome:       "KeyHome",
	KeyEnd:        "KeyEnd",
	KeyPgUp:       "KeyPgUp",
	KeyPgDn:       "KeyPgDn",
	KeyArrowUp:    "KeyArrowUp",
	KeyArrowDown:  "KeyArrowDown",
	KeyArrowLeft:  "KeyArrowLeft",
	KeyArrowRight: "KeyArrowRight",
	KeyCtrlTilde:  "KeyCtrlTilde",
	KeyCtrlA:      "KeyCtrlA",
	KeyCtrlB:      "KeyCtrlB",
	KeyCtrlC:      "KeyCtrlC",
	KeyCtrlD:      "KeyCtrlD",
	KeyCtrlE:      "KeyCtrlE",
	KeyCtrlF:      "KeyCtrlF",
	KeyCtrlG:      "KeyCtrlG",
	KeyBackspace:  "KeyBackspace",
	KeyTab:        "KeyTab",
	KeyCtrlJ:      "KeyCtrlJ",
	KeyCtrlK:      "KeyCtrlK",
	KeyCtrlL:      "KeyCtrlL",
	KeyEnter:      "KeyEnter",
	KeyCtrlN:      "KeyCtrlN",
	KeyCtrlO:      "KeyCtrlO",
	KeyCtrlP:      "KeyCtrlP",
	KeyCtrlQ:      "KeyCtrlQ",
	KeyCtrlR:      "KeyCtrlR",
	KeyCtrlS:      "KeyCtrlS",
	KeyCtrlT:      "KeyCtrlT",
	KeyCtrlU:      "KeyCtrlU",
	KeyCtrlV:      "KeyCtrlV",
	KeyCtrlW:      "KeyCtrlW",
	KeyCtrlX:      "KeyCtrlX",
	KeyCtrlY:      "KeyCtrlY",
	KeyCtrlZ:      "KeyCtrlZ",
	KeyEsc:        "KeyEsc",
	KeyCtrl4:      "KeyCtrl4",
	KeyCtrl5:      "KeyCtrl5",
	KeyCtrl6:      "KeyCtrl6",
	KeyCtrl7:      "KeyCtrl7",
	KeySpace:      "KeySpace",
	KeyBackspace2: "KeyBackspace2",
}

// Printable characters, but worth having constants for them.
const (
	KeySpace = ' '
)

// Negative values for non-printable characters.
const (
	KeyF1 Key = -(iota + 1)
	KeyF2
	KeyF3
	KeyF4
	KeyF5
	KeyF6
	KeyF7
	KeyF8
	KeyF9
	KeyF10
	KeyF11
	KeyF12
	KeyInsert
	KeyDelete
	KeyHome
	KeyEnd
	KeyPgUp
	KeyPgDn
	KeyArrowUp
	KeyArrowDown
	KeyArrowLeft
	KeyArrowRight
	KeyCtrlTilde
	KeyCtrlA
	KeyCtrlB
	KeyCtrlC
	KeyCtrlD
	KeyCtrlE
	KeyCtrlF
	KeyCtrlG
	KeyBackspace
	KeyTab
	KeyCtrlJ
	KeyCtrlK
	KeyCtrlL
	KeyEnter
	KeyCtrlN
	KeyCtrlO
	KeyCtrlP
	KeyCtrlQ
	KeyCtrlR
	KeyCtrlS
	KeyCtrlT
	KeyCtrlU
	KeyCtrlV
	KeyCtrlW
	KeyCtrlX
	KeyCtrlY
	KeyCtrlZ
	KeyEsc
	KeyCtrl4
	KeyCtrl5
	KeyCtrl6
	KeyCtrl7
	KeyBackspace2
)

// Keys declared as duplicates by termbox.
const (
	KeyCtrl2          Key = KeyCtrlTilde
	KeyCtrlSpace      Key = KeyCtrlTilde
	KeyCtrlH          Key = KeyBackspace
	KeyCtrlI          Key = KeyTab
	KeyCtrlM          Key = KeyEnter
	KeyCtrlLsqBracket Key = KeyEsc
	KeyCtrl3          Key = KeyEsc
	KeyCtrlBackslash  Key = KeyCtrl4
	KeyCtrlRsqBracket Key = KeyCtrl5
	KeyCtrlSlash      Key = KeyCtrl7
	KeyCtrlUnderscore Key = KeyCtrl7
	KeyCtrl8          Key = KeyBackspace2
)
