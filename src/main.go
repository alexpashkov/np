package main

import (
	"fmt"
	"log"
	"os"
	"github.com/alexpashkov/npuzzle/src/puzzle"
	"github.com/alexpashkov/npuzzle/src/path"
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
