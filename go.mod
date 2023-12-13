module github.com/multiverse-os/cli

go 1.19

require (
	github.com/multiverse-os/cli/data v0.1.0
	github.com/multiverse-os/cli/terminal/ansi v0.1.0
	github.com/multiverse-os/cli/terminal/loading v0.1.0
	github.com/multiverse-os/cli/terminal/text/banner v0.1.0
)

// THIS IS BECAUSE OF banner/big
require github.com/multiverse-os/banner v0.1.0 // indirect

replace (
	github.com/multiverse-os/cli/data => ./data
	github.com/multiverse-os/cli/terminal/ansi => ./terminal/ansi
	github.com/multiverse-os/cli/terminal/loading => ./terminal/loading
	github.com/multiverse-os/cli/terminal/text/banner => ./terminal/text/banner
)

// Where the fuck is this indirect coming form!!!
//github.com/multiverse-os/banner v0.1.0 // indirect
require golang.org/x/text v0.14.0 // indirect
