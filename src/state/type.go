package state

import (
	"github.com/alexpashkov/npuzzle/src/puzzle"
	"strconv"
)

type State struct {
	parent *State
	puzzle  puzzle.Puzzle
}

func (s State) Id() (id string) {
	s.puzzle.ForEach(func(t puzzle.Tile, _, _ int) {
		id += strconv.Itoa(t.Val())
	})
	return
}

func (s State) G() (g int) {
	p := s.parent
	for ; p != nil; p = p.parent {
		g++
	}
	return
}

type HeuristicFunc func(b puzzle.Puzzle) int

// takes heuristics function as a parameter
func (s State) H(f HeuristicFunc) int {
	return f(s.puzzle)
}

// takes heuristics function as a parameter
func (s State) F(f HeuristicFunc) int {
	return s.G() + s.H(f)
}
