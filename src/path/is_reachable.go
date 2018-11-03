package path

import "github.com/alexpashkov/npuzzle/src/puzzle"

// Counts number of inversions for a tile in position x, y
func tileInversions(p puzzle.Puzzle, x, y int) (inversions int) {
	examinedTile := p.Tile(x, y)
	for currY := y; currY < p.Size(); currY++ {
		currX := 0
		if currY == y {
			currX = x
		}
		for ; currX < p.Size(); currX++ {
			currTile := p.Tile(currX, currY)
			if currTile.Val() != 0 && examinedTile.Val() > currTile.Val() {
				inversions++
			}
		}
	}
	return
}

// Counts total number of inversions for the p
func inversions(p puzzle.Puzzle) (invs int) {
	var (
		ch       = make(chan int)
		zeroXPos = 0
		zeroYPos = 0
	)
	p.ForEach(func(t puzzle.Tile, x, y int) {
		if t.Val() == 0 {
			zeroXPos = x
			zeroYPos = y
		}
		go func() {
			ch <- tileInversions(p, x, y)
		}()
	})
	p.ForEach(func(_ puzzle.Tile, _, _ int) {
		invs += <-ch
	})
	// if p size is even (there is no central tile), do some magic:
	if p.Size()%2 == 0 {
		invs += zeroYPos*p.Size() + zeroXPos
	}
	return
}

// Checks if it is possible to go from to to
func CanBeReached(from, to puzzle.Puzzle) bool {
	return inversions(from)%2 == inversions(to)%2
}
