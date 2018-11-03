package state

import (
	"github.com/alexpashkov/npuzzle/src/heuristics"
	"github.com/alexpashkov/npuzzle/src/puzzle"
	"strconv"
)

type State struct {
	id     string
	Parent *State
	Puzzle puzzle.Puzzle
}

func (s *State) Id() string {
	if s.id == "" {
		s.Puzzle.ForEach(func(t puzzle.Tile, _, _ int) (shellContinue bool) {
			s.id += strconv.Itoa(t.Val())
			return true
		})
	}
	return s.id
}

func (s *State) G() (g int) {
	p := s.Parent
	for ; p != nil; p = p.Parent {
		g++
	}
	return
}

// takes heuristics function as a parameter
func (s *State) H(f heuristics.Fn, sp puzzle.Puzzle) int {
	return f(s.Puzzle, sp)
}

// takes heuristics function as a parameter
func (s *State) F(f heuristics.Fn, sp puzzle.Puzzle) int {
	return s.G() + s.H(f, sp)
}
