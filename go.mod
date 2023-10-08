module github.com/multiverse-os/cli

go 1.19

require (
	github.com/multiverse-os/ansi v0.1.0
	github.com/multiverse-os/banner v0.1.0
	github.com/multiverse-os/color v0.1.0
)

replace (
	github.com/multiverse-os/terminal/ansi/color => github.com/multiverse-os/color v0.1.0
	github.com/multiverse-os/terminal/loading => ./terminal/loading
// github.com/multiverse-os/terminal/loading => github.com/multiverse-os/loading v0.1.0
)

require (
	github.com/multiverse-os/loading v0.1.0 // indirect
	golang.org/x/text v0.13.0 // indirect
)
