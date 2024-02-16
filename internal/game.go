package internal

import (
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	hud   HUD
	world World
	state State
}

func NewGame() (g *Game) {
	g = new(Game)
	g.state.Init()
	g.InitWorld()
	g.hud.Init(HUDOptions{
		Play:    g.Play,
		Pause:   g.Pause,
		ZoomIn:  g.ZoomIn,
		ZoomOut: g.ZoomOut,
		Restart: g.Restart,
		Fast:    g.Fast,
		Slow:    g.Slow,
		Step:    g.Step,
	})
	return
}

func (g *Game) InitWorld() {
	g.world.Init(g.state.Width(), g.state.Height(), g.state.Zoom())
}

func (g *Game) Play() {
	g.state.Play()
	g.UpdateFPS()
}

func (g *Game) Pause() {
	g.state.Pause()
	g.UpdateFPS(1)
}

func (g *Game) PlayPause() {
	g.state.PlayPause()
}

func (g *Game) Restart() {
	g.Pause()
	g.InitWorld()
}

func (g *Game) ZoomIn() {
	if g.state.CanZoomIn() {
		g.state.ZoomIn()
		g.InitWorld()
	}
}

func (g *Game) ZoomOut() {
	if g.state.CanZoomOut() {
		g.state.ZoomOut()
		g.InitWorld()
	}
}

func (g *Game) Fast() {
	if g.state.CanFast() {
		g.state.Fast()
		g.UpdateFPS()
	}
}

func (g *Game) Slow() {
	if g.state.CanSlow() {
		g.state.Slow()
		g.UpdateFPS()
	}
}

func (g *Game) Step() {
	if g.state.Paused() {
		g.world.Update()
	}
}

func (g *Game) UpdateFPS(speed ...int) {
	spd := g.state.Speed()
	if len(speed) >= 1 {
		spd = speed[0]
	}
	ebiten.SetTPS(ebiten.DefaultTPS / spd)
}

func (g *Game) HandleDisableButtons() {
	g.hud.play.GetWidget().Disabled = !g.state.Paused()
	g.hud.pause.GetWidget().Disabled = g.state.Paused()
	g.hud.step.GetWidget().Disabled = !g.state.Paused()
	g.hud.zoomIn.GetWidget().Disabled = !g.state.CanZoomIn()
	g.hud.zoomOut.GetWidget().Disabled = !g.state.CanZoomOut()
	g.hud.fast.GetWidget().Disabled = !g.state.CanFast()
	g.hud.slow.GetWidget().Disabled = !g.state.CanSlow()
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
	g.hud.fps.Label = "FPS: " + strconv.Itoa(int(ebiten.ActualFPS()))
}

func (g *Game) Update() (err error) {
	g.HandleShortcuts()
	g.UpdateFps()
	g.HandleDisableButtons()

	if !g.state.Paused() {
		err = g.world.Update()
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
