package heuristics

import (
	"github.com/alexpashkov/npuzzle/src/puzzle"
	"testing"
)

type testCase struct {
	p   puzzle.Puzzle
	exp int
}

var testCases = []testCase{
	{
		exp: 0,
		p: puzzle.Puzzle{
			{1, 2},
			{0, 3},
		},
	},
	{
		exp: 4,
		p: puzzle.Puzzle{
			{3, 2},
			{0, 1},
		},
	},
	{
		p: puzzle.Puzzle{
			{1, 2, 3},
			{8, 0, 4},
			{7, 6, 5},
		},
		exp: 0,
	},
	{
		p: puzzle.Puzzle{
			{1, 2, 3, 4},
			{12, 13, 14, 5},
			{11, 0, 15, 6},
			{10, 9, 8, 7},
		},
		exp: 0,
	},
	{
		p: puzzle.Puzzle{
			{1, 2, 3, 4},
			{12, 13, 14, 5},
			{11, 0, 15, 6},
			{10, 8, 9, 7},
		},
		exp: 2,
	},
	{
		p: puzzle.Puzzle{
			{7, 2, 3, 4},
			{12, 13, 14, 5},
			{11, 0, 15, 6},
			{10, 8, 9, 1},
		},
		exp: 14,
	},
}

func TestManhattan(t *testing.T) {
	for _, tc := range testCases {
		v := manhattan(tc.p)
		if v != tc.exp {
			t.Error("wrong h() for puzzle: ", tc.p,
				"expected: ", tc.exp, "got: ", v)
		}
	}
}
