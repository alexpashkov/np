package board

type Tile int

func (t Tile) Val() int {
	return int(t)
}

type Row []Tile
type Board []Row

func (b Board) Size() int {
	return len(b)
}

func (b Board) Tile(x, y int) Tile {
	return b[y][x]
}

func (b Board) ForEach(fn func(t Tile, x, y int)) {
	for y := range b {
		for x, t := range b[y] {
			fn(t, x, y)
		}
	}
}
