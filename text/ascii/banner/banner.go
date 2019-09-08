package banner

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"path"
	"strings"
)

const defaultFont = "standard"

type Banner struct {
	Phrase string
	Lines  []string
	Font   string
	Height int
	Width  int
	Length int
}

func (self Banner) String() (output string) {
	for _, line := range self.Lines {
		output += line + "\n"
	}
	return output
}

func New(fontName, text string) Banner {
	font := newFont(fontName)

	banner := Banner{
		Text:   text,
		Font:   font,
		Width:  width,
		Lines:  banner.Slicify(),
		Length: len(text),
	}
	banner.Width = len(banner.Lines[0])

	return banner
}

func BigFont(text string) Banner         { return New("big", text) }
func ChunkyFont(text string) Banner      { return New("chunky", text) }
func CyberLargeFont(text string) Banner  { return New("cyberlarge", text) }
func CyberMediumFont(text string) Banner { return New("cybermedium", text) }
func DoomFont(text string) Banner        { return New("doom", text) }
func DrPepperFont(text string) Banner    { return New("drpepper", text) }
func EliteFont(text string) Banner       { return New("elite", text) }
func Isometric3Font(text string) Banner  { return New("isometric3", text) }
func Isometric4Font(text string) Banner  { return New("isometric4", text) }
func IvritFont(text string) Banner       { return New("ivrit", text) }
func JerusalemFont(text string) Banner   { return New("jerusalem", text) }
func Larry3DFont(text string) Banner     { return New("larry3d", text) }
func LCDFont(text string) Banner         { return New("lcd", text) }
func LeanFont(text string) Banner        { return New("lean", text) }
func LettersFont(text string) Banner     { return New("letters", text) }
func LinuxFont(text string) Banner       { return New("linux", text) }
func LockerGnomeFont(text string) Banner { return New("lockergnome", text) }
func MaxFourFont(text string) Banner     { return New("maxfour", text) }
func NancyJFont(text string) Banner      { return New("nancyj", text) }
func NTGreekFont(text string) Banner     { return New("ntgreek", text) }
func PepperFont(text string) Banner      { return New("pepper", text) }
func RectanglesFont(text string) Banner  { return New("rectangles", text) }
func ReliefFont(text string) Banner      { return New("relief", text) }
func Relief2Font(text string) Banner     { return New("relief2", text) }
func SmallFont(text string) Banner       { return New("small", text) }
func Smisome1Font(text string) Banner    { return New("smisome1", text) }
func StandardFont(text string) Banner    { return New("standard", text) }
func TicksFont(text string) Banner       { return New("ticks", text) }
func TicksSlantFont(text string) Banner  { return New("ticksslant", text) }
