package text

import (
	"fmt"
)

type BoxType int

const (
	DefaultBox BoxType = iota
	HashBox
	EqualsBox
)

type BoxBorder int

const (
	TopBorder BoxBorder = iota
	RightBorder
	BottomBorder
	LeftBorder
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
