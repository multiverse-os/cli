package terminal

import (
	"syscall"
	"unsafe"

	"github.com/multiverse-os/cli/os"
)

type Terminal struct {
	Title  string
	Cursor cursor.Cursor // TODO: Really want to support multiseat

	// TODO: At least an active buffer, and shaddow buffer, but possibly several
	//       so one can be an overlay, etc.

	// TODO: Require, Ticks, Subscribers, information about the logged in user
	//       etc
	User             os.User
	WorkingDirectory os.Path

	Dimensions terminalDimensions
}

type terminalDimensions struct {
	CharacterWidth  uint16
	CharacterHeight uint16
	PixelWidth      uint16
	PixelHeight     uint16
}

func New() Terminal {
	return Terminal{
		Dimensions: terminalDimensions{
			CharacterWidth:  0,
			CharacterHeight: 0,
		},
	}
}

func Width() uint {
	dimensions := &terminalDimensions{}
	data, _, err := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(dimensions)),
	)

	if int(data) == -1 {
		panic(err)
	}
	return uint(dimensions.CharacterWidth)
}
package cli

type Short int16

type Coord struct {
	X Short
	Y Short
}

// TODO:
//        We will load the terminal, with the cursor location, height.
//        then modify values and provide a hook on resize
//        and use this for the output.go logging functionality
//         -
//        **keep in mind that command output like help is execpted to be
//        different than logging otuput, it can be the same but its typical to
//        have that go to  a log file.**

//func NewTerminal() *Terminal {
//Stdio: terminal.Stdio{
//	In:  os.Stdin,
//	Out: os.Stdout,
//	Err: os.Stderr,
//	w, h := terminal.WindowSize()
//	return &terminal.Terminal{
//		ConsoleStream: types.NewConsoleStream(),
//		Cursor:        *terminal.NewCursor(),
//		width:         w,
//		height:        h,
//	}
//}

//type Icon struct {
//	Text   string
//	Format string
//}
//type IconSet struct {
//	HelpInput      Icon
//	Error          Icon
//	Help           Icon
//	Question       Icon
//	MarkedOption   Icon
//	UnmarkedOption Icon
//	SelectFocus    Icon
//}

// Validator is a function passed to a Question after a user has provided a response.
// If the function returns an error, then the user will be prompted again for another
// response.
//type Validator func(ans interface{}) error
//
//// Transformer is a function passed to a Question after a user has provided a response.
//// The function can be used to implement a custom logic that will result to return
//// a different representation of the given answer.
////
//// Look `TransformString`, `ToLower` `Title` and `ComposeTransformers` for more.
//type Transformer func(ans interface{}) (newAns interface{})
//
//
//func (r *Renderer) NewCursor() *terminal.Cursor {
//	return &terminal.Cursor{
//		In:  r.stdio.In,
//		Out: r.stdio.Out,
//	}
//}

// var ErrorTemplate = `{{color .Icon.Format }}{{ .Icon.Text }} Sorry, your reply was invalid: {{ .Error.Error }}{{color "reset"}}

//type ErrorTemplateData struct {
//	Error error
//	Icon  Icon
//}

//func (r *Renderer) Render(tmpl string, data interface{}) error {
//	// cleanup the currently rendered text
//	lineCount := r.countLines(r.renderedText)
//	r.resetPrompt(lineCount)
//	r.renderedText.Reset()
//
//	// render the template summarizing the current state
//	userOut, layoutOut, err := core.RunTemplate(tmpl, data)
//	if err != nil {
//		return err
//	}
//
//	// print the summary
//	fmt.Fprint(terminal.NewAnsiStdout(r.stdio.Out), userOut)
//
//	// add the printed text to the rendered text buffer so we can cleanup later
//	r.AppendRenderedText(layoutOut)
//
//	// nothing went wrong
//	return nil
//}

// Required does not allow an empty value
//func Required(val interface{}) error {
//	// the reflect value of the result
//	value := reflect.ValueOf(val)
//
//	// if the value passed in is the zero value of the appropriate type
//	if isZero(value) && value.Kind() != reflect.Bool {
//		return errors.New("Value is required")
//	}
//	return nil
//}
//
//// MaxLength requires that the string is no longer than the specified value
//func MaxLength(length int) Validator {
//	// return a validator that checks the length of the string
//	return func(val interface{}) error {
//		if str, ok := val.(string); ok {
//			// if the string is longer than the given value
//			if len([]rune(str)) > length {
//				// yell loudly
//				return fmt.Errorf("value is too long. Max length is %v", length)
//			}
//		} else {
//			// otherwise we cannot convert the value into a string and cannot enforce length
//			return fmt.Errorf("cannot enforce length on response of type %v", reflect.TypeOf(val).Name())
//		}
//
//		// the input is fine
//		return nil
//	}
//}
//
//// MinLength requires that the string is longer or equal in length to the specified value
//func MinLength(length int) Validator {
//	// return a validator that checks the length of the string
//	return func(val interface{}) error {
//		if str, ok := val.(string); ok {
//			// if the string is shorter than the given value
//			if len([]rune(str)) < length {
//				// yell loudly
//				return fmt.Errorf("value is too short. Min length is %v", length)
//			}
//		} else {
//			// otherwise we cannot convert the value into a string and cannot enforce length
//			return fmt.Errorf("cannot enforce length on response of type %v", reflect.TypeOf(val).Name())
//		}
//
//		// the input is fine
//		return nil
//	}
//}
//
//// ComposeValidators is a variadic function used to create one validator from many.
//func ComposeValidators(validators ...Validator) Validator {
//	// return a validator that calls each one sequentially
//	return func(val interface{}) error {
//		// execute each validator
//		for _, validator := range validators {
//			// if the answer's value is not valid
//			if err := validator(val); err != nil {
//				// return the error
//				return err
//			}
//		}
//		// we passed all validators, the answer is valid
//		return nil
//	}
//}
//
//// isZero returns true if the passed value is the zero object
//func isZero(v reflect.Value) bool {
//	switch v.Kind() {
//	case reflect.Slice, reflect.Map:
//		return v.Len() == 0
//	}
//
//	// compare the types directly with more general coverage
//	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
//}
