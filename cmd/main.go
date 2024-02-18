package main

import (
	"game-of-life/internal"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	r := internal.InitialResolution

	ebiten.SetWindowSize(r.Width, r.Height)
	ebiten.SetWindowTitle(internal.Title)
	ebiten.SetTPS(ebiten.DefaultTPS / internal.InitialSpeed)

	if err := ebiten.RunGame(internal.NewGame()); err != nil {
		panic(err)
	}
}
