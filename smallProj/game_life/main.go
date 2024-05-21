package main

import (
	"flag"
	"github.com/FeiNiaoBF/Practice-Go/smallProj/gamelife/cmd"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

var (
	setRand bool
)

func init() {
	flag.BoolVar(&setRand, "r", true, "use random generation")
}

func main() {
	flag.Parse()

	gameLife := cmd.NewGame(setRand)

	ebiten.SetWindowSize(cmd.ScreenWidth, cmd.ScreenHeight)
	ebiten.SetWindowTitle("Game of Life")
	if err := ebiten.RunGame(gameLife); err != nil {
		log.Fatal(err)
	}
}
