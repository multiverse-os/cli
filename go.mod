module github.com/multiverse-os/cli

go 1.19

replace (
	github.com/multiverse-os/cli/data => github.com/multiverse-os/data v0.1.0
	github.com/multiverse-os/cli/terminal/ansi => github.com/multiverse-os/ansi v0.1.0
	github.com/multiverse-os/cli/terminal/loading => github.com/multiverse-os/loading v0.1.0
	github.com/multiverse-os/cli/terminal/text => github.com/multiverse-os/text v0.1.0
)

require golang.org/x/text v0.13.0 // indirect
