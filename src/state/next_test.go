package state

import (
	"github.com/alexpashkov/npuzzle/src/puzzle"
	"testing"
)

func TestNext(t *testing.T) {
	Next(State{Puzzle: puzzle.Puzzle{
		puzzle.Row{1, 2, 3},
		puzzle.Row{4, 5, 6},
		puzzle.Row{7, 8, 9},
	}})
}
