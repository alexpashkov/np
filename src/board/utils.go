package board

// Flattens the board. Assumes, that the board is valid.
// If includeZero is false, zero would not be included
func Flatten(board Board, includeZero bool) Row {
	var size = len(board) * len(board)
	if !includeZero {
		size--
	}
	row := make(Row, size)
	for y := range board {
		for x := range board[y] {
			if !includeZero && board[y][x] == 0 {
				continue
			}
			row[y*len(board)+x] = board[y][x]
		}
	}
	return row
}
