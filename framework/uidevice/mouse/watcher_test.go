package mouse

import "testing"

var wantStr = `mouse.Watcher(
	One: Down,
	Two: Up,
	Button(255): Down,
)`

func TestWatcher(t *testing.T) {
	m := NewWatcher()
	m.SetState(Left, Down)
	m.SetState(Right, Up)
	if !m.Down(Left) {
		t.Fatal("expect mouse.Left in state mouse.Down")
	}
	if !m.Up(Right) {
		t.Fatal("expect mouse.Right in state mouse.Up")
	}
	if !m.Up(Wheel) {
		t.Fatal("expect mouse.Wheel in state mouse.Up")
	}

	// Verify the state lookup table.
	want := map[Button]State{
		Left:  Down,
		Right: Up,
		Wheel: InvalidState,
	}
	states := m.States()
	if len(states) != 8 {
		t.Fatalf("got %d states, want 8\n", len(states))
	}
	for b, s := range states {
		wantState := want[Button(b)]
		if wantState != s {
			t.Fatalf("got %v=%v, want %v=%v\n", Button(b), s, Button(b), wantState)
		}
	}

	// Verify that expansion on the lookup table works OK.
	m.SetState(255, Down)
	got := m.State(255)
	if got != Down {
		t.Fatalf("Wanted Button(255) == Down, got Button(255) ==", got)
	}

	if m.String() != wantStr {
		t.Fatal("Watcher.String returned invalid string.")
		t.Logf("%q\n", m)
	}
}
