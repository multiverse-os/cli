package text

import (
	"fmt"

	color "github.com/multiverse-os/cli-framework/text/color"
)

func PrintBanner(appName, version string) {
	fmt.Println(color.Header(appName) + "  " + color.Light(version))
	fmt.Println(Repeat("=", 80))
}
