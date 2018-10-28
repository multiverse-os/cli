package text

import (
	"fmt"
)

func PrintBanner(appName, version string) {
	fmt.Println(Header(appName) + "  " + Light(version))
	fmt.Println(Repeat("=", 80))
}
