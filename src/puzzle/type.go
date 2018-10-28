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

func (p Puzzle) Coords(t Tile) (x, y int) {
	p.ForEach(func(ct Tile, cx, cy int) (shellContinue bool) {
		if ct.Val() == t.Val() {
			x, y = cx, cy
			return false
		}
		return true
	})
	return
}

// Iterates over puzzle tiles. Iteration will stop if provided function returns false
func (p Puzzle) ForEach(fn func(t Tile, x, y int) (shellContinue bool)) {
	for y := range p {
		for x, t := range p[y] {
			shellStop := !fn(t, x, y)
			if shellStop {
				return
			}
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
