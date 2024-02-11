package internal

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	ScreenWidth  = 640
	ScreenHeight = 480
)

type Game struct {
	world World
}

func NewGame() (self *Game) {
	self = new(Game)
	self.world.Init()
	return
}

func (self *Game) HandlePaused() {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		state.paused = !state.paused
	}

}

func (self *Game) HandleRestart() {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		self.world.Init()
		state.paused = false
	}
}

func (self *Game) HandleZoomOut() {
	if state.paused && inpututil.IsKeyJustPressed(ebiten.KeyMinus) && state.zoom > 0 {
		state.zoom--
		self.world.Init()
	}
}

func (self *Game) HandleZoomIn() {
	if state.paused && inpututil.IsKeyJustPressed(ebiten.KeyShift) && inpututil.IsKeyJustPressed(ebiten.KeyEqual) && state.zoom < 5 {
		state.zoom++
		self.world.Init()
	}
}

func (self *Game) Update() (err error) {
	self.HandlePaused()
	self.HandleRestart()
	self.HandleZoomOut()
	self.HandleZoomIn()

	if !state.paused {
		err = self.world.Update()
	}
	return
}

func (self *Game) Draw(screen *ebiten.Image) {
	self.world.Draw()
	pixels := self.world.Pixels()
	screen.WritePixels(pixels)
}

func (self *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth / state.zoom, outsideHeight / state.zoom
}
