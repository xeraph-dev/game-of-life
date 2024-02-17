package internal

import (
	"game-of-life/internal/assets"
	"game-of-life/internal/custom"
	"image/color"
	"strconv"

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

func (h *HUD) Init(actions Actions) {
	h.ui = new(ebitenui.UI)

	h.ui.Container = widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(
			widget.AnchorLayoutOpts.Padding(widget.NewInsetsSimple(4)),
		)),
	)

	h.play = custom.NewIconButton(assets.PlayIcon, actions.Play, custom.NewShortcut("Play the world", assets.KeySpaceIcon))
	h.pause = custom.NewIconButton(assets.PauseIcon, actions.Pause, custom.NewShortcut("Pause the world", assets.KeySpaceIcon))
	h.zoomIn = custom.NewIconButton(assets.PlusIcon, actions.ZoomIn, custom.NewShortcut("Increase cell's size", assets.KeyPlusIcon))
	h.zoomOut = custom.NewIconButton(assets.MinusIcon, actions.ZoomOut, custom.NewShortcut("Decrease cell's size", assets.KeyMinusIcon))
	h.fast = custom.NewIconButton(assets.FastIcon, actions.Fast, custom.NewShortcut("Increase generation speed", assets.KeyGreaterIcon))
	h.slow = custom.NewIconButton(assets.SlowIcon, actions.Slow, custom.NewShortcut("Decrease generation speed", assets.KeyLowerIcon))
	h.step = custom.NewIconButton(assets.StepIcon, actions.Step, custom.NewShortcut("Advance one generation", assets.KeyPeriodIcon))
	h.restart = custom.NewIconButton(assets.RestartIcon, actions.Restart, custom.NewShortcut("Regenerate the world", assets.KeyRIcon))

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

func (h *HUD) updateFps() {
	h.fps.Label = "FPS: " + strconv.Itoa(int(ebiten.ActualFPS()))
}

func (h *HUD) handleDisableButtons() {
	h.play.GetWidget().Disabled = !state.Paused()
	h.pause.GetWidget().Disabled = state.Paused()
	h.step.GetWidget().Disabled = !state.Paused()
	h.zoomIn.GetWidget().Disabled = !state.CanZoomIn()
	h.zoomOut.GetWidget().Disabled = !state.CanZoomOut()
	h.fast.GetWidget().Disabled = !state.CanFast()
	h.slow.GetWidget().Disabled = !state.CanSlow()
}

func (h *HUD) Update() {
	h.updateFps()
	h.handleDisableButtons()
	h.ui.Update()
}

func (h *HUD) Draw(screen *ebiten.Image) {
	h.ui.Draw(screen)
}
