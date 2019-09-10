package main

import (
	"fmt"

	color "github.com/multiverse-os/cli/text/ansi/color"
	style "github.com/multiverse-os/cli/text/ansi/style"
)

func main() {
	// TODO: Fix the order
	// Primary 8 ANSI Colors
	fmt.Printf(" +=====================+\n")
	fmt.Printf(" | Primary ANSI Colors |\n")
	fmt.Printf(" +=====================+\n")
	fmt.Printf(" | %s   | %s   |  \n", color.Black("Black"), color.BlackBg("BlackBg"))
	fmt.Printf(" | %s  | %s  | \n", color.Maroon("Maroon"), color.MaroonBg("MaroonBg"))
	fmt.Printf(" | %s   | %s   |  \n", color.Green("Green"), color.GreenBg("GreenBg"))
	fmt.Printf(" | %s   | %s   | \n", color.Olive("Olive"), color.OliveBg("OliveBg"))
	fmt.Printf(" | %s    | %s    |  \n", color.Blue("Blue"), color.BlueBg("BlueBg"))
	fmt.Printf(" | %s | %s |  \n", color.Magenta("Magenta"), color.MagentaBg("MagentaBg"))
	fmt.Printf(" | %s    | %s    |  \n", color.Cyan("Cyan"), color.CyanBg("CyanBg"))
	fmt.Printf(" | %s  | %s  | \n", color.Silver("Silver"), color.SilverBg("SilverBg"))
	fmt.Printf(" +---------+-----------+\n\n")

	// Secondary 8 ANSI Colors
	fmt.Printf(" +=====================+\n")
	fmt.Printf(" |Secondary ANSI Colors|\n")
	fmt.Printf(" +=====================+\n")
	fmt.Printf(" | %s    | %s    | \n", color.Gray("Gray"), color.GrayBg("GrayBg"))
	fmt.Printf(" | %s     | %s     |  \n", color.Red("Red"), color.RedBg("RedBg"))
	fmt.Printf(" | %s    | %s    | \n", color.Lime("Lime"), color.LimeBg("LimeBg"))
	fmt.Printf(" | %s  | %s  |  \n", color.Yellow("Yellow"), color.YellowBg("YellowBg"))
	fmt.Printf(" | %s | %s | \n", color.SkyBlue("SkyBlue"), color.SkyBlueBg("SkyBlueBg"))
	fmt.Printf(" | %s | %s | \n", color.Fuchsia("Fuchsia"), color.FuchsiaBg("FuchsiaBg"))
	fmt.Printf(" | %s    | %s    | \n", color.Aqua("Aqua"), color.AquaBg("AquaBg"))
	fmt.Printf(" | %s   | %s   |  \n", color.White("White"), color.WhiteBg("WhiteBg"))
	fmt.Printf(" +---------+-----------+\n\n")

	fmt.Printf(" ANSI Style Options\n")
	fmt.Printf(" +=====================+\n")
	fmt.Printf("  %s  \n", style.Bold("Bold"))
	fmt.Printf("  %s  \n", style.Dim("Dim"))
	fmt.Printf("  %s  \n", style.Italic("Italic"))
	fmt.Printf("  %s  \n", style.Underline("Underline"))
	fmt.Printf("  %s  \n", style.SlowBlink("Slow Blink"))
	fmt.Printf("  %s  \n", style.FastBlink("Fast Blink"))
	fmt.Printf("  %s  \n", style.Inverse("Inverse"))
	fmt.Printf("  %s  \n", style.Conceal("Conceal"))
	fmt.Printf("  %s  \n", style.Strikethrough("Strikethough"))
	fmt.Printf("  %s  \n", style.Framed("Framed"))
	fmt.Printf("  %s  \n", style.Encircle("Encircle"))
	fmt.Printf("  %s  \n", style.Overline("Overline"))
}
