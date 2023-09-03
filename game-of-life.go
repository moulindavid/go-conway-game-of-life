package main

import (
	"fmt"
	"moulindavid/game-of-life/cmd/conway"
)

func main() {
	Board := conway.NewBoard(5, 5)
	Board.GenerateCells()
	Board.Random()
	fmt.Println(Board.Cells)
}
