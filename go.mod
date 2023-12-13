module github.com/multiverse-os/cli

go 1.19

//require github.com/multiverse-os/banner v0.1.0
//
//require (
//	github.com/multiverse-os/cli/terminal/ansi v0.1.0
//	github.com/multiverse-os/cli/terminal/loading v0.0.0-00010101000000-000000000000
//)

require (
	github.com/multiverse-os/banner v0.1.0
	github.com/multiverse-os/cli/data v0.1.0
	github.com/multiverse-os/cli/terminal/ansi v0.1.0
	github.com/multiverse-os/cli/terminal/loading v0.1.0
)

require golang.org/x/text v0.13.0 // indirect

//replace github.com/multiverse-os/cli/terminal/loading => github.com/multiverse-os/loading v0.1.0

//replace (
//	github.com/multiverse-os/ansi => github.com/multiverse-os/cli/terminal/ansi v0.1.0
//	github.com/multiverse-os/loading => github.com/multiverse-os/cli/terminal/loading v0.1.0
//)

replace (
	github.com/multiverse-os/cli/data => ./data
	github.com/multiverse-os/cli/terminal/ansi => ./terminal/ansi
	github.com/multiverse-os/cli/terminal/loading => ./terminal/loading
	github.com/multiverse-os/cli/terminal/text => ./terminal/text
	github.com/multiverse-os/cli/terminal/text/banner => ./terminal/text/banner
)
