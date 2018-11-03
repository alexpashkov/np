package puzzle

type Tile int
type TileCoords struct {
	X, Y int
}

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

func (p Puzzle) Find(t Tile) TileCoords {
	for y := range p {
		for x, ct := range p[y] {
			if ct.Val() == t.Val() {
				return TileCoords{X: x, Y: y}
			}
		}
	}
	return TileCoords{X: -1, Y: -1}
}

// Iterates over puzzle tiles. Iteration will stop if provided function returns false
func (p Puzzle) ForEach(fn func(t Tile, x, y int)) {
	for y := range p {
		for x, t := range p[y] {
			fn(t, x, y)
		}
	}
}

func Equal(a, b Puzzle) bool {
	if a == nil || b == nil {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for y := range a {
		if len(a[y]) != len(b[y]) {
			return false
		}
		for x := range a[y] {
			if a[y][x] != b[y][x] {
				return false
			}
		}
	}
	return true
}
