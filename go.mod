module github.com/multiverse-os/cli

go 1.19

replace github.com/multiverse-os/cli/terminal/ansi => github.com/multiverse-os/ansi v0.0.0-20230212053502-2711fc61f14d

require (
	github.com/multiverse-os/banner v0.0.0-20230122081958-39bc0e2a3c54
	github.com/multiverse-os/cli/terminal/ansi v0.0.0-00010101000000-000000000000
	github.com/multiverse-os/loading v0.0.0-20230205140225-67dcaf84ca47
)

require github.com/multiverse-os/ansi v0.0.0-20230122075550-10efed87b9d4 // indirect
