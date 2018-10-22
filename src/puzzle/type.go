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

func CoordsFinder(p Puzzle) func(t Tile) (x, y int) {
	type Coords struct{ x, y int }
	knownCoords := make(map[Tile]Coords)
	cy, cx := 0, 0
	return func(t Tile) (x, y int) {
		coords, ok := knownCoords[t]
		if ok {
			return coords.x, coords.y
		}
		for ; cy < p.Size(); cy++ {
			for ; cx < p.Size(); cx++ {
				tile := p[cy][cx]
				knownCoords[tile] = Coords{x: cx, y: cy}
				if tile.Val() == t.Val() {
					break
				}
			}
		}
		return cx, cy
	}
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
