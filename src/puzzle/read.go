package puzzle

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

func removeComments(line string) string {
	return regexp.MustCompile("#.+").ReplaceAllLiteralString(line, "")
}

func parseSize(line string) (size int, err error) {
	_, err = fmt.Sscanf(line, "%d", &size)
	return size, err
}

func parseRow(line string, size int) (Row, error) {
	row := make(Row, size)
	line = strings.Trim(line, " ")
	tileStringVals := regexp.MustCompile(`\s+`).Split(line, -1)
	if len(tileStringVals) != size {
		return nil, errors.New("invalid number of tiles in a line")
	}
	for i, tileStringVal := range tileStringVals {
		tileIntVal, err := strconv.Atoi(tileStringVal)
		if err != nil {
			return nil, errors.New("tile value is not an integer")
		}
		row[i] = Tile(tileIntVal)
	}
	return row, nil
}

func Read(from io.Reader) (puzzle Puzzle, err error) {
	var (
		scanner        = bufio.NewScanner(from)
		size           int
		sizeHasBeenSet = false
	)
	for scanner.Scan() {
		line := removeComments(scanner.Text())
		if len(line) == 0 {
			continue
		}
		if !sizeHasBeenSet {
			size, err = parseSize(line)
			sizeHasBeenSet = true
		} else {
			var row Row
			if len(puzzle) >= size {
				return puzzle, errors.New("invalid p size")
			}
			row, err = parseRow(line, size)
			if err == nil {
				puzzle = append(puzzle, row)
			}
		}
		if err != nil {
			return
		}
	}
	return
}
