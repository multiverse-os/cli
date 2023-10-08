package cli

import (
	loading "github.com/multiverse-os/cli/terminal/loading"
)

func (c CLI) LoadingBar(animationFrames []string) *loading.Bar {
	return loading.NewBar(animationFrames)
}

func (c CLI) Spinner(animationFrames []string) *loading.Spinner {
	return loading.NewSpinner(animationFrames)
}
