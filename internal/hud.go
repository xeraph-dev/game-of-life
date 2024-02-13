package internal

import (
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

type HUD struct {
	ui *ebitenui.UI
}

func (h *HUD) Init() {
	h.ui = new(ebitenui.UI)
	h.ui.Container = widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
			widget.RowLayoutOpts.Spacing(4),
		)),
	)

	// h.playButton = NewPlayButton()

	// h.ui.Container.AddChild(h.playButton)
}

func (h *HUD) Update() {
	h.ui.Update()
}

func (h *HUD) Draw(screen *ebiten.Image) {
	h.ui.Draw(screen)
}
