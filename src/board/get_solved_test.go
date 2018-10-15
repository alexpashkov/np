package board

import "testing"

type TestCase struct {
	size  int
	board Board
}

var (
	passingTestCases = []TestCase{
		{
			size: 2,
			board: Board{
				{1, 2},
				{0, 3},
			},
		},
		{
			size: 3,
			board: Board{
				{1, 2, 3},
				{8, 0, 4},
				{7, 6, 5},
			},
		},
		{
			size: 4,
			board: Board{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 0, 15, 6},
				{10, 9, 8, 7},
			},
		},
	}
	failingTestCases = []TestCase{
		{
			size: 2,
			board: Board{
				{1, 2},
				{3, 4},
			},
		},
		{
			size: 3,
			board: Board{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		{
			size: 4,
			board: Board{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 9, 15, 6},
				{10, 9, 8, 7},
			},
		},
	}
)

// Checks if 2 boards are equal
func boardEq(a Board, b Board) bool {
	if len(a) != len(b) {
		return false
	}
	for y := range a {
		if len(a[y]) != len(b[y]) {
			return false
		}
		for x := range a[y] {
			if a[y][x] != b[y][x] {
				return false
			}
		}
	}
	return true
}

func TestGetSolved(t *testing.T) {
	for _, testCase := range passingTestCases {
		solvedBoard := GetSolved(testCase.size)
		if !boardEq(solvedBoard, testCase.board) {
			t.Error("Boards are not equal, size: ", testCase.size,
				"expected: ", testCase.board, "got: ", solvedBoard)
		}
	}
	for _, testCase := range failingTestCases {
		solvedBoard := GetSolved(testCase.size)
		if boardEq(solvedBoard, testCase.board) {
			t.Error("Boards are not equal, size: ", testCase.size,
				"expected: ", testCase.board, "got: ", solvedBoard)
		}
	}
}
