package log

import (
	"fmt"
	"os"
	"sync"
)

type Outputs []LogOutput

type Output int

const (
	STDOUT Output = iota
	FILE
)

type LogOutput interface {
	Open() error
	Close()

	Append(entry Entry)
}

// STREAM: File I/O
///////////////////////////////////////////////////////////////////////////////
type LogFile struct {
	format Format
	mutex  sync.Mutex
	path   string
	data   *os.File
}

func (self *LogFile) Open() (err error) {
	defer self.Close()
	self.mutex.Lock()
	self.data, err = os.OpenFile(self.path, os.O_APPEND|os.O_WRONLY, 0660)
	self.mutex.Unlock()
	return err
}

// For a clean shutdown, this should be called in shutdown/signal
func (self *LogFile) Close() {
	self.mutex.Lock()
	self.data.Close()
	self.mutex.Unlock()
}

func (self *LogFile) Append(entry Entry) {
	self.mutex.Lock()
	_, err := self.data.WriteString(entry.FormattedOutput())
	self.mutex.Unlock()
	if err != nil {
		FatalError(err)
	}
}

// STREAM: OS stdout
///////////////////////////////////////////////////////////////////////////////
// There is two versions of this, (1) where we pull apart the "fmt" library
// and implement all that is needed to interact with StdOut the other, (2)
// is to just make it simple as possible leveraging the existing "fmt".
type StdOut struct {
	format Format
	mutex  sync.Mutex
}

func (self *StdOut) Open() error { return nil }
func (self *StdOut) Close()      {}

func (self *StdOut) Append(entry Entry) {
	fmt.Println(entry.Format(self.format))
}
