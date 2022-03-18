package parse_test

// NOTE: Because its external we have to import the library we are testing
import (
  "testing"

  cli "github.com/multiverse-os/cli"
)

type testArg []string

func testArgs() []testArg {
  return  []testArg{
    testArg{"h"},
    testArg{"help"},
    testArg{"-h"},
    testArg{"-help"},
    testArg{"v"},
    testArg{"version"},
    testArg{"-v"},
    testArg{"-version"},
  }
}

// TODO: Will need to define different cli objects using New() 

// TODO: Some obvio problems with these tests lol, wtf if you think this is
// normal omgffg
func Test_Parse(t *testing.T) {
  for _, arg := range testArgs() {
    cli := cli.New().Parse(arg)
    if cli == nil {
      t.Errorf("cli failed to create with empty cli.App, returned nil cli.CLI object")
    }
  }
}
