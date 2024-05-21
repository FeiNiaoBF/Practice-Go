package cmd

import (
	"github.com/hajimehoshi/ebiten"
)

// Game represents a game state.
type Game struct {
	board      *Board
	boardImage *ebiten.Image
}

// NewGame generates a new Game object.
func NewGame(isRand bool) *Game {

	g := &Game{
		board: NewBoard(boardDimension, isRand),
	}
	return g
}

// Layout implements ebiten.Game's Layout.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

// Update the current game state.
func (g *Game) Update(screen *ebiten.Image) error {
	if err := g.board.Update(); err != nil {
		return err
	}
	return nil
}

// Draw the current game state to the given screen.
func (g *Game) Draw(screen *ebiten.Image) {
	if g.boardImage == nil {
		w, h := ScreenWidth, ScreenHeight
		g.boardImage, _ = ebiten.NewImage(w, h, ebiten.FilterDefault)
	}
	screen.Fill(CellDead)
	g.board.Draw(g.boardImage)
	screen.DrawImage(g.boardImage, &ebiten.DrawImageOptions{})

}
