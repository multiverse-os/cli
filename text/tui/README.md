# TUI
Need to pull simplest elements out to make a usable TUI that is not too bloated. 

I like the following API:

```
type UI interface {
	Say(message string, args ...interface{})
	Warn(message string, args ...interface{})
	Ask(prompt string, args ...interface{}) (answer string)
	AskForPassword(prompt string, args ...interface{}) (answer string)
	Confirm(message string, args ...interface{}) bool
	Failed(message string, args ...interface{})
	Wait(duration time.Duration)
	EmptyLine()
	Ok()
}

type terminalUI struct {
	stdin   io.Reader
	printer Printer
}

func NewUI(r io.Reader, printer Printer) UI {
	return &terminalUI{
		stdin:   r,
		printer: printer,
	}
}

```

[Source](https://github.com/mcastilho/terminal/blob/master/ui.go)
