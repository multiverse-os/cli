package banner

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"path"
	"strconv"
	"strings"
)

const ascii_offset = 32
const first_ascii = ' '
const last_ascii = '~'

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
	font.setName(name)
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

func (font *font) setName(name string) {
	font.name = name
	if len(name) < 1 {
		font.name = defaultFont
	}
}

func (font *font) setAttributes(scanner *bufio.Scanner) {
	for scanner.Scan() {
		text := scanner.Text()
		if strings.HasPrefix(text, signature) {
			font.height = getHeight(text)
			font.baseline = getBaseline(text)
			font.hardblank = getHardblank(text)
			break
		}
	}
}

func (font *font) setLetters(scanner *bufio.Scanner) {
	font.letters = append(font.letters, make([]string, font.height, font.height)) //TODO: set spaces from flf
	for i := range font.letters[0] {
		font.letters[0][i] = "  "
	} //TODO: set spaces from flf
	letterIndex := 0
	for scanner.Scan() {
		text, cutLength, letterIndexInc := scanner.Text(), 1, 0
		if lastCharLine(text, font.height) {
			font.letters = append(font.letters, []string{})
			letterIndexInc = 1
			if font.height > 1 {
				cutLength = 2
			}
		}
		if letterIndex > 0 {
			appendText := ""
			if len(text) > 1 {
				appendText = text[:len(text)-cutLength]
			}
			font.letters[letterIndex] = append(font.letters[letterIndex], appendText)
		}
		letterIndex += letterIndexInc
	}
}

func (font *font) evenLetters() {
	var longest int
	for _, letter := range font.letters {
		if len(letter) > 0 && len(letter[0]) > longest {
			longest = len(letter[0])
		}
	}
	for _, letter := range font.letters {
		for i, row := range letter {
			letter[i] = row + strings.Repeat(" ", longest-len(row))
		}
	}
}

func scrub(text string, char byte) string {
	return strings.Replace(text, string(char), " ", -1)
}

func (figure figure) Slicify() (rows []string) {
	for r := 0; r < figure.font.height; r++ {
		printRow := ""
		for _, char := range figure.phrase {
			if char < first_ascii || char > last_ascii {
				char = '?'
			}
			fontIndex := char - ascii_offset
			charRowText := scrub(figure.font.letters[fontIndex][r], figure.font.hardblank)
			printRow += charRowText
		}
		if r < figure.font.baseline || len(strings.TrimSpace(printRow)) > 0 {
			rows = append(rows, strings.TrimRight(printRow, " "))
		}
	}
	return rows
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
