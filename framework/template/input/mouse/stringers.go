package mouse

import "fmt"

const _State_name = "InvalidStateDownUp"

var _State_index = [...]uint8{12, 16, 18}

func (i State) String() string {
	if i >= State(len(_State_index)) {
		return fmt.Sprintf("State(%d)", i)
	}
	hi := _State_index[i]
	lo := uint8(0)
	if i > 0 {
		lo = _State_index[i-1]
	}
	return _State_name[lo:hi]
}

const _Button_name = "InvalidOneTwoThreeFourFiveSixSevenEight"

var _Button_index = [...]uint8{7, 10, 13, 18, 22, 26, 29, 34, 39}

func (i Button) String() string {
	if i >= Button(len(_Button_index)) {
		return fmt.Sprintf("Button(%d)", i)
	}
	hi := _Button_index[i]
	lo := uint8(0)
	if i > 0 {
		lo = _Button_index[i-1]
	}
	return _Button_name[lo:hi]
}
