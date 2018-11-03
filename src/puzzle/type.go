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

func (p Puzzle) Coords(t Tile) (x, y int, ok bool) {
	for y := range p {
		for x, ct := range p[y] {
			if ct.Val() == t.Val() {
				return x, y, true
			}
		}
	}
	return 0, 0, false
}

// Iterates over puzzle tiles. Iteration will stop if provided function returns false
func (p Puzzle) ForEach(fn func(t Tile, x, y int)) {
	for y := range p {
		for x, t := range p[y] {
			fn(t, x, y)
		}
	}
}
