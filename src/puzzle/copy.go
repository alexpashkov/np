package puzzle

func Copy(src Puzzle) Puzzle {
	dst := make(Puzzle, src.Size())
	for y := range src {
		dst[y] = make(Row, src.Size())
		for x, t := range src[y] {
			dst[y][x] = t
		}
	}
	return dst
}
