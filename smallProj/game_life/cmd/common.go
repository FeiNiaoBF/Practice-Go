package cmd

import "image/color"

// game windows
const (
	ScreenWidth    int = 600
	ScreenHeight   int = 600
	boardDimension int = 100
	scale          int = ScreenWidth / boardDimension
)

// color
var (
	// cell color
	CellLive = color.Black
	CellDead = color.White
)
