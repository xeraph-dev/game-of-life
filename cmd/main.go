package main

import (
	"game-of-life/internal"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	width, height, zoom := internal.InitialScreenWidth, internal.InitialScreenHeight, internal.InitialZoom

	ebiten.SetWindowSize(width, height)
	ebiten.SetWindowTitle("Game of Life")

	if err := ebiten.RunGame(internal.NewGame(width, height, zoom)); err != nil {
		panic(err)
	}
}
