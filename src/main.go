package main

import (
	"flag"
	"fmt"
	"github.com/alexpashkov/npuzzle/src/heuristics"
	"github.com/alexpashkov/npuzzle/src/path"
	"github.com/alexpashkov/npuzzle/src/puzzle"
	"log"
	"os"
)

func main() {
	heuristicsName := flag.String("h", "manhattan", "specify heuristics")
	flag.Parse()
	heuristicsFn, ok := heuristics.Funcs[*heuristicsName]

	if !ok {
		log.Fatalln("Invalid heuristics function name")
	}

	initialPuzzle, err := puzzle.Read(os.Stdin)
	if err == nil {
		fmt.Println("Puzzle:", initialPuzzle)
		solvedPuzzle := puzzle.GetSolved(initialPuzzle.Size())
		if !path.CanBeReached(initialPuzzle, solvedPuzzle) {
			log.Fatalln("Puzzle is not solvable")
		}
		path.Find(initialPuzzle, solvedPuzzle, heuristicsFn)
	} else {
		log.Fatalln("Input error:", err)
	}
}
