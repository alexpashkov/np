package puzzle

type Tile int

func (t Tile) Val() int {
	return int(t)
}

type Row []Tile
type Puzzle []Row

func (p Puzzle) Size() int {
	return len(p)
}

func (p Puzzle) Tile(x, y int) Tile {
	return p[y][x]
}

func (p Puzzle) ForEach(fn func(t Tile, x, y int)) {
	for y := range p {
		for x, t := range p[y] {
			fn(t, x, y)
		}
	}
}
