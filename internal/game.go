package internal

import (
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	hud           HUD
	world         World
	width, height int
	zoom          int
	paused        bool
	speed         int
}

func NewGame() (g *Game) {
	g = new(Game)
	g.width, g.height, g.zoom, g.speed = InitialScreenWidth, InitialScreenHeight, InitialZoom, InitialSpeed
	g.paused = true
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
	g.world.Init(g.width, g.height, g.zoom)
}

func (g *Game) Play() {
	g.paused = false
	g.UpdateFPS()
}

func (g *Game) Pause() {
	g.paused = true
	g.UpdateFPS(ebiten.DefaultTPS)
}

func (g *Game) PlayPause() {
	g.paused = !g.paused
}

func (g *Game) Restart() {
	g.Pause()
	g.InitWorld()
}

func (g *Game) ZoomIn() {
	if g.paused && g.zoom < MaxZoom {
		g.zoom++
		g.InitWorld()
	}
}

func (g *Game) ZoomOut() {
	if g.paused && g.zoom > MinZoom {
		g.zoom--
		g.InitWorld()
	}
}

func (g *Game) Fast() {
	if !g.paused && g.speed > MinSpeed {
		g.speed--
		g.UpdateFPS()
	}
}

func (g *Game) Slow() {
	if !g.paused && g.speed < MaxSpeed {
		g.speed++
		g.UpdateFPS()
	}
}

func (g *Game) Step() {
	if g.paused {
		g.world.Update()
	}
}

func (g *Game) UpdateFPS(speed ...int) {
	spd := g.speed
	if len(speed) >= 1 {
		spd = speed[0]
	}
	ebiten.SetTPS(ebiten.DefaultTPS / spd)
}

func (g *Game) HandleDisableButtons() {
	g.hud.play.GetWidget().Disabled = !g.paused
	g.hud.pause.GetWidget().Disabled = g.paused
	g.hud.zoomIn.GetWidget().Disabled = !g.paused || g.zoom >= MaxZoom
	g.hud.zoomOut.GetWidget().Disabled = !g.paused || g.zoom <= MinZoom
	g.hud.fast.GetWidget().Disabled = g.paused || g.speed <= MinSpeed
	g.hud.slow.GetWidget().Disabled = g.paused || g.speed >= MaxSpeed
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

	if !g.paused {
		err = g.world.Update()
	}

	g.HandleDisableButtons()
	g.UpdateFps()
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
