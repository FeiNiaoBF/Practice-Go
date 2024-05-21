package cmd

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	board *Board
	image *ebiten.Image
}

func NewGame(isRand bool) *Game {
	g := &Game{
		board: NewBoard(boardDimension, isRand),
	}
	return g
}

// Update game 的更新频率
func (g *Game) Update() error {
	if err := g.board.Update(); err != nil {
		return err
	}
	return nil
}

// Draw 画图---需要显示画面
func (g *Game) Draw(screen *ebiten.Image) {
	g.board.Draw(screen)
	screen.DrawImage(g.image, &ebiten.DrawImageOptions{})
}

// Layout 大小
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
