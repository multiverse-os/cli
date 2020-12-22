package data

func IsZero(value int) bool  { return value == 0 }
func NotZero(value int) bool { return value != 0 }

func IsBlank(str string) bool  { return IsZero(len(str)) }
func NotBlank(str string) bool { return !IsZero(len(str)) }

func IsEmpty(value []interface{}) bool  { return IsZero(len(value)) }
func NotEmpty(value []interface{}) bool { return !IsZero(len(value)) }

func IsNil(value interface{}) bool {
	switch value.(type) {
	case int:
		return IsZero(value.(int))
	case string:
		return IsBlank(value.(string))
	case error:
		return value.(error) == nil
	default:
		return value == nil
	}
}
func NotNil(value interface{}) bool { return !IsNil(value) }

func IsGreaterThan(gt, value int) bool         { return (gt > value) }
func IsGreaterOrEqualThan(gte, value int) bool { return (gte >= value) }
func IsLessThan(lt, value int) bool            { return (lt < value) }
func IsLessOrEqualThan(lte, value int) bool    { return (lte <= value) }

func IsBetween(start, end, value int) bool  { return start < value && value < end }
func NotBetween(start, end, value int) bool { return start > value && value > end }
