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
func tileDist(tile puzzle.Tile, x, y int, solvedP puzzle.Puzzle) int {
	sx, sy := solvedP.Coords(tile)
	return Abs(x - sx) + Abs(y - sy)
}

// Takes puzzle and solved puzzle states
func Manhattan(p, solvedP puzzle.Puzzle) (dist int) {
	distCh := make(chan int)
	p.ForEach(func(t puzzle.Tile, x, y int) (shellContinue bool) {
		go func() {
			distCh <- tileDist(t, x, y, solvedP)
		}()
		return true
	})
	p.ForEach(func(_ puzzle.Tile, _, _ int) (shellContinue bool) {
		dist += <-distCh
		return true
	})
	return
}
