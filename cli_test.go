package cli_test

import (
  "testing"

  cli "github.com/multiverse-os/cli"

)


func Test_New(t *testing.T) {

  // TODO: Empty should actually work, instead of failing it should fallback
  cli := cli.New(cli.App{
     
  })

  if cli == nil {
    t.Errorf("cli failed to create with empty cli.App, returned nil cli.CLI object")
  }
}


