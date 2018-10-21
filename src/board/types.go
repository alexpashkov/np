package board

import "strconv"

type Tile int

func (t Tile) Val() int {
	return int(t)
}

type Row []Tile
type Board []Row

func (b Board) Size() int {
	return len(b)
}

func (b Board) Tile(x, y int) Tile {
	return b[y][x]
}

func (b Board) ForEach(fn func(t Tile, x, y int)) {
	for y := range b {
		for x, t := range b[y] {
			fn(t, x, y)
		}
	}
}

type State struct {
	parent *State
	board  Board
}

func (s State) Id() (id string) {
	s.board.ForEach(func(t Tile, _, _ int) {
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

type HeuristicFunc func(b Board) int

// takes heuristics function as a parameter
func (s State) H(f HeuristicFunc) int {
	return f(s.board)
}

// takes heuristics function as a parameter
func (s State) F(f HeuristicFunc) int {
	return s.G() + s.H(f)
}
