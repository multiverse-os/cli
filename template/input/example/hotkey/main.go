package main

import (
	"fmt"
	"os"

	signals "github.com/multiverse-os/cli/framework/service/signal"
	keyboard "github.com/multiverse-os/cli/framework/uidevice/keyboard"
)

func main() {
	fmt.Println("hotkey watcher")
	fmt.Println("==============")
	fmt.Println("Watching for keyboard input to react to it")

	m := keyboard.NewWatcher()
	m.SetState(keyboard.A, keyboard.Down)
	m.SetState(keyboard.ArrowLeft, keyboard.Up)
	if !m.Down(keyboard.A) {
		fmt.Println("[error] expected keyboard.A in state keyboard.Down")
	} else {
		fmt.Println("!!m.Down(keyboard.A)")
	}
	if !m.Up(keyboard.ArrowLeft) {
		fmt.Println("[error] expected keyboard.ArrowLeft in state keyboard.Up")
	} else {
		fmt.Println("!!m.Down(keyboard.ArrowLeft)")
	}
	if !m.Up(keyboard.Escape) {
		fmt.Println("[error] expected keyboard.Esc in state keyboard.Up")
	} else {
		fmt.Println("!!m.Down(keyboard.Escape)")
	}

	// Verify the state lookup table.
	want := map[keyboard.Key]keyboard.State{
		keyboard.A:         keyboard.Down,
		keyboard.ArrowLeft: keyboard.Up,
		keyboard.Escape:    keyboard.Up,
	}
	states := m.States()
	if len(states) != len(want) {
		fmt.Println("got %d states, want %d\n", len(states), len(want))
	}
	for key, state := range states {
		wantState := want[key]
		if wantState != state {
			fmt.Println("got %v=%v, want %v=%v\n", key, state, key, wantState)
		}
	}
	fmt.Println("holding open, use CTRL+C to shutdown")
	signals.ShutdownHandler(func(s os.Signal) {
		fmt.Println("[CTRL+C] user terminated the application, shutting down received a singal:", s)
		os.Exit(1)
	})
	for {

	}
}
