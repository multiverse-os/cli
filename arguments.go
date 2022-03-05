package cli

//import (
//  "strings"
//)

type ArgumentType int

const (
  UndefinedArgumentType ArgumentType = iota
  CommandArgument 
  FlagArgument
  ParamArgument
)

func (self ArgumentType) String() string {
  switch self {
  case CommandArgument:
    return "command"
  case FlagArgument:
    return "flag"
  case ParamArgument:
    return "param"
  default: // UndefinedArgumentType
    return ""
  }
}

type Argument interface {
  Type()    ArgumentType
  IsValid() bool
}

type arguments []*Argument 

func Arguments(arguments ...Argument) (argumentPointers arguments) { 
  for index, _ := range arguments {
    argumentPointers = append(argumentPointers, &arguments[index])
  }
  return argumentPointers
}

//func (self arguments) Strings() (argumentStrings []string) {
//  for _, argument := range self {
//    // TODO: Should it be value or name? It has to get the relevant data for
//    // each type and then each type needs to implement that new function. And
//    // the new function can't interfere with existing functionality.
//    argumentStrings = append(argumentStrings, argument.Value())
//  }
//}
//
//func (self arguments) String() string {
//  return strings.Join(self.Strings(), " ")
//}

func (self arguments) Last() *Argument { return self[self.Count()-1] }
func (self arguments) Count() int { return len(self) }

func (self arguments) Add(argument Argument) arguments {
  return append(self, &argument)
}
