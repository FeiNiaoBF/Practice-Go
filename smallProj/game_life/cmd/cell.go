package cmd

import (
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
)

// Cell 是点/块
type Cell struct {
	row   int
	col   int
	scale int
	// cell is alive or dead
	state bool
}

func newCell(x, y int) *Cell {
	return &Cell{
		row:   x,
		col:   y,
		scale: scale,
		state: false,
	}
}

// 初始cells
func NewCells(dimension int, isRand bool) [][]Cell {
	cells := make([][]Cell, dimension)

	// rand cell
	seed := time.Now().UnixNano()
	random := rand.New(rand.NewSource(seed))
	for r := 0; r < dimension; r++ {
		cells[r] = make([]Cell, dimension)
		for c := 0; c < dimension; c++ {
			cells[r][c] = *newCell(r, c)
			if isRand && random.Float32() < 0.5 {
				cells[r][c].state = true
			}
		}
	}
	return cells
}

func (c *Cell) Draw(board *ebiten.Image) {
	var colorCell color.Gray16

	if c.state {
		// If alive, the cell fill is white
		colorCell = CellLive
	} else {
		// Otherwise, it is black (i.e. dead)
		colorCell = CellDead
	}

	ebitenutil.DrawRect(board,
		float64(c.row*scale),
		float64(c.col*scale),
		float64(scale),
		float64(scale),
		colorCell,
	)
}
