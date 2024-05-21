package main

import (
	"flag"
	"log"

	"github.com/FeiNiaoBF/Practice-Go/smallProj/gamelife/cmd"
	"github.com/hajimehoshi/ebiten"
)

var setRand bool

func init() {
	flag.BoolVar(&setRand, "r", false, "use random generation")
}

func main() {
	flag.Parse()

	gameLife := cmd.NewGame(setRand)
	ebiten.SetMaxTPS(30)
	ebiten.SetWindowSize(cmd.ScreenWidth, cmd.ScreenHeight)
	ebiten.SetWindowTitle("Game of Life")
	if err := ebiten.RunGame(gameLife); err != nil {
		log.Fatal(err)
	}
}
