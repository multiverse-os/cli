package main

import (
	"errors"
	"fmt"
	"os"

	cli "github.com/multiverse-os/cli-framework"
	log "github.com/multiverse-os/cli-framework/log"
)

func main() {
	cmd := cli.New(&cli.CLI{
		Name: "example",
		//Version: cli.Version{Major: 0, Minor: 1, Patch: 1},
		Usage: "make an explosive entrance",
		DefaultAction: func(c *cli.Context) error {
			fmt.Println("Example output for an action (or command)!")
			fmt.Println("version is: ", c.CLI.Version.ANSIFormattedString())
			c.CLI.Logger.Info("This should log to both stdout and file")
			log.Log(log.INFO, "Test info JSON log").PrettyJSON().StdOut()
			log.Log(log.INFO, "Test info JSON log").Format(log.JSON).File("./test.log")
			log.Log(log.INFO, "Test JSON with values").WithValue("test", "value").JSON().StdOut()
			log.Log(log.DEBUG, "Lying for good things isnt bad.").WithValue("test", 10).StdOut()
			log.Log(log.DEBUG, "Lying for good things isnt bad.").WithValue("pi", 3.14).WithError(errors.New("Test Error")).StdOut()
			log.Log(log.DEBUG, "Lying for good things isnt bad.").WithValue("testBoolean", true).WithErrors([]error{errors.New("Test Error"), errors.New("Another test error")}).StdOut()
			return nil
		},
	})

	cmd.Run(os.Args)
}
