package keyboard

import "testing"

func TestWatcher(t *testing.T) {
	m := NewWatcher()
	m.SetState(A, Down)
	m.SetState(ArrowLeft, Up)
	if !m.Down(A) {
		t.Fatal("expect keyboard.A in state keyboard.Down")
	}
	if !m.Up(ArrowLeft) {
		t.Fatal("expect keyboard.ArrowLeft in state keyboard.Up")
	}
	if !m.Up(Escape) {
		t.Fatal("expect keyboard.Esc in state keyboard.Up")
	}

	// Verify the state lookup table.
	want := map[Key]State{
		A:         Down,
		ArrowLeft: Up,
		Escape:    Up,
	}
	states := m.States()
	if len(states) != len(want) {
		t.Fatalf("got %d states, want %d\n", len(states), len(want))
	}
	for key, state := range states {
		wantState := want[key]
		if wantState != state {
			t.Fatalf("got %v=%v, want %v=%v\n", key, state, key, wantState)
		}
	}
}
