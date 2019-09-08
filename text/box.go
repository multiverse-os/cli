package text

import (
	"fmt"
)

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
