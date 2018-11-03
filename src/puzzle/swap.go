package puzzle

import (
	"fmt"
)

func outOfBounds(p Puzzle, c TileCoords) bool {
	if c.X >= p.Size() || c.Y >= p.Size() || c.X < 0 || c.Y < 0 {
		return true
	}
	return false
}

func Swap(p Puzzle, a, b TileCoords) (err error) {
	if outOfBounds(p, a) {
		return fmt.Errorf("invalid coords: %v", a)
	}
	if outOfBounds(p, b) {
		return fmt.Errorf("invalid coords: %v", b)
	}
	p[a.Y][a.X] = p[b.Y][b.X]
	return
}
