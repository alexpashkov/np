package path

import (
	"fmt"
	"github.com/alexpashkov/npuzzle/src/heuristics"
	"github.com/alexpashkov/npuzzle/src/priority-queue"
	"github.com/alexpashkov/npuzzle/src/puzzle"
)

func Find(from, to puzzle.Puzzle, f heuristics.Func) {
	openSet := priority_queue.New(f)
	fmt.Println("Finding path...")
}
