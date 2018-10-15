package board

// Counts number of inversions for a tile in position x, y
func TileInversions(board Board, x, y int) (inversions int) {
	examinedTile := board.Tile(x, y)
	for currY := y; currY < len(board); currY++ {
		for currX := x; currX < len(board); currX++ {
			currTile := board.Tile(currX, currY)
			if currTile.Val() != 0 &&
				examinedTile.Val() > currTile.Val() {
				inversions++
			}
		}
	}
	return
}

// Counts total number of inversions for the board
func Inversions(board Board) (inversions int) {
	ch := make(chan int)
	board.ForEach(func(_ Tile, x, y int) {
		go func() {
			ch <- TileInversions(board, x, y)
		}()
	})
	board.ForEach(func(_ Tile, _, _ int) {
		inversions += <-ch
	})
	return
}

// Checks if a board is solvable
//func IsSolvable(board Board) bool {
//
//}
