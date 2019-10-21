package terminal

import (
	"fmt"
	"os/exec"
)

type Shell struct {
}

func Bash(script string) (string, error) {
	out, err := exec.Command("/bin/bash", script).Output()
	return string(out), err
}

func Sh(script string) (string, error) {
	out, err := exec.Command("/bin/sh", script).Output()
	return string(out), err
}

func ParseShell() ([]string, error) {
	if path, err := exec.LookPath("bash"); err == nil {
		return []string{path, "-c"}, nil
	}
	if path, err := exec.LookPath("sh"); err == nil {
		return []string{path, "-c"}, nil
	}
	return nil, fmt.Errorf("Could not find bash or sh on path.")
}
