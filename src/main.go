package main

import (
	"fmt"
	"github.com/alexpashkov/npuzzle/src/puzzle"
	"log"
	"os"
)

func main() {
	p, err := puzzle.Read(os.Stdin)
	if err == nil {
		fmt.Println("Puzzle:", p)
		sp := puzzle.GetSolved(p.Size())
		if !puzzle.IsSolvable(p, sp) {
			log.Fatalln("Puzzle is not solvable")
		}
	} else {
		log.Fatalln("Input error:", err)
	}
}
