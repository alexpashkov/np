package board

func TileInversions(board Board, x, y int) (inversions int) {
	examinedTile := board[y][x]
	board.ForEach(func(currTile Tile, _, _ int) {
		if examinedTile.Val() > currTile.Val() {
			inversions++
		}
	})
	return
}

func BoardInversions(board Board) {

}

// Checks if a board is solvable
func IsSolvable(board Board) bool {

}
