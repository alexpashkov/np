package state

import "github.com/alexpashkov/npuzzle/src/puzzle"

func Next(currState State) (nextStates []*State) {
	moves := [4]puzzle.TileCoords{
		{
			X: 1, Y: 0,
		},
		{
			X: 0, Y: -1,
		},
		{
			X: -1, Y: 0,
		},
		{
			X: 0, Y: -1,
		},
	}
	for _, m := range moves {

		nextState := State{
			Parent: &currState,
		}
		copy(currState.Puzzle, nextState.Puzzle)
		emptyTileCoords := currState.EmptyTileCoords()
		puzzle.Swap(nextState.Puzzle, emptyTileCoords, puzzle.TileCoords{
			X: emptyTileCoords.X + m.X,
			Y: emptyTileCoords.Y + m.Y,
		})
		if currState.Parent != nil && nextState.Id() != currState.Parent.Id() {
			nextStates = append(nextStates, &nextState)
		}
	}
	return
}
