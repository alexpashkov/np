package main

import (
	"fmt"
	"os"
	"github.com/alexpashkov/npuzzle/src/puzzle"
)

func main() {
	b, err := puzzle.Read(os.Stdin)
	if err == nil {
		fmt.Println("Puzzle:", b)
		if !puzzle.IsSolvable(b) {
			fmt.Println("Puzzle is not solvable")
		}
	} else {
		fmt.Println("Input error:", err)
	}
}
