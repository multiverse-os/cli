package data

// TODO: Perhaps have built-in
type Type int

// We doon't want too offer too many options, because that goes out of scope,
// but providing the tools that will enable validation for 98% of the input is
// desirable if its not too much. Since it will enable developers too have input
// validation and consistent input across all applications, even simple scripts
// with little cost. Like MAC could be included, but its not nearly as common as
// IP or Filename as a datatype. We could do a automated survey of GNU coreutils to
// try to determine the most common types in the future with greater certainty.
const (
	Bool Type = iota
	Int
	Float
	String
	Directory
	Filename
	Filenames // Via globbing or comma separated values
	URL
	Port
	IPv4
	IPv6
)

// TODO: This should migrate into its own package (the generic equivialence
// helpers) the idea would be to provide a collection of helpers to make all of
// our code cleaner and more expressive in the same way rails extends the
// epxressiveness of default ruby
// These may seem pointless but they will also simplify validation of values and
// provide helpers for developers to simplify validation

//
// Alias (For more expressive code)
///////////////////////////////////////////////////////////////////////////////
//var Blank = ""
//var Whitespace = " "
//var Tab = "\t"
//var NewLine = "\n"

//var True = true
//var False = false

//
// Transform
///////////////////////////////////////////////////////////////////////////////

//
// Validate
///////////////////////////////////////////////////////////////////////////////

// Encode
///////////////////////////////////////////////////////////////////////////////
// Normally convention is not to use "to" but that is regarding methods, if we
// migrate this to a more generic flexible data type for holding variable user
// input. We can migrate
//func (self Data) String() string {
//	return fmt.Sprintf(self)
//}
//
//// String Subtypes
//func (self Data) Directory() string { return self.(string) }
//func (self Data) Filename() string  { return self.(string) }
//func (self Data) Filenames() string { return self.(string) }
//func (self Data) URL() string       { return self.(string) }
//func (self Data) IPv4() string      { return self.(string) }
//func (self Data) IPv6() string      { return self.(string) }
//
//func (self Data) Int() int {
//	intValue, err := strconv.Atoi(self.(string))
//	if err != nil {
//		return 0
//	} else {
//		return intValue
//	}
//}
//
//// Int Subtypes
//func (self Data) Port() int { return toInt(value) }
//
//func (self Data) Float() float64 {
//	floatParts := strings.Split(value.(string), ".")
//	if len(floatParts) > 1 {
//		floatValue, err := strconv.ParseFloat(value.(string), len(floatParts[1]))
//		if err != nil {
//			return float64(0.00)
//		} else {
//			return floatValue
//		}
//	}
//	return float64(0)
//}
//
//func (self Data) Bool() bool {
//	for _, trueValue := range trueValues {
//		if value.(string) == trueValue {
//			return true
//		}
//	}
//	return false
//}
//
////
//// Public Methods
/////////////////////////////////////////////////////////////////////////////////
//// TODO: Basic validation should move here, and basic conversion (which is in
//// flags since this is where it orginated, but it became clear it would also be
//// important for parameters) [BUT maybe parameters don't have a type? Just hand
//// off the string, just seems if we are dealing with flag type we should go
//// ahead and extend this to parameters, but it may be wise to just ignore both]
//func (self Data) Valid(flagType DataType) (bool, error) {
//	switch flagType {
//	case Bool:
//		boolValues := append(trueValues, falseValues...)
//		for _, boolValue := range boolValues {
//			if boolValue == value {
//				return true, nil
//			}
//		}
//		return false, errors.New("[error] could not parse valid boolean value")
//	//case Int:
//	//case String:
//	//case Directory:
//	case Filename:
//		_, err := os.Stat(value.(string))
//		return (err == nil), nil
//	//case Filenames:
//	//case URL:
//	//case IPv4:
//	//case IPv6:
//	//case Port:
//	default:
//		return false, errors.New("[error] failed to parse data type")
//	}
//}
