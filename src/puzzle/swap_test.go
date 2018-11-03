package puzzle

import "testing"

type swapTestCases struct {
	a, b      TileCoords
	orig, exp Puzzle
}

var tcs = []swapTestCases{
	{
		a: TileCoords{X: 0, Y: 0},
		b: TileCoords{X: 1, Y: 1},
		orig: Puzzle{
			{1, 2},
			{0, 3},
		},
		exp: Puzzle{
			{3, 2},
			{0, 1},
		},
	},
	{
		a: TileCoords{X: 0, Y: 0},
		b: TileCoords{X: 2, Y: 2},
		orig: Puzzle{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 0},
		},
		exp: Puzzle{
			{0, 2, 3},
			{4, 5, 6},
			{7, 8, 1},
		},
	},
	{
		a: TileCoords{X: 0, Y: 0},
		b: TileCoords{X: 1, Y: 2},
		orig: Puzzle{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 0},
		},
		exp: Puzzle{
			{8, 2, 3},
			{4, 5, 6},
			{7, 1, 0},
		},
	},
}

func TestSwap(t *testing.T) {
	for _, tc := range tcs {
		Swap(tc.orig, tc.a, tc.b)
		if !Equal(tc.orig, tc.exp) {
			t.Errorf("puzzles are not equal: coords: %v, %v", tc.orig, tc.exp)
		}
	}
}
