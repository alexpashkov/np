package state

import (
	"github.com/alexpashkov/npuzzle/src/heuristics"
	"github.com/alexpashkov/npuzzle/src/puzzle"
	"strconv"
)

type State struct {
	empty  puzzle.TileCoords
	id     string
	Parent *State
	Puzzle puzzle.Puzzle
}

func (s *State) Id() string {
	if s.id == "" {
		s.Puzzle.ForEach(func(t puzzle.Tile, _, _ int) {
			s.id += strconv.Itoa(t.Val())
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
func (s *State) H(f heuristics.Func) int {
	return f(s.Puzzle)
}

// takes heuristics function as a parameter
func (s *State) F(f heuristics.Func) int {
	return s.G() + s.H(f)
}

func (s *State) EmptyTileCoords() puzzle.TileCoords {
	if s.empty.X == 0 && s.empty.Y == 0 {
		s.empty = s.Puzzle.Find(0)
	}
	return s.empty
}
