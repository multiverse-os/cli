package ansi

// TODO: Consider having a pointer object that can be tracked and more easily
// moved. Or perhaps just leverage this and put this concept in a higher level
// terminal library
func ShowCursor() string { return prefix + "?25h" }
func HideCursor() string { return prefix + "?25l" }

// Absolute Movement
func MoveCursorTo(row, column int) string { return prefix + row + ";" + column + "H" }

// Relative Movement
func MoveCursorUp(rows int) string         { return prefix + rows + "A" }
func MoveCursorDown(rows int) string       { return prefix + rows + "B" }
func MoveCursorRight(columns int) string   { return prefix + columns + "C" }
func MoveCursorLeft(columns int) string    { return prefix + columns + "D" }
func MoveCursorUpperLeft(count int) string { return prefix + count + "H" }
func MoveCursorToNextLine() string         { return escape + "E" }

func ClearLineRight() string    { return prefix + "0K" }
func ClearLineLeft() string     { return prefix + "1K" }
func ClearEntireLine() string   { return prefix + "2K" }
func ClearScreenDown() string   { return prefix + "0J" }
func ClearScreenUp() string     { return prefix + "1J" }
func ClearEntireScreen() string { return prefix + "2J" }

func SaveAttributes() string    { return escape + "7" }
func RestoreAttributes() string { return escape + "8" }
