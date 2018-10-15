package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/alexpashkov/npuzzle/src/board"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type Puzzle struct {
	board board.Board
}

func removeComments(line string) string {
	return regexp.MustCompile("#.+").ReplaceAllLiteralString(line, "")
}

func parseSize(line string) (size int, err error) {
	_, err = fmt.Sscanf(line, "%d", &size)
	return size, err
}

func parseRow(line string, size int) (board.Row, error) {
	row := make(board.Row, size)
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
		row[i] = board.Tile(tileIntVal)
	}
	return row, nil
}

func readPuzzle(from io.Reader) (puzzle Puzzle, err error) {
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
			var row board.Row
			if len(puzzle.board) >= size {
				return puzzle, errors.New("invalid puzzle size")
			}
			row, err = parseRow(line, size)
			if err == nil {
				puzzle.board = append(puzzle.board, row)
			}
		}
		if err != nil {
			return
		}
	}
	return
}

func main() {
	//puzzle, err := readPuzzle(os.Stdin)
	//if err == nil {
	//	fmt.Println("Puzzle:", puzzle)
	//} else {
	//	fmt.Println("Error:", err)
	//}
	fmt.Println(board.GetSolved(6))
}
