package internal

import (
	"game-of-life/internal/assets"
	"game-of-life/internal/custom"
	"image/color"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font/gofont/goregular"
)

type HUD struct {
	ui      *ebitenui.UI
	play    *widget.Button
	pause   *widget.Button
	zoomIn  *widget.Button
	zoomOut *widget.Button
	fast    *widget.Button
	slow    *widget.Button
	step    *widget.Button
	restart *widget.Button
	fps     *widget.Label
}

type HUDOptions struct {
	Play    func()
	Pause   func()
	ZoomIn  func()
	ZoomOut func()
	Fast    func()
	Slow    func()
	Step    func()
	Restart func()
}

func (h *HUD) Init(opts HUDOptions) {
	h.ui = new(ebitenui.UI)

	h.ui.Container = widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(
			widget.AnchorLayoutOpts.Padding(widget.NewInsetsSimple(4)),
		)),
	)

	h.play = custom.NewIconButton(assets.PlayIcon, opts.Play, custom.NewShortcut("play the world", assets.PlayIcon))
	h.pause = custom.NewIconButton(assets.PauseIcon, opts.Pause, custom.NewShortcut("pause the world", assets.PlayIcon))
	h.zoomIn = custom.NewIconButton(assets.PlusIcon, opts.ZoomIn, custom.NewShortcut("increase cell's size", assets.PlayIcon))
	h.zoomOut = custom.NewIconButton(assets.MinusIcon, opts.ZoomOut, custom.NewShortcut("decrease cell's size", assets.PlayIcon))
	h.fast = custom.NewIconButton(assets.FastIcon, opts.Fast, custom.NewShortcut("increase generation speed", assets.PlayIcon))
	h.slow = custom.NewIconButton(assets.SlowIcon, opts.Slow, custom.NewShortcut("decrease generation speed", assets.PlayIcon))
	h.step = custom.NewIconButton(assets.StepIcon, opts.Step, custom.NewShortcut("advance one generation", assets.PlayIcon))
	h.restart = custom.NewIconButton(assets.RestartIcon, opts.Restart, custom.NewShortcut("regenerate the world", assets.KeyRIcon))

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
	buttonsContainer.AddChild(h.step)
	buttonsContainer.AddChild(h.pause)
	buttonsContainer.AddChild(h.zoomIn)
	buttonsContainer.AddChild(h.zoomOut)
	buttonsContainer.AddChild(h.slow)
	buttonsContainer.AddChild(h.fast)

	h.ui.Container.AddChild(buttonsContainer)

	fpsContainer := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(
			widget.AnchorLayoutOpts.Padding(widget.NewInsetsSimple(6)),
		)),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(
			widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionStart,
			},
		)),
	)

	ttfFont, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic(err)
	}
	face := truetype.NewFace(ttfFont, &truetype.Options{
		Size: 18,
	})

	h.fps = widget.NewLabel(
		widget.LabelOpts.Text("FPS", face, &widget.LabelColor{
			Idle: color.White,
		}),
	)

	fpsContainer.AddChild(h.fps)

	h.ui.Container.AddChild(fpsContainer)
}

func (h *HUD) Update() {
	h.ui.Update()
}

func (h *HUD) Draw(screen *ebiten.Image) {
	h.ui.Draw(screen)
}
