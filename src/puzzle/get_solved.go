package puzzle

type Direction struct {
	x, y int
}

func nextDirectionGenerator() func() Direction {
	dirs, i := [4]Direction{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}, 0
	return func() Direction {
		n := dirs[i%4]
		i++
		return n
	}
}

func getNextCoords(x, y int, dir Direction) (int, int) {
	x += dir.x
	y += dir.y
	return x, y
}

func shouldChangeDirection(p Puzzle, x, y int, dir Direction) bool {
	x, y = getNextCoords(x, y, dir)
	return x < 0 || y < 0 || x >= len(p) || y >= len(p) || p.Tile(TileCoords{X: x, Y: y}).Val() != 0
}

var solvedPuzzles = make(map[int]Puzzle)

func GetSolved(size int) Puzzle {
	p, ok := solvedPuzzles[size]
	if ok {
		return p
	}
	p = make(Puzzle, size)
	for i := range p {
		p[i] = make(Row, size)
	}
	getNextDirection := nextDirectionGenerator()
	direction, x, y := getNextDirection(), 0, 0
	for i := 1; i < size*size; i++ {
		p[y][x] = Tile(i)
		if shouldChangeDirection(p, x, y, direction) {
			direction = getNextDirection()
		}
		x, y = getNextCoords(x, y, direction)
	}
	solvedPuzzles[size] = p
	return p
}
