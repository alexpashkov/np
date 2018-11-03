package puzzle

func Swap(p Puzzle, a, b TileCoords) {
	p[a.Y][a.X] = p[b.Y][b.X]
}
