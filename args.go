package cli

import (
	"fmt"
)

type Argument int

const (
	Add Argument = iota
	Push
	Commit
)

func (self Argument) Marshal(arg string) {
	switch arg {
	case Add.String():
		fmt.Println("do thing")
	}
}

//func isShort(arg string) bool {
//	return strings.HasPrefix(arg, "-") && !strings.HasPrefix(arg, "--") && len(arg) >= 2
//}
//
//func isLong(arg string) bool {
//	return strings.HasPrefix(arg, "--") && len(arg) >= 3
//}

// It use for parsing the boolean of flag or environnement
// TODO: Uhh just use ToLower to cut down some of this
//func toBool(value string) bool {
//	switch value {
//	case "true", "True", "TRUE", "1", "":
//		return true
//	case "false", "False", "FALSE", "0":
//		return false
//	default:
//		ErrLog.Println(Cerror{
//			Reason: E_SUPBTOBOOL,
//			Str:    value,
//		})
//		return true
//	}
//}
