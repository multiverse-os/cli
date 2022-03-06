package data

import (
	"fmt"
	"strings"
)

type Boolean bool

const (
	True  Boolean = true
	False Boolean = false
)

func MarshalBool(input interface{}) (bool, error) {
	switch input.(type) {
	case string:
		if IsTrue(input.(string)) {
			return true, nil
		} else if IsFalse(input.(string)) {
			return false, nil
		}
	case int:
		if input.(int) == 1 {
			return true, nil
		} else if input.(int) == 0 {
			return false, nil
		}
	case bool:
		if input.(bool) {
			return true, nil
		} else if input.(bool) {
			return false, nil
		}

	}
	return false, fmt.Errorf("failed to marshal boolean value")
}

func IsTrue(value string) bool {
	value = strings.ToLower(value)
	for _, trueValue := range True.Strings() {
		if trueValue == value {
			return true
		}
	}
	return false
}

func IsFalse(value string) bool {
	value = strings.ToLower(value)
	for _, falseValue := range False.Strings() {
		if falseValue == value {
			return true
		}
	}
	return false
}

func IsBoolean(value string) bool { return IsFalse(value) || IsTrue(value) }

///////////////////////////////////////////////////////////////////////////////

func (self Boolean) Bool() bool { return bool(self) }

// TODO: Maybe in future give more options for string output, as in "1" "t"
// "yes" 
func (self Boolean) String() string {
	if self == True {
		return "true"
	} else {
		return "false"
	}
}

func (self Boolean) Int() int {
	if self == True {
		return 1
	} else {
		return 0
	}
}

func (self Boolean) Strings() []string {
	if self == True {
		return []string{"true", "yes", "y", "t", "1"}
	} else {
		return []string{"false", "no", "n", "f", "0"}
	}
}
