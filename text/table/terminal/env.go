package term

import (
	"os"
	"strconv"
)

func GetEnvWindowSize() *Size {
	lines := os.Getenv("LINES")
	columns := os.Getenv("COLUMNS")
	if lines == "" && columns == "" {
		return nil
	}
	nLines := 0
	nColumns := 0
	var err error
	if lines != "" {
		nLines, err = strconv.Atoi(lines)
		if err != nil || nLines < 0 {
			return nil
		}
	}
	if columns != "" {
		nColumns, err = strconv.Atoi(columns)
		if err != nil || nColumns < 0 {
			return nil
		}
	}

	return &Size{
		Lines:   nLines,
		Columns: nColumns,
	}
}
