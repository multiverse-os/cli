package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/multiverse-os/cli/framework/loaders/spinner"
	"github.com/multiverse-os/cli/framework/loaders/spinner/clock"
	"github.com/multiverse-os/cli/framework/loaders/spinner/dots"
	"github.com/multiverse-os/cli/framework/text"
)

var Rainbow = []string{text.LIGHT_RED, text.RED, text.LIGHT_RED,
	text.LIGHT_YELLOW, text.YELLOW, text.LIGHT_YELLOW, text.LIGHT_GREEN,
	text.GREEN, text.LIGHT_GREEN, text.LIGHT_CYAN, text.CYAN, text.LIGHT_CYAN,
	text.LIGHT_BLUE, text.BLUE, text.LIGHT_BLUE, text.LIGHT_MAGENTA,
	text.MAGENTA, text.LIGHT_MAGENTA}

var Blues = []string{text.LIGHT_BLUE, text.LIGHT_BLUE, text.BLUE, text.BLUE,
	text.BLUE, text.BLUE, text.BLUE, text.BLUE, text.BLUE, text.LIGHT_BLUE,
	text.LIGHT_BLUE, text.LIGHT_BLUE, text.WHITE, text.WHITE, text.LIGHT_CYAN,
	text.LIGHT_CYAN, text.LIGHT_CYAN, text.CYAN, text.LIGHT_CYAN,
	text.LIGHT_CYAN, text.LIGHT_CYAN}

func randomWait() {
	time.Sleep(time.Duration(rand.Intn(2)+2) * time.Second)
}

func main() {
	fmt.Println("Loader Example")
	fmt.Println("==============")
	RainbowSpinner()
	//MultiMessageSpinner()
}

// Options are defined using function chaining, and the available options
// are:
//   1) Speed(int), with aliases spinner.(Slowest,Slow,Normal,Fast,Fastest)
//   2) Messasge(string)
//   3) Palette([]string), expecting ansi colors to cycle through

func RainbowSpinner() {
	rainbowDots := spinner.New(dots.Animation).Message("Loading...").Speed(spinner.Normal).Palette(Rainbow).Start()
	randomWait()
	rainbowDots.Complete(text.Green("Loading Complete!"))
}

func MultiMessageSpinner() {
	s := spinner.New(clock.Animation).Message("Loading...").Speed(spinner.Normal)
	s.Start()
	randomWait()
	randomWait()

	// Currently there are only six (6) available spinner options
	// any pull requests for additional spinners will be reviewed
	// and accepted if the code is consistent with the existing
	// library codebase. Pull requests welcome.
	// Loader Animation Options:
	//    (1) circle.Animation   ["◐","◓", "◑", "◒"]
	//    (2) clock.Animation    ["🕐","🕑","🕒","🕓",...]
	//    (3) dots.Animation     ["⠋","⠙","⠹","⠸",...]
	//    (4) lines.Animation    ["-","\","|","/"]
	//    (5) moon.Animation     ["🌑","🌒","🌓","🌔",...]
	//    (6) triangle.Animation ["◢","◣","◤","◥"]

	// To provide a more interactive UI, messages, speed and even palette can
	// be updated while the spinner animation is active to provide updates
	// which lets the user know the program is not frozen.
	s.Message("Water, Dirt & Grass")
	randomWait()
	s.Message("Trees, Debris & Hideouts")
	randomWait()
	s.Message("Wildlife, Werewolves & Bandits")
	randomWait()
	s.Message("Sounds of wildlife & trees waving in the wind")
	randomWait()
	s.Message("Hiding treasure in the haunted woods...")
	randomWait()
	s.Complete("Completed")

}
