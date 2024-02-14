package internal

import (
	"game-of-life/internal/assets"
	"game-of-life/internal/custom"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

type HUD struct {
	ui      *ebitenui.UI
	play    *widget.Button
	pause   *widget.Button
	zoomIn  *widget.Button
	zoomOut *widget.Button
	restart *widget.Button
}

type HUDOptions struct {
	Play    func()
	Pause   func()
	ZoomIn  func()
	ZoomOut func()
	Restart func()
}

func (h *HUD) Init(opts HUDOptions) {
	h.ui = new(ebitenui.UI)

	h.ui.Container = widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(
			widget.AnchorLayoutOpts.Padding(widget.NewInsetsSimple(4)),
		)),
	)

	h.play = custom.NewIconButton(assets.PlayIcon, func(args *widget.ButtonClickedEventArgs) {
		opts.Play()
	})

	h.pause = custom.NewIconButton(assets.PauseIcon, func(args *widget.ButtonClickedEventArgs) {
		opts.Pause()
	})

	h.zoomIn = custom.NewIconButton(assets.PlusIcon, func(args *widget.ButtonClickedEventArgs) {
		opts.ZoomIn()
	})

	h.zoomOut = custom.NewIconButton(assets.MinusIcon, func(args *widget.ButtonClickedEventArgs) {
		opts.ZoomOut()
	})

	h.restart = custom.NewIconButton(assets.RestartIcon, func(args *widget.ButtonClickedEventArgs) {
		opts.Restart()
	})

	buttonsContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
			widget.RowLayoutOpts.Spacing(2),
		)),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(
			widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
			},
		)),
	)

	buttonsContainer.AddChild(h.restart)
	buttonsContainer.AddChild(h.play)
	buttonsContainer.AddChild(h.pause)
	buttonsContainer.AddChild(h.zoomIn)
	buttonsContainer.AddChild(h.zoomOut)

	h.ui.Container.AddChild(buttonsContainer)
}

func (h *HUD) Update() {
	h.ui.Update()
}

func (h *HUD) Draw(screen *ebiten.Image) {
	h.ui.Draw(screen)
}
