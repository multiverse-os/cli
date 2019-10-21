package main

import (
	"fmt"
	"os"

	log "github.com/multiverse-os/cli/log"
	color "github.com/multiverse-os/cli/text/ansi/color"
	style "github.com/multiverse-os/cli/text/ansi/style"
)

func main() {
	wd, _ := os.Getwd()
	log.Info("Print log to terminal (stdout)")
	log.Log(log.WARN, "Write WARN level log to file (working directory)").File(wd + "/debub.log")

}

func banner() {
	fmt.Println(style.Dim("(Minimalistic)") + color.Gray(" Log Example"))
	fmt.Println(style.Dim(color.White("==========================")))
}
