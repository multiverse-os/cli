package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

const (
	TimeFormat = "Mon Jan _2 15:04:05 MST"
)

var (
	Width int

	ExitCode    = flag.String("exitCode", "0", "exit code of last process")
	ScreenWidth = flag.String("screenWidth", "", "screen width in chars")
)

func main() {
	flag.Parse()

	GetConfigDir()
	GetUser()
	GetHost()
	GetLoadAverage()
	GetCwd()
	GetBattery()

	GetWidth()

	firstLine := []*RoundBoxInfo{
		Time(),
		Who(),
		RoundBox(""), // placeholder for truncated directory
	}
	if HasBattery {
		firstLine = append(firstLine, Battery())
	}
	firstLine = append(firstLine,
		LoadAverage(),
		CommandStatus(),
	)

	path := FitPath(Cwd, RemainingWidth(FirstLine, firstLine))
	firstLine[2] = RoundBox(path)

	secondLine := []*RoundBoxInfo{
		GitBox(),
		DirnameBox(),
	}

	var b bytes.Buffer
	NewlineIfNecessary(&b)
	b.WriteString(strings.Repeat(" ", Width))
	PrintLine(&b, FirstLine, firstLine)
	b.WriteRune('\n')
	PrintLine(&b, SecondLine, secondLine)

	if IsRoot {
		SetColour(&b, "31")
		b.WriteString(" # ")
	} else {
		SetColour(&b, "32")
		b.WriteString(" $ ")
	}
	SetColour(&b, "")

	os.Stdout.Write(b.Bytes())
}

// GetWidth interprets the screen width command line parameter, saving it in
// Width.
func GetWidth() {
	Width = 80 // failsafe value

	if *ScreenWidth == "" {
		CaptureError(errors.New("-screenWidth not specified"))
	}

	w, err := strconv.Atoi(*ScreenWidth)
	if err != nil {
		CaptureError(err)
		return
	}
	if w < 2 {
		CaptureError(fmt.Errorf("width too small: %d", w))
		return
	}

	Width = w
}

// NewlineIfNecessary determines if a newline is required, and arranges for it
// to be present if so. This is actually done by printing out a line's worth of
// filler characters, then resetting the cursor position to the left-hand side.
// If no newline was required, then the fillers will be overwritten by our
// normal output; otherwise, the fillers on the starting line (with partial
// output from the previous command) will be left in place.
func NewlineIfNecessary(w TerminalWriter) {
	SetColour(w, "35")
	w.WriteString(strings.Repeat("·", Width))
	w.Write([]byte(BashNPStart + CSI + "G" + BashNPEnd))
	SetColour(w, "")
}

// Time returns a RoundBox formatted with the system time.
func Time() *RoundBoxInfo {
	t := time.Now().Format(TimeFormat)
	printUTC := strings.HasSuffix(t, "GMT")

	var b bytes.Buffer
	SetColour(&b, "36")
	if printUTC {
		b.WriteString(t[:len(t)-3])
		b.Write([]byte("UTC"))
	} else {
		b.WriteString(t)
	}
	return RoundBox(b.String())
}

// Who returns a RoundBox formatted with the username and hostname, with
// highlighting for root/remote systems.
func Who() *RoundBoxInfo {
	var (
		b           bytes.Buffer
		leftColour  = 7
		rightColour = 7
	)

	// write username, draw attention to root
	if IsRoot {
		leftColour = 1
		SetColour(&b, "1;37;41")
	}
	b.WriteString(User.Username)

	// write separator char
	if IsRoot || !IsLocalhost {
		SetColour(&b, "0;34;41")
	}
	b.WriteRune('@')

	if hostInfo := LoadHostInfo(); hostInfo != nil {
		b.Write(hostInfo)
	} else {
		// write hostname, draw attention to remote
		if !IsLocalhost {
			rightColour = 1
			SetColour(&b, "1;37;41")
		} else {
			SetColour(&b, "0;47")
		}
		b.WriteString(Hostname)
	}

	r := RoundBox(b.String())
	r.SetColour(leftColour, rightColour)
	return r
}

// Battery returns a RoundBox displaying the current battery level and charge
// status.
func Battery() *RoundBoxInfo {
	var b bytes.Buffer
	switch {
	case BatteryPercent > 80: // green
		SetColour(&b, "32")
	case BatteryPercent > 40: // yellow
		SetColour(&b, "33")
	default: // red
		SetColour(&b, "31")
	}
	fmt.Fprintf(&b, "%d%%", BatteryPercent)

	if BatteryDischarging {
		SetColour(&b, "31") // red
		b.WriteRune('▽')
	} else {
		SetColour(&b, "33") // yellow
		b.WriteRune('⚡')    // \u26A1 (high voltage sign)
	}

	return RoundBox(b.String())
}

// LoadAverage returns a RoundBox displaying the colour-coded load average.
func LoadAverage() *RoundBoxInfo {
	var b bytes.Buffer
	switch {
	case LoadAvg < 0.2: // green
		SetColour(&b, "32")
	case LoadAvg < 2.0: // yellow
		SetColour(&b, "33")
	default: // red
		SetColour(&b, "31")
	}

	fmt.Fprintf(&b, "%.2f", LoadAvg)

	return RoundBox(b.String())
}

// CommandStatus returns a roundbox signifying whether the last command
// succeeded or failed.
func CommandStatus() *RoundBoxInfo {
	var b bytes.Buffer
	if *ExitCode != "0" {
		SetRGB(&b, RGB{255, 0, 0}, RGBUnset)
		b.WriteRune('▼')
	} else {
		SetRGB(&b, RGB{0, 255, 0}, RGBUnset)
		b.WriteRune('▲')
	}
	return RoundBox(b.String())
}

// DirnameBox returns a box with the current directory name.
func DirnameBox() *RoundBoxInfo {
	var d string
	if Cwd == os.Getenv("HOME") {
		d = "~"
	} else {
		d = filepath.Base(Cwd)
	}

	var b bytes.Buffer
	SetColour(&b, "1;37")
	b.WriteString(d)
	r := RoundBox(b.String())
	r.SetColour(4, 4)
	return r
}

func FitPath(p string, max int) string {
	if utf8.RuneCountInString(p) <= max {
		return p
	}

	components := strings.Split(p, "/")
	for i := 1; i < len(components); i++ {
		components[i-1] = "…/"
		p := filepath.Join(components[i-1:]...)
		if len(p) <= max {
			return p
		}
	}
	return "…"
}
