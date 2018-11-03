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
		puzzle.Swap(nextState.Puzzle, currState.EmptyTile(), m)
		if nextState.Id() != currState.Parent.Id() {
			nextStates = append(nextStates, &nextState)
		}
	}
	return
}
