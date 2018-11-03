package puzzle

import "testing"

type testCase struct {
	size int
	p    Puzzle
}

var (
	passingTestCases = []testCase{
		{
			size: 2,
			p: Puzzle{
				{1, 2},
				{0, 3},
			},
		},
		{
			size: 3,
			p: Puzzle{
				{1, 2, 3},
				{8, 0, 4},
				{7, 6, 5},
			},
		},
		{
			size: 4,
			p: Puzzle{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 0, 15, 6},
				{10, 9, 8, 7},
			},
		},
	}
	failingTestCases = []testCase{
		{
			size: 2,
			p: Puzzle{
				{1, 2},
				{3, 4},
			},
		},
		{
			size: 3,
			p: Puzzle{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		{
			size: 4,
			p: Puzzle{
				{1, 2, 3, 4},
				{12, 13, 14, 5},
				{11, 9, 15, 6},
				{10, 9, 8, 7},
			},
		},
	}
)

func TestGetSolved(t *testing.T) {
	for _, testCase := range passingTestCases {
		sp := GetSolved(testCase.size)
		if !Equal(sp, testCase.p) {
			t.Error("puzzles are not equal, size: ", testCase.size,
				"expected: ", testCase.p, "got: ", sp)
		}
	}
	for _, testCase := range failingTestCases {
		sp := GetSolved(testCase.size)
		if Equal(sp, testCase.p) {
			t.Error("puzzles are not equal, size: ", testCase.size,
				"expected: ", testCase.p, "got: ", sp)
		}
	}
}
