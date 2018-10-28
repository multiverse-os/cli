package log

import (
	"fmt"
	"os"
	"sync"
)

type Outputs []LogOutput

type Output int

const (
	DEFAULT_LOG Output = iota
	USER_LOG
	OS_LOG
	TERMINAL
	FILE
	//SYSLOG     // Not yet supported
	//JOURNALCTL // Not yet supported
	//HTTP       // Not yet supported (use case: doing HTTP requests against REST APIs to send logs to chat servers)
)

// Output Aliases
const (
	STDOUT = TERMINAL
)

type LogOutput interface {
	Open() error
	Close() error
	Append(entry Entry) error
}

// OUTPUT: File I/O
///////////////////////////////////////////////////////////////////////////////
type LogFile struct {
	format  Format
	mutex   sync.Mutex
	path    string
	data    *os.File
	rotate  bool
	maxSize int // kb
}

func (self *LogFile) Open() (err error) {
	defer self.Close()
	self.mutex.Lock()
	self.data, err = os.OpenFile(self.path, os.O_APPEND|os.O_WRONLY, 0660)
	self.mutex.Unlock()
	return err
}

// For a clean shutdown, this should be called in shutdown/signal
func (self *LogFile) Close() error {
	self.mutex.Lock()
	self.data.Close()
	self.mutex.Unlock()
	return nil
}

func (self *LogFile) Append(entry Entry) (err error) {
	self.mutex.Lock()
	_, err = self.data.WriteString(entry.String())
	self.mutex.Unlock()
	return err
}

// OUTPUT: OS/User (Default) Log File
///////////////////////////////////////////////////////////////////////////////

// Just use above File I/O, or build custom UserLogFile and OSLogFile types?

// [OS]  [Log File] [Path:'/var/log/APP_NAME/APP_NAME.log\']
// [User][Log File] [Path:'/home/USER_NAME/.local/share/APP_NAME/APP_NAME.log')

// OUTPUT: Terminal Output (stdout)
///////////////////////////////////////////////////////////////////////////////
type Terminal struct{ format Format }

func (self *Terminal) Open() error  { return nil }
func (self *Terminal) Close() error { return nil }
func (self *Terminal) Append(entry Entry) error {
	fmt.Println(entry.String())
	return nil
}

// OUTPUT: Syslog Output
///////////////////////////////////////////////////////////////////////////////
