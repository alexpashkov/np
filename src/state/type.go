package state

import (
	"github.com/alexpashkov/npuzzle/src/heuristics"
	"github.com/alexpashkov/npuzzle/src/puzzle"
	"strconv"
)

type Id string

type State struct {
	Parent *State
	Puzzle puzzle.Puzzle
}

func (s State) Id() (id Id) {
	s.Puzzle.ForEach(func(t puzzle.Tile, _, _ int) (shellContinue bool) {
		id += Id(strconv.Itoa(t.Val()))
		return true
	})
	return
}

func (s State) G() (g int) {
	p := s.Parent
	for ; p != nil; p = p.Parent {
		g++
	}
	return
}

// takes heuristics function as a parameter
func (s *State) H(f heuristics.Fn) int {
	return f(s.Puzzle, puzzle.GetSolved(s.Puzzle.Size()))
}

// takes heuristics function as a parameter
func (s State) F(f heuristics.Fn) int {
	return s.G() + s.H(f)
}
