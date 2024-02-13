package internal

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	hud   HUD
	world World
}

func NewGame() (g *Game) {
	g = new(Game)
	g.world.Init()
	g.hud.Init()
	return
}

func (g *Game) HandleShortcuts() {
	// handle play/pause
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		state.Paused = !state.Paused
	}

	// handle restart
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		state.Paused = true
		g.world.Init()

	}

	// handle zoom out
	if state.Paused && inpututil.IsKeyJustPressed(ebiten.KeyComma) && state.Zoom > 1 {
		state.Zoom--
		g.world.Init()
	}

	// handle zoom in
	if state.Paused && inpututil.IsKeyJustPressed(ebiten.KeyPeriod) && state.Zoom < 5 {
		state.Zoom++
		g.world.Init()
	}
}

func (g *Game) Update() (err error) {
	g.HandleShortcuts()
	if !state.Paused {
		err = g.world.Update()
	}
	g.hud.Update()
	return
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.world.Draw()
	if screen.Bounds().Dx()*screen.Bounds().Dy()*4 == len(g.world.Pixels) {
		screen.WritePixels(g.world.Pixels)
	}
	g.hud.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
