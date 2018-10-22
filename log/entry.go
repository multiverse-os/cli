package log

import (
	"fmt"
	"time"

	color "github.com/multiverse-os/cli-framework/text/color"
)

// TODO: Add variables that are being output
// TODO: Add hooks like after write, before write, etc
type Entry struct {
	Logger *Logger
	Time   time.Time
	Type   LogType
	Text   string
	errors []error
}

func NewLog(logger *Logger, logType LogType, text string) {
	logger.Log(logType, text)
}

func (self *Entry) Print() {
	fmt.Println(self.Type.FormattedString(true) + color.Blue("["+self.Time.String()+"]") + self.Text)
}

// TODO: Make a .String() and .JSON() function for writing to those
// and a .WriteToFile or AppendToFile function may be better way
// of handling structure
