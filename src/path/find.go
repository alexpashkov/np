package path

import (
	"fmt"
	"github.com/alexpashkov/npuzzle/src/heuristics"
	"github.com/alexpashkov/npuzzle/src/priority-queue"
	"github.com/alexpashkov/npuzzle/src/puzzle"
	"github.com/alexpashkov/npuzzle/src/state"
)

func Find(from, to puzzle.Puzzle, f heuristics.Func) {
	openSet := priority_queue.New(f)
	fromState := state.State{Parent: nil, Puzzle: from}
	openSet.Push(&fromState)

	fmt.Println(state.Next(fromState))

	//for openSet.Len() != 0 && !puzzle.Equal(openSet.Peek().(*state.State).Puzzle, to) {
	//
	//	fmt.Println(openSet.Peek())
	//}
}
