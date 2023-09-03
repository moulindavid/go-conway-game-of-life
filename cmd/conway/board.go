package conway

import "math/rand"

// Dead : 0 Alive: 1
type CellState int

const (
	Dead  CellState = iota
	Alive CellState = 1
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
