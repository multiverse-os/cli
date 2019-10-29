package keyboard

import "fmt"

const _State_name = "InvalidStateDownUp"

var _State_index = [...]uint8{0, 12, 16, 18}

func (i State) String() string {
	if i+1 >= State(len(_State_index)) {
		return fmt.Sprintf("State(%d)", i)
	}
	return _State_name[_State_index[i]:_State_index[i+1]]
}

const _Key_name = "InvalidTildeDashEqualsSemicolonApostropheCommaPeriodForwardSlashBackSlashBackspaceTabCapsLockSpaceEnterEscapeInsertPrintScreenDeletePageUpPageDownHomeEndPauseSleepClearSelectPrintExecuteHelpApplicationsScrollLockPlayZoomArrowLeftArrowRightArrowDownArrowUpLeftBracketLeftShiftLeftCtrlLeftSuperLeftAltRightBracketRightShiftRightCtrlRightSuperRightAltZeroOneTwoThreeFourFiveSixSevenEightNineF1F2F3F4F5F6F7F8F9F10F11F12F13F14F15F16F17F18F19F20F21F22F23F24F25ABCDEFGHIJKLMNOPQRSTUVWXYZNumLockNumMultiplyNumDivideNumAddNumSubtractNumZeroNumOneNumTwoNumThreeNumFourNumFiveNumSixNumSevenNumEightNumNineNumDecimalNumCommaNumEnterBrowserBackBrowserForwardBrowserRefreshBrowserStopBrowserSearchBrowserFavoritesBrowserHomeMediaNextMediaPreviousMediaStopMediaPlayPauseLaunchMailLaunchMediaLaunchAppOneLaunchAppTwoKanaKanjiJunjaAttnCrSelExSelEraseEOF"

var _Key_index = [...]uint16{0, 7, 12, 16, 22, 31, 41, 46, 52, 64, 73, 82, 85, 93, 98, 103, 109, 115, 126, 132, 138, 146, 150, 153, 158, 163, 168, 174, 179, 186, 190, 202, 212, 216, 220, 229, 239, 248, 255, 266, 275, 283, 292, 299, 311, 321, 330, 340, 348, 352, 355, 358, 363, 367, 371, 374, 379, 384, 388, 390, 392, 394, 396, 398, 400, 402, 404, 406, 409, 412, 415, 418, 421, 424, 427, 430, 433, 436, 439, 442, 445, 448, 451, 454, 455, 456, 457, 458, 459, 460, 461, 462, 463, 464, 465, 466, 467, 468, 469, 470, 471, 472, 473, 474, 475, 476, 477, 478, 479, 480, 487, 498, 507, 513, 524, 531, 537, 543, 551, 558, 565, 571, 579, 587, 594, 604, 612, 620, 631, 645, 659, 670, 683, 699, 710, 719, 732, 741, 755, 765, 776, 788, 800, 804, 809, 814, 818, 823, 828, 836}

func (i Key) String() string {
	if i < 0 || i+1 >= Key(len(_Key_index)) {
		return fmt.Sprintf("Key(%d)", i)
	}
	return _Key_name[_Key_index[i]:_Key_index[i+1]]
}
