package cmd

import "github.com/hajimehoshi/ebiten/v2"

// Board 游戏的具体实现
type Board struct {
	cells [][]Cell
}

func NewBoard(dimension int, isRand bool) *Board {
	board := &Board{
		cells: NewCells(dimension, isRand),
	}
	return board
}

func (b *Board) Update() error {
	// 点击生成cell
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		b.flip(x, y)
	}
	b.cellMove()

	return nil
}

// Draw draws the board to the given boardImage.
func (b *Board) Draw(screen *ebiten.Image) {
	for row := 0; row < boardDimension; row++ {
		for col := 0; col < boardDimension; col++ {
			c := b.cells[row][col]
			c.Draw(screen)
		}
	}
}

// cellMove Cell 的移动规则
func (b *Board) cellMove() {
	// 清空
	nextCells := NewCells(boardDimension, false)

	for r := 0; r < boardDimension; r++ {
		for c := 0; c < boardDimension; c++ {
			currCell := &b.cells[r][c]
			neighborsNums := b.countCellNeighbors(currCell)
			state := currCell.state
			// 规则
			if state && (neighborsNums == 2 || neighborsNums == 3) {
				nextCells[r][c].state = true
			} else if !state && neighborsNums == 3 {
				nextCells[r][c].state = true
			} else {
				nextCells[r][c].state = false
			}
		}
	}
	// 更新
	b.cells = nextCells
}

func (b *Board) countCellNeighbors(c *Cell) (neighbors int) {
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			if dr == 0 && dc == 0 {
				continue
			}
			neighbors += b.isNeigh(c.row+dr, c.col+dc)
		}
	}
	return
}

// Returns either a 1 or 0 if the neighbor is either valid and alive or not
func (b *Board) isNeigh(row, col int) int {
	if row < 0 || col < 0 || row >= boardDimension || col >= boardDimension {
		return 0
	}
	if b.cells[row][col].state {
		return 1
	} else {
		return 0
	}
}

// Sets the state of the clicked cell and its neighboring cells to alive
func (b *Board) flip(x int, y int) {
	// Scale the mouse click x and y pixel coordinates to the board dimensions
	row := int(x / scale)
	col := int(y / scale)

	if row >= 0 && col >= 0 && row < boardDimension && col < boardDimension {
		b.cells[row][col].state = true
	}
}
