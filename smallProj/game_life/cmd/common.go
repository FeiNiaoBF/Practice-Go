package cmd

import "image/color"

// game windows
const (
	ScreenWidth    int = 400
	ScreenHeight   int = 400
	boardDimension int = 100
	scale          int = ScreenWidth / boardDimension
)

// color
var (
	// cell color
	CellLive = color.Black
	CellDead = color.White
)
