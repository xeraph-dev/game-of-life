package internal

import (
	"fmt"

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
	g.hud.Init(HUDOptions{
		Play:    g.Play,
		Pause:   g.Pause,
		ZoomIn:  g.ZoomIn,
		ZoomOut: g.ZoomOut,
		Restart: g.Restart,
	})
	return
}

func (g *Game) Play() {
	state.Paused = false
}

func (g *Game) Pause() {
	state.Paused = true
}

func (g *Game) PlayPause() {
	state.Paused = !state.Paused
}

func (g *Game) Restart() {
	g.Pause()
	g.world.Init()
}

func (g *Game) ZoomIn() {
	if state.Paused && state.Zoom < 5 {
		state.Zoom++
		g.world.Init()
	}
}

func (g *Game) ZoomOut() {
	if state.Paused && state.Zoom > 1 {
		state.Zoom--
		g.world.Init()
	}
}

func (g *Game) HandleDisableButtons() {
	g.hud.play.GetWidget().Disabled = !state.Paused
	g.hud.pause.GetWidget().Disabled = state.Paused
	g.hud.zoomIn.GetWidget().Disabled = !state.Paused || state.Zoom >= 5
	g.hud.zoomOut.GetWidget().Disabled = !state.Paused || state.Zoom <= 1
}

func (g *Game) HandleShortcuts() {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		g.PlayPause()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		g.Restart()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyComma) {
		g.ZoomOut()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyPeriod) {
		g.ZoomIn()
	}
}

func (g *Game) UpdateFps() {
	g.hud.fps.Label = fmt.Sprintf("FPS: %d", int(ebiten.ActualFPS()))
}

func (g *Game) Update() (err error) {
	g.HandleShortcuts()

	if !state.Paused {
		err = g.world.Update()
	}

	g.HandleDisableButtons()
	g.UpdateFps()
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
