package state

import (
	"github.com/alexpashkov/npuzzle/src/puzzle"
	"strconv"
	"github.com/alexpashkov/npuzzle/src/heuristics"
)

type State struct {
	parent *State
	puzzle puzzle.Puzzle
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

// takes heuristics function as a parameter
func (s State) H(f heuristics.Fn) int {
	return f(s.puzzle)
}

// takes heuristics function as a parameter
func (s State) F(f heuristics.Fn) int {
	return s.G() + s.H(f)
}
