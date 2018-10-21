package puzzle

// Counts number of inversions for a tile in position x, y
func TileInversions(p Puzzle, x, y int) (inversions int) {
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
func Inversions(p Puzzle) (invs int) {
	var (
		ch       = make(chan int)
		zeroXPos = 0
		zeroYPos = 0
	)
	p.ForEach(func(t Tile, x, y int) (shellContinue bool) {
		if t.Val() == 0 {
			zeroXPos = x
			zeroYPos = y
		}
		go func() {
			ch <- TileInversions(p, x, y)
		}()
		return true
	})
	p.ForEach(func(_ Tile, _, _ int) (shellContinue bool) {
		invs += <-ch
		return true
	})
	// if p size is even (there is no central tile), do some magic:
	if p.Size()%2 == 0 {
		invs += zeroYPos*p.Size() + zeroXPos
	}
	return
}

// Checks if a p is solvable
func IsSolvable(p, sp Puzzle) bool {
	if sp == nil {
		sp = GetSolved(p.Size())
	}
	return Inversions(p)%2 == Inversions(sp)%2
}
