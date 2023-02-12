module github.com/multiverse-os/cli

go 1.19

replace (
  github.com/multiverse-os/cli/terminal/ansi    => github.com/multiverse-os/ansi    latest
  github.com/multiverse-os/cli/terminal/loading => github.com/multiverse-os/loading latest
  github.com/multiverse-os/cli/terminal/text    => github.com/multiverse-os/text    latest
)

