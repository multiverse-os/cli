module github.com/multiverse-os/cli

go 1.19

replace (
	github.com/multiverse-os/banner v0.1.0 => github.com/multiverse-os/cli/terminal/text/banner v0.1.0

	github.com/multiverse-os/cli/data v0.1.0 => ./data
	github.com/multiverse-os/cli/terminal/ansi v0.1.0 => ./terminal/ansi
	github.com/multiverse-os/cli/terminal/loading v0.1.0 => ./terminal/loading
	github.com/multiverse-os/cli/terminal/text/banner v0.1.0 => ./terminal/text/banner
)

require (
	github.com/multiverse-os/cli/data v0.1.0
	github.com/multiverse-os/cli/terminal/ansi v0.1.0
	github.com/multiverse-os/cli/terminal/loading v0.1.0
	github.com/multiverse-os/cli/terminal/text/banner v0.1.0
)

//exclude github.com/multiverse-os/banner v0.1.0 // indirect

//exclude github.com/multiverse-os/banner v0.1.0

//github.com/multiverse-os/banner v0.1.0 // indirect
require golang.org/x/text v0.14.0 // indirect
