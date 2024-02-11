package main

import (
	. "game-of-life/internal"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Game of Life")

	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
