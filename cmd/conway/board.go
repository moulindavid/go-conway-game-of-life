package conway

import (
	"fmt"
	"math/rand"
)

// Dead : 0 Alive: 1
type CellState int

const (
	whiteSquare           = "\xF0\x9F\x9F\xAA"
	greenSquare           = "\xF0\x9F\x9F\xA9"
	Dead        CellState = 0
	Alive       CellState = 1
)

type Dimension struct {
	X, Y int
}

type Stats struct {
	GenerationCount int
	Population      int
}

type Board struct {
	Stats
	Dimension
	Cells [][]CellState
}

func NewBoard(x, y int) *Board {
	return &Board{
		Dimension: Dimension{
			x,
			y,
		},
	}
}

func (b *Board) GenerateCells() {
	b.Cells = make([][]CellState, b.X)
	for x := 0; x < b.X; x++ {
		b.Cells[x] = make([]CellState, b.Y)
	}
}

func (b *Board) Random() {
	for i := 0; i < b.X; i++ {
		for j := 0; j < b.Y; j++ {
			b.Cells[i][j] = CellState(rand.Int() % 2)
		}
	}
}

func (b *Board) Draw() {
	for _, row := range b.Cells {
		for _, cell := range row {
			switch {
			case cell == Alive:
				fmt.Printf(greenSquare)
			default:
				fmt.Printf(whiteSquare)
			}
		}
		fmt.Printf("\n")
	}
}

func (b *Board) Alive(x, y int) bool {
	if x < 0 || y < 0 || x >= b.X || y >= b.Y {
		return false
	}
	return b.Cells[x][y] == Alive
}

func (b *Board) CountNeighbours(x, y int) int {
	var neighbours int

	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i == x && j == y {
				continue
			}
			if b.Alive(i, j) {
				neighbours++
			}
		}
	}

	return neighbours
}

func (b *Board) NextGeneration() int {
	nextGrid := make([][]CellState, b.X)
	for x := 0; x < b.X; x++ {
		nextGrid[x] = make([]CellState, b.Y)
	}

	var livingCells int
	for i := 0; i < b.X; i++ {
		for j := 0; j < b.Y; j++ {
			value := b.Cells[i][j]
			neighbours := b.CountNeighbours(i, j)
			if value == Dead && containsInt([]int{3}, neighbours) {
				nextGrid[i][j] = Alive
			} else if value == Alive && !containsInt([]int{2, 3}, neighbours) {
				nextGrid[i][j] = Dead
			} else {
				nextGrid[i][j] = value
			}
			if nextGrid[i][j] == Alive {
				livingCells++
			}
		}
	}

	b.Cells = nextGrid
	return livingCells
}
