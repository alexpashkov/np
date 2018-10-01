package board

type Board [][]int

type Direction struct {
	x, y int
}

func nextDirectionGenerator() func() Direction {
	directions, i := [4]Direction{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}, 0
	return func() Direction {
		n := directions[i%4]
		i++
		return n
	}
}

func getNextCoords(x, y int, direction Direction) (int, int) {
	x += direction.x
	y += direction.y
	return x, y
}

func shouldChangeDirection(board [][]int, x, y int, direction Direction) bool {
	x, y = getNextCoords(x, y, direction)
	return x < 0 || y < 0 || x >= len(board) || y >= len(board) || board[y][x] != 0
}

func GetSolved(size int) Board {
	board := make(Board, size)
	for i := range board {
		board[i] = make([]int, size)
	}
	getNextDirection := nextDirectionGenerator()
	direction, x, y := getNextDirection(), 0, 0
	for i := 1; i < size*size; i++ {
		board[y][x] = i
		if shouldChangeDirection(board, x, y, direction) {
			direction = getNextDirection()
		}
		x, y = getNextCoords(x, y, direction)
	}
	return board
}
