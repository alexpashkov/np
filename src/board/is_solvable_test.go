package board

import "testing"

func TestTileInversions(t *testing.T) {
	type testCase struct {
		x, y, expected int
	}
	board := Board{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 0},
	}
	testCases := []testCase{
		{
			x: 0, y: 0, expected: 0,
		},
		{
			x: 1, y: 0, expected: 0,
		},
		{
			x: 2, y: 0, expected: 0,
		},
		{
			x: 0, y: 1, expected: 0,
		},
		{
			x: 1, y: 2, expected: 0,
		},
	}
	for _, tc := range testCases {
		got := TileInversions(board, tc.x, tc.y)
		if got != tc.expected {
			t.Error("Wrong number of inversions, got ", got, "expected: ", tc.expected)
		}
	}
}

func TestInversions(t *testing.T) {
	type testCase struct {
		board    Board
		expected int
	}

	testCases := []testCase{
		{
			board: Board{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 0},
			},
			expected: 15,
		},
		{
			board: Board{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 0},
			},
			expected: 0,
		},
		{
			board: Board{
				{2, 1, 3},
				{4, 5, 6},
				{7, 8, 0},
			},
			expected: 1,
		},
		{
			board: Board{
				{2, 3, 1},
				{4, 5, 6},
				{7, 8, 0},
			},
			expected: 2,
		},
		{
			board: Board{
				{1, 4, 2},
				{5, 6, 3},
				{7, 8, 0},
			},
			expected: 4,
		},
	}
	for _, tc := range testCases {
		bi := Inversions(tc.board)
		if bi != tc.expected {
			t.Error("Wrong number of inversions, got ", bi, "expected: ", tc.expected)
		}
	}
}
