package state

import (
	"github.com/alexpashkov/npuzzle/src/puzzle"
	"testing"
)

var nextOKTestCases = []struct {
	curr State
	exp  []State
}{
	{
	},
}

func (states States) IdPresent(id string) bool {
	for _, s := range states {
		if s.Id() == id {
			return true
		}
	}
	return false
}

func TestNext(t *testing.T) {
	currState := State{
		Puzzle: puzzle.Puzzle{
			{1, 2},
			{0, 3},
		},
	}
	expectedNextStates := States{
		{
			Parent: &currState,
			Puzzle: puzzle.Puzzle{
				{0, 2},
				{1, 3},
			},
		},
		{
			Parent: &currState,
			Puzzle: puzzle.Puzzle{
				{1, 2},
				{3, 0},
			},
		},
	}
	nextStates := Next(currState)

	if len(nextStates) != len(expectedNextStates) {
		t.Error("No eq")
	} else {
		for _, expectedNextState := range expectedNextStates {
			if !nextStates.IdPresent(expectedNextState.Id()) {

				t.Error("No eq")
			}
		}
	}
}
