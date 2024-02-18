package internal

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	hud   HUD
	world World
	input Input
}

func NewGame() (g *Game) {
	g = new(Game)
	state.Init()
	g.InitWorld()
	g.hud.Init(g)
	g.input.Init(g)
	return
}

func (g *Game) InitWorld() {
	g.world.Init()
}

func (g *Game) Play() {
	state.Play()
	g.UpdateFPS()
}

func (g *Game) Pause() {
	state.Pause()
	g.UpdateFPS(1)
}

func (g *Game) PlayPause() {
	state.PlayPause()
}

func (g *Game) Restart() {
	g.Pause()
	g.InitWorld()
}

func (g *Game) ZoomIn() {
	if state.CanZoomIn() {
		state.ZoomIn()
		g.InitWorld()
	}
}

func (g *Game) ZoomOut() {
	if state.CanZoomOut() {
		state.ZoomOut()
		g.InitWorld()
	}
}

func (g *Game) Fast() {
	if state.CanFast() {
		state.Fast()
		g.UpdateFPS()
	}
}

func (g *Game) Slow() {
	if state.CanSlow() {
		state.Slow()
		g.UpdateFPS()
	}
}

func (g *Game) Step() {
	if state.Paused() {
		g.world.Update()
	}
}

func (g *Game) UpdateFPS(speed ...int) {
	spd := state.Speed()
	if len(speed) >= 1 {
		spd = speed[0]
	}
	ebiten.SetTPS(ebiten.DefaultTPS / spd)
}

func (g *Game) Update() (err error) {
	g.input.Update()
	if !state.Paused() {
		g.world.Update()
	}
	g.hud.Update()
	return
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.world.Draw()
	screen.WritePixels(g.world.Pixels)
	g.hud.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
