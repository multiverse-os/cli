package banner

import (
	"bufio"
	"bytes"
	"io"
	"path"
	"strconv"
	"strings"
)

var charDelimiters = [3]string{"@", "#", "$"}
var hardblanksBlacklist = [2]byte{'a', '2'}

type font struct {
	name      string
	height    int
	baseline  int
	hardblank byte
	letters   [][]string
}

func newFont(name string) (font font) {
	font.name = name
	if len(name) < 1 {
		font.name = defaultFont
	}
	fontBytes, err := Asset(path.Join("fonts", font.name+".flf"))
	if err != nil {
		panic(err)
	}
	fontBytesReader := bytes.NewReader(fontBytes)
	scanner := bufio.NewScanner(fontBytesReader)
	font.setAttributes(scanner)
	font.setLetters(scanner)
	return font
}

func newFontFromReader(reader io.Reader) (font font) {
	scanner := bufio.NewScanner(reader)
	font.setAttributes(scanner)
	font.setLetters(scanner)
	return font
}

func (self *font) setAttributes(scanner *bufio.Scanner) {
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, "flf2") {
			self.height = getHeight(text)
			self.baseline = getBaseline(text)
			self.hardblank = getHardblank(text)
			break
		}
	}
}

func (self *font) setLetters(scanner *bufio.Scanner) {
	self.letters = append(self.letters, make([]string, self.height, self.height)) //TODO: set spaces from flf
	for i := range self.letters[0] {
		self.letters[0][i] = "  "
	} //TODO: set spaces from flf
	letterIndex := 0
	for scanner.Scan() {
		text, cutLength, letterIndexInc := scanner.Text(), 1, 0
		if lastCharLine(text, self.height) {
			self.letters = append(self.letters, []string{})
			letterIndexInc = 1
			if self.height > 1 {
				cutLength = 2
			}
		}
		if letterIndex > 0 {
			appendText := ""
			if len(text) > 1 {
				appendText = text[:len(text)-cutLength]
			}
			self.letters[letterIndex] = append(self.letters[letterIndex], appendText)
		}
		letterIndex += letterIndexInc
	}
}

func (self *font) evenLetters() {
	var longest int
	for _, letter := range self.letters {
		if len(letter) > 0 && len(letter[0]) > longest {
			longest = len(letter[0])
		}
	}
	for _, letter := range self.letters {
		for i, row := range letter {
			letter[i] = row + strings.Repeat(" ", longest-len(row))
		}
	}
}

func getHeight(metadata string) int {
	datum := strings.Fields(metadata)[1]
	height, _ := strconv.Atoi(datum)
	return height
}

func getBaseline(metadata string) int {
	datum := strings.Fields(metadata)[2]
	baseline, _ := strconv.Atoi(datum)
	return baseline
}

func getHardblank(metadata string) byte {
	datum := strings.Fields(metadata)[0]
	hardblank := datum[len(datum)-1]
	if hardblank == hardblanksBlacklist[0] || hardblank == hardblanksBlacklist[1] {
		return ' '
	} else {
		return hardblank
	}
}

func lastCharLine(text string, height int) bool {
	endOfLine, length := "  ", 2
	if height == 1 && len(text) > 0 {
		length = 1
	}
	if len(text) >= length {
		endOfLine = text[len(text)-length:]
	}
	return endOfLine == strings.Repeat(charDelimiters[0], length) ||
		endOfLine == strings.Repeat(charDelimiters[1], length) ||
		endOfLine == strings.Repeat(charDelimiters[2], length)
}
