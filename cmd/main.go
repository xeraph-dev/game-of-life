package main

import (
	"game-of-life/internal"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(internal.ScreenWidth, internal.ScreenHeight)
	ebiten.SetWindowTitle("Game of Life")

	if err := ebiten.RunGame(internal.NewGame()); err != nil {
		panic(err)
	}
}
