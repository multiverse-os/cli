package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	ConfigDir     string
	User          *user.User
	IsRoot        bool
	Hostname, Cwd string
	IsLocalhost   bool
	LoadAvg       float32

	HasBattery         bool
	BatteryPercent     int
	BatteryDischarging bool
)

const (
	SysPowerPath = "/sys/class/power_supply"
)

// GetConfigDir finds the directory name for configuration files, saving it in
// ConfigDir.
func GetConfigDir() {
	ConfigDir = filepath.Join(os.Getenv("HOME"), ".config", "bprompt")
}

// GetUser finds the current user's details, saving them in User.
func GetUser() {
	var err error
	User, err = user.Current()
	if err != nil {
		CaptureError(err)
		User = &user.User{
			Uid:      "(err)",
			Gid:      "(err)",
			Username: "(err)",
			Name:     "(err)",
			HomeDir:  "/tmp",
		}
	}
	IsRoot = User.Uid == "0"
}

// GetHost finds the current hostname, and tests for the SSH_CONNECTION
// environment variable to determine whether or not this is a local or remote
// connection. Sets Hostname and IsLocalhost.
func GetHost() {
	var err error

	Hostname, err = os.Hostname()
	if err != nil {
		CaptureError(err)
		Hostname = "(err)"
	}

	IsLocalhost = (os.Getenv("SSH_CONNECTION") == "")
}

// GetLoadAverage finds the system's 1-minute load average, saving it in
// LoadAvg.
func GetLoadAverage() {
	LoadAvg = -1 // error condition

	b, err := ioutil.ReadFile("/proc/loadavg")
	if err != nil {
		CaptureError(err)
		return
	}

	p := bytes.IndexByte(b, ' ')
	if p == -1 {
		p = len(b)
	}

	var load float64
	if load, err = strconv.ParseFloat(string(b[:p]), 32); err != nil {
		CaptureError(err)
		return
	}

	LoadAvg = float32(load)
}

// GetCwd finds the current working directory, saving it in Cwd.
func GetCwd() {
	var err error
	Cwd, err = os.Getwd()
	if err != nil {
		Cwd = "/"
		CaptureError(err)
	}
}

// GetBattery finds the battery level, and whether or not we are on AC. Sets
// HasBattery, BatteryPercent and BatteryDischarging.
func GetBattery() {
	batPath := filepath.Join(SysPowerPath, "BAT0")
	if _, err := os.Stat(batPath); err != nil {
		batPath = filepath.Join(SysPowerPath, "BAT1")
		if _, err := os.Stat(batPath); err != nil {
			return
		}
	}
	HasBattery = true
	BatteryPercent = -1
	BatteryDischarging = true

	cap, err := ioutil.ReadFile(filepath.Join(batPath, "capacity"))
	if err != nil {
		CaptureError(err)
	} else {
		i, err := strconv.ParseUint(strings.TrimSpace(string(cap)),
			10, 32)
		if err != nil {
			CaptureError(fmt.Errorf("battery capacity: %v", err))
		} else {
			BatteryPercent = int(i)
			if BatteryPercent > 100 {
				BatteryPercent = 100
			}
		}
	}

	acPath := filepath.Join(SysPowerPath, "AC")
	if _, err := os.Stat(acPath); err != nil {
		acPath = filepath.Join(SysPowerPath, "ADP1")
	}
	ac, err := ioutil.ReadFile(filepath.Join(acPath, "online"))
	if err != nil {
		CaptureError(err)
	} else {
		t, err := strconv.ParseBool(strings.TrimSpace(string(ac)))
		if err != nil {
			CaptureError(fmt.Errorf("ac online: %v", err))
		} else {
			BatteryDischarging = !t
		}
	}
}
