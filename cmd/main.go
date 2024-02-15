package main

import (
	"game-of-life/internal"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(internal.InitialScreenWidth, internal.InitialScreenHeight)
	ebiten.SetWindowTitle("Game of Life")
	ebiten.SetTPS(ebiten.DefaultTPS / internal.InitialSpeed)

	if err := ebiten.RunGame(internal.NewGame()); err != nil {
		panic(err)
	}
}
