package text

type BoxType int

const (
	DefaultBox BoxType = 0
	HashBox            = 1
	EqualsBox          = 2
)

type BoxBorder int

const (
	TopBorder    BoxBorder = 0
	RightBorder            = 1
	BottomBorder           = 2
	LeftBorder             = 3
)

func Border(text string, border BoxBorder) string {
	switch border {
	case TopBorder:
		// P
	case RightBorder:
	case BottomBorder:
	case LeftBorder:
	}
	return text
}

func Box(text string) {
	l := len(text)
	fmt.Printf("╭")
	for i := 0; i < l+2; i++ {
		fmt.Printf("─")
	}
	fmt.Printf("╮\n")
	fmt.Printf("│ %v │\n", text)
	fmt.Printf("╰")
	for i := 0; i < l+2; i++ {
		fmt.Printf("─")
	}
	fmt.Printf("╯\n")
}
