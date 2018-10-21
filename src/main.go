package main

import (
	"fmt"
	"github.com/alexpashkov/npuzzle/src/puzzle"
	"log"
	"os"
)

func main() {
	b, err := puzzle.Read(os.Stdin)
	if err == nil {
		fmt.Println("Puzzle:", b)
		if !puzzle.IsSolvable(b) {
			log.Fatalln("Puzzle is not solvable")
		}
	} else {
		log.Fatalln("Input error:", err)
	}
}
