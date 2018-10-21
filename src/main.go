package main

import (
	"fmt"
	"github.com/alexpashkov/npuzzle/src/path"
	"github.com/alexpashkov/npuzzle/src/puzzle"
	"log"
	"os"
)

func main() {
	initialPuzzle, err := puzzle.Read(os.Stdin)
	if err == nil {
		fmt.Println("Puzzle:", initialPuzzle)
		solvedPuzzle := puzzle.GetSolved(initialPuzzle.Size())
		if !puzzle.IsReachable(initialPuzzle, solvedPuzzle) {
			log.Fatalln("Puzzle is not solvable")
		}
		path.Find(initialPuzzle, solvedPuzzle)
	} else {
		log.Fatalln("Input error:", err)
	}
}
