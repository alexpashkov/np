package board

// Counts number of inversions for a tile in position x, y
func TileInversions(board Board, x, y int) (inversions int) {
	examinedTile := board.Tile(x, y)
	for currY := y; currY < board.Size(); currY++ {
		currX := 0
		if currY == y {
			currX = x
		}
		for ; currX < board.Size(); currX++ {
			currTile := board.Tile(currX, currY)
			if currTile.Val() != 0 && examinedTile.Val() > currTile.Val() {
				inversions++
			}
		}
	}
	return
}

// Counts total number of inversions for the board
func Inversions(board Board) (inversions int) {
	var (
		ch       = make(chan int)
		zeroXPos = 0
		zeroYPos = 0
	)
	board.ForEach(func(t Tile, x, y int) {
		if t.Val() == 0 {
			zeroXPos = x
			zeroYPos = y
		}
		go func() {
			ch <- TileInversions(board, x, y)
		}()
	})
	board.ForEach(func(_ Tile, _, _ int) {
		inversions += <-ch
	})
	// if board size is even (there is no central tile), do some magic:
	if board.Size()%2 == 0 {
		inversions += zeroYPos*board.Size() + zeroXPos
	}
	return
}

// Checks if a board is solvable
func IsSolvable(board Board) bool {
	return Inversions(board)%2 == Inversions(GetSolved(board.Size()))%2
}
