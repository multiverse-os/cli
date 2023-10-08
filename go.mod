module github.com/multiverse-os/cli

go 1.19

require (
	github.com/multiverse-os/ansi v0.0.0-20230122075550-10efed87b9d4
	github.com/multiverse-os/banner v0.1.0
	github.com/multiverse-os/color v0.1.0
	github.com/multiverse-os/loading v0.1.0
)

replace github.com/multiverse-os/terminal/ansi/color => github.com/multiverse-os/color v0.1.0

require golang.org/x/text v0.13.0 // indirect
