package keyboard

import (
	"bytes"
	"fmt"
	"sync"
)

// Watcher watches the state of various keyboard keys.
type Watcher struct {
	access    sync.RWMutex
	states    map[Key]State
	rawStates map[uint64]State
}

// String returns a multi-line string representation of this keyboard watcher
// and it's associated states (but not raw ones).
func (w *Watcher) String() string {
	bb := new(bytes.Buffer)
	fmt.Fprintf(bb, "keyboard.Watcher(\n")
	for k, s := range w.States() {
		fmt.Fprintf(bb, "\t%v: %v,\n", k, s)
	}
	fmt.Fprintf(bb, ")")
	return bb.String()
}

// SetState specifies the current state of the specified key.
func (w *Watcher) SetState(k Key, s State) {
	w.access.Lock()
	defer w.access.Unlock()

	w.states[k] = s
}

// States returns an copy of the internal key state map used by this watcher.
func (w *Watcher) States() map[Key]State {
	w.access.RLock()
	defer w.access.RUnlock()

	cpy := make(map[Key]State, len(w.states))
	for key, state := range w.states {
		cpy[key] = state
	}
	return cpy
}

// EachState calls f with each known key to this watcher and it's current key
// state. It does so until the function returns false or there are no more keys
// known to the watcher.
func (w *Watcher) EachState(f func(k Key, s State) bool) {
	w.access.RLock()
	defer w.access.RUnlock()

	for key, state := range w.states {
		// Call the function without the lock being held, so they can access
		// methods on this watcher still.
		w.access.RUnlock()
		cont := f(key, state)
		w.access.RLock()

		if !cont {
			return
		}
	}
}

// State returns the current state of the specified key.
func (w *Watcher) State(k Key) State {
	w.access.Lock()
	defer w.access.Unlock()

	state, ok := w.states[k]
	if !ok {
		w.states[k] = Up
		return Up
	}
	return state
}

// Down tells whether the specified key is currently in the down state.
func (w *Watcher) Down(k Key) bool {
	return w.State(k) == Down
}

// Up tells whether the specified key is currently in the up state.
func (w *Watcher) Up(k Key) bool {
	return w.State(k) == Up
}

// SetRawState specifies the current state of the specified raw key value.
func (w *Watcher) SetRawState(raw uint64, s State) {
	w.access.Lock()
	defer w.access.Unlock()

	w.rawStates[raw] = s
}

// RawStates returns an copy of the internal raw key state map used by this
// watcher.
func (w *Watcher) RawStates() map[uint64]State {
	w.access.RLock()
	defer w.access.RUnlock()

	cpy := make(map[uint64]State, len(w.rawStates))
	for raw, state := range w.rawStates {
		cpy[raw] = state
	}
	return cpy
}

// RawState returns the current state of the specified raw key value.
func (w *Watcher) RawState(raw uint64) State {
	w.access.Lock()
	defer w.access.Unlock()

	state, ok := w.rawStates[raw]
	if !ok {
		w.rawStates[raw] = Up
		return Up
	}
	return state
}

// RawDown tells whether the specified raw key value is currently in the down
// state.
func (w *Watcher) RawDown(raw uint64) bool {
	return w.RawState(raw) == Down
}

// RawUp tells whether the specified raw key value is currently in the up
// state.
func (w *Watcher) RawUp(raw uint64) bool {
	return w.RawState(raw) == Up
}

// NewWatcher returns a new, initialized, watcher.
func NewWatcher() *Watcher {
	w := new(Watcher)
	w.states = make(map[Key]State)
	w.rawStates = make(map[uint64]State)
	return w
}
