package heuristics

import (
	"github.com/alexpashkov/npuzzle/src/puzzle"
	"github.com/alexpashkov/npuzzle/src/utils"
)

// calc distance for particular tile
func distance(origTile puzzle.Tile, origX, origY int, target puzzle.Puzzle) int {
	targetX, targetY := target.Coords(origTile)
	return utils.Abs(origX-targetX) + utils.Abs(origY-targetY)
}

// Takes puzzle and solved puzzle states
var Manhattan Fn = func(p puzzle.Puzzle) (dist int) {
	distancesCh := make(chan int)
	p.ForEach(func(t puzzle.Tile, x, y int) (shellContinue bool) {
		go func() {
			distancesCh <- distance(t, x, y, puzzle.GetSolved(p.Size()))
		}()
		return true
	})
	p.ForEach(func(_ puzzle.Tile, _, _ int) (shellContinue bool) {
		dist += <-distancesCh
		return true
	})
	return
}
