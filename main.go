package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

type Puzzle struct {
	board [][]int
}

func removeComments(line string) string {
	return regexp.MustCompile("#.+").ReplaceAllLiteralString(line, "")
}

func parseSize(line string) (size int, err error) {
	_, err = fmt.Sscanf(line, "%d", &size)
	return size, err
}

func parseRow(line string, size int) ([]int, error) {
	row := make([]int, size)
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
		row[i] = tileIntVal
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
			var row []int
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

type Direction struct {
	x, y int
}

func nextDirectionGenerator() func() Direction {
	directions, i := [4]Direction{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}, 0
	return func() Direction {
		n := directions[i%4]
		i++
		return n
	}
}

func getNextCoords(x, y int, direction Direction) (int, int) {
	x += direction.x
	y += direction.y
	return x, y
}

func shouldChangeDirection(board [][]int, x, y int, direction Direction) bool {
	x, y = getNextCoords(x, y, direction)
	return x < 0 || y < 0 || x >= len(board) || y >= len(board) || board[y][x] != 0
}

func GetSolvedBoard(size int) [][]int {
	board := make([][]int, size)
	for i := range board {
		board[i] = make([]int, size)
	}
	getNextDirection := nextDirectionGenerator()
	direction, x, y := getNextDirection(), 0, 0
	for i := 1; i < size*size; i++ {
		board[y][x] = i
		if shouldChangeDirection(board, x, y, direction) {
			direction = getNextDirection()
		}
		x, y = getNextCoords(x, y, direction)
	}
	return board
}

func main() {
	//puzzle, err := readPuzzle(os.Stdin)
	//if err == nil {
	//	fmt.Println("Puzzle:", puzzle)
	//} else {
	//	fmt.Println("Error:", err)
	//}
	fmt.Println(GetSolvedBoard(6))
}
