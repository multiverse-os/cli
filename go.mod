module github.com/multiverse-os/cli

go 1.19

replace (
	github.com/multiverse-os/cli/terminal/ansi => ./terminal/ansi
	github.com/multiverse-os/cli/terminal/banner => ./terminal/banner
	github.com/multiverse-os/cli/terminal/loading => ./terminal/loading
	github.com/multiverse-os/cli/terminal/text => ./terminal/text
	github.com/multiverse-os/loading => github.com/multiverse-os/cli/terminal/loading v0.0.0-00010101000000-000000000000
// TODO: If I include loading as a submodule I get issues
)

require (
	github.com/multiverse-os/banner v0.0.0-20231003171846-79934d6d30a0
	github.com/multiverse-os/cli/terminal/ansi v0.0.0-00010101000000-000000000000
)

require github.com/multiverse-os/ansi v0.0.0-20230122075550-10efed87b9d4 // indirect
