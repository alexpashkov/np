package state

import (
	"github.com/alexpashkov/npuzzle/src/puzzle"
	"fmt"
)

func Next(currState State) (nextStates []*State) {
	moves := [4]puzzle.TileCoords{
		{
			X: 1, Y: 0,
		},
		{
			X: 0, Y: 1,
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
			Puzzle: puzzle.Copy(currState.Puzzle),
		}
		fmt.Println(nextState)
		emptyTileCoords := currState.EmptyTileCoords()
		fmt.Println(emptyTileCoords)
		err := puzzle.Swap(nextState.Puzzle, emptyTileCoords, puzzle.TileCoords{
			X: emptyTileCoords.X + m.X,
			Y: emptyTileCoords.Y + m.Y,
		})
		if err != nil {
			fmt.Println(err)
		}
		if currState.Parent != nil && nextState.Id() != currState.Parent.Id() {
			nextStates = append(nextStates, &nextState)
		}
	}
	return
}
