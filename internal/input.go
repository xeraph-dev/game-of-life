package internal

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Input struct {
	a     Actions
	shift bool
}

func (i *Input) Init(actions Actions) {
	i.a = actions
}

func (i *Input) handleShift() {
	if inpututil.IsKeyJustPressed(ebiten.KeyShift) {
		i.shift = true
	}
	if inpututil.IsKeyJustReleased(ebiten.KeyShift) {
		i.shift = false
	}
}

func (i *Input) Update() {
	i.handleShift()

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		i.a.PlayPause()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		i.a.Restart()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyMinus) {
		i.a.ZoomOut()
	}

	if i.shift && inpututil.IsKeyJustPressed(ebiten.KeyEqual) {
		i.a.ZoomIn()
	}

	if i.shift && inpututil.IsKeyJustPressed(ebiten.KeyComma) {
		i.a.Slow()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyPeriod) {
		if i.shift {
			i.a.Fast()
		} else {
			i.a.Step()
		}
	}
}
