module github.com/multiverse-os/cli

go 1.19

replace (
	github.com/multiverse-os/cli/data => ./data
	github.com/multiverse-os/cli/terminal/ansi => ./terminal/ansi
	github.com/multiverse-os/cli/terminal/loading => ./terminal/loading
	github.com/multiverse-os/cli/terminal/text => ./terminal/text
	github.com/multiverse-os/cli/terminal/text/banner => ./terminal/text/banner
)

require (
	github.com/multiverse-os/cli/data v0.1.0
	github.com/multiverse-os/cli/terminal/ansi v0.1.0
	github.com/multiverse-os/cli/terminal/loading v0.1.0
	github.com/multiverse-os/cli/terminal/text/banner v0.1.0
)

require (
	github.com/multiverse-os/banner v0.0.0-20231006133835-80f8c892b073 // indirect
	golang.org/x/text v0.13.0 // indirect
)
