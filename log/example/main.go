package main

import (
	"fmt"

	log "github.com/multiverse-os/cli-framework/log"
	"github.com/multiverse-os/cli-framework/os/fs"
	text "github.com/multiverse-os/cli-framework/text"
)

func main() {
	banner()
	// Ephemeral Log Interface Example
	log.Info("Print INFO level log to terminal (stdout)")
	log.Log(log.WARN, "Write WARN level log to file (working directory)").File(fs.WorkingDirectory + "/debub.log")

}

func banner() {
	fmt.Println(text.Light("(Minimalistic)") + text.Gray(" Log Example"))
	fmt.Println(text.Light(text.White("==========================")))
}
