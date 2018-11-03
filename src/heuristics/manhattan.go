package heuristics

import (
	"github.com/alexpashkov/npuzzle/src/puzzle"
)

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// calc distance for particular tile
func tileDist(tile puzzle.Tile, x, y int, solved puzzle.Puzzle) int {
	sx, sy := solved.Coords(tile)
	return Abs(x-sx) + Abs(y-sy)
}

// Takes puzzle and solved puzzle states
var manhattan Func = func(p puzzle.Puzzle) (dist int) {
	distCh := make(chan int)
	p.ForEach(func(t puzzle.Tile, x, y int) {
		go func() {
			distCh <- tileDist(t, x, y, puzzle.GetSolved(p.Size()))
		}()
	})
	p.ForEach(func(_ puzzle.Tile, _, _ int) {
		dist += <-distCh
	})
	return
}
