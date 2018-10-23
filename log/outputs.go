package log

import (
	"os"
	"sync"
)

type Outputs []LogOutput

func (self Outputs) Open() {
	for _, output := range self {
		output.Open()
	}
}

func (self Outputs) Close() {
	for _, output := range self {
		output.Close()
	}
}

func (self Outputs) Append(entry Entry) {
	for _, output := range self {
		output.Append(entry)
	}
}

// Default supported output types
///////////////////////////////////////////////////////////////////////////////
//    (1) File
//    (2) StdOut
type LogOutput interface {
	Open()
	Close()
	Append(entry Entry)
}

//
// STREAM: File I/O
///////////////////////////////////////////////////////////////////////////////
type LogFile struct {
	Path     string
	Filename string
	Mutex    sync.Mutex
	Data     *os.File
}

func (self *LogFile) Open() (err error) {
	defer self.Close()
	self.Mutex.Lock()
	self.Data, err = os.OpenFile(self.FilePath(), os.O_APPEND|os.O_WRONLY, 0660)
	self.Mutex.Unlock()
	return err
}

// TODO: This should be called in signal cancel and shutdown
func (self *LogFile) Close() {
	self.Mutex.Lock()
	self.Data.Close()
	self.Mutex.Unlock()
}

// TODO: We should support an array or map of io.writers, then we can cycle through and write to each one
// supporting stdout, logfile and so on, potentially even writing to held open html requests
func (self *LogFile) Append(logEntry Entry) {
	formattedMessage := logEntry.Message
	self.Mutex.Lock()
	_, err := self.Data.WriteString(formattedMessage)
	self.Mutex.Unlock()
	if err != nil {
		FatalError(err)
	}
}

//
// STREAM: OS stdout
///////////////////////////////////////////////////////////////////////////////
// There is two versions of this, (1) where we pull apart the "fmt" library
// and implement all that is needed to interact with StdOut the other, (2)
// is to just make it simple as possible leveraging the existing "fmt".
type StdOut struct{}

func (self *StdOut) Open()  {}
func (self *StdOut) Close() {}

func (self *StdOut) Append(logEntry Entry) {
	//logEntry.Print()
}
