package puzzle

import "testing"

func TestTileInversions(t *testing.T) {
	type testCase struct {
		x, y, expected int
	}
	p := Puzzle{
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
		got := TileInversions(p, tc.x, tc.y)
		if got != tc.expected {
			t.Error("Wrong number of inversions, got ", got, "expected: ", tc.expected)
		}
	}
}

func TestInversions(t *testing.T) {
	type testCase struct {
		p        Puzzle
		expected int
	}

	testCases := []testCase{
		{
			p: Puzzle{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 0},
			},
			expected: 15,
		},
		{
			p: Puzzle{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 0},
			},
			expected: 0,
		},
		{
			p: Puzzle{
				{2, 1, 3},
				{4, 5, 6},
				{7, 8, 0},
			},
			expected: 1,
		},
		{
			p: Puzzle{
				{2, 3, 1},
				{4, 5, 6},
				{7, 8, 0},
			},
			expected: 2,
		},
		{
			p: Puzzle{
				{1, 4, 2},
				{5, 6, 3},
				{7, 8, 0},
			},
			expected: 4,
		},
	}
	for _, tc := range testCases {
		bi := Inversions(tc.p)
		if bi != tc.expected {
			t.Error("Wrong number of inversions, got ", bi, "expected: ", tc.expected)
		}
	}
}
