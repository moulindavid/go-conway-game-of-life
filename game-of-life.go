package main

import (
	"fmt"
	"moulindavid/game-of-life/cmd/conway"
	"time"
)

const ansiEscapeCode = "\033c\x0c"

func main() {
	fmt.Println(ansiEscapeCode)
	Board := conway.NewBoard(25, 25)
	Board.GenerateCells()
	Board.Random()
	for {
		Board.Draw()
		Board.NextGeneration()
		time.Sleep(200 * time.Millisecond)
		fmt.Println(ansiEscapeCode)
	}
}
