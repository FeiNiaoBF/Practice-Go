package cmd

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// Game represents a game state.
type Game struct {
	board      *Board
	boardImage *ebiten.Image
}

// NewGame generates a new Game object.
func NewGame(isRand bool) *Game {
	g := &Game{}
	g.board = NewBoard(boardDimension, isRand)
	return g
}

// Layout implements ebiten.Game's Layout.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

// Update the current game state.
func (g *Game) Update() error {
	if err := g.board.Update(); err != nil {
		return err
	}
	return nil
}

// Draw the current game state to the given screen.
func (g *Game) Draw(screen *ebiten.Image) {
	if g.boardImage == nil {
		w, h := ScreenWidth, ScreenHeight
		g.boardImage = ebiten.NewImage(w, h)
	}
	screen.Fill(CellDead)
	g.board.Draw(g.boardImage)
	screen.DrawImage(g.boardImage, &ebiten.DrawImageOptions{})
}
