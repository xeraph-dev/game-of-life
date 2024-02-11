package internal

import "github.com/hajimehoshi/ebiten/v2"

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

func (self *Game) Update() (err error) {
	err = self.world.Update()
	return
}

func (self *Game) Draw(screen *ebiten.Image) {
	self.world.Draw()
	pixels := self.world.Pixels()
	screen.WritePixels(pixels)
}

func (self *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
