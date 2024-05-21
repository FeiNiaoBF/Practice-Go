package cmd

import (
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// Cell 是点/块
type Cell struct {
	image *ebiten.Image
	row   int
	col   int
	scale int
	// cell is alive or dead
	state bool
	color color.Color
}

func newCell(x, y int) *Cell {
	img := ebiten.NewImage(1, 1) // 创建一个1x1的图像
	img.Fill(CellDead)
	img.Fill(CellDead)
	return &Cell{
		image: img,
		row:   x,
		col:   y,
		scale: scale,
		state: false,
		color: CellDead,
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
			if isRand {
				cells[r][c].state = random.Float32() < 0.5
			}
		}
	}
	return cells
}

func (c *Cell) Draw(board *ebiten.Image) {
	if c.state {
		c.color = CellLive
	} else {
		c.color = CellDead
	}
	vector.DrawFilledRect(
		board,
		float32(c.row*c.scale),
		float32(c.col*c.scale),
		float32(c.scale),
		float32(c.scale),
		c.color,
		true,
	)
}
