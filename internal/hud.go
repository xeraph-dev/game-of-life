package internal

import (
	"game-of-life/internal/assets"
	"game-of-life/internal/custom"
	"image/color"
	"strconv"

	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
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
	resDown *widget.Button
	resUp   *widget.Button
	resText *widget.Text
}

func (h *HUD) Init(actions Actions) {
	h.ui = new(ebitenui.UI)

	h.initContainer()
	h.initFPS()
	h.initButtons(actions)
	h.initResolutions(actions)
}

func (h *HUD) initContainer() {
	h.ui.Container = widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewStackedLayout(
			widget.StackedLayoutOpts.Padding(widget.NewInsetsSimple(4)),
		)),
	)
}

func (h *HUD) initFPS() {
	root := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	container := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(
			widget.AnchorLayoutOpts.Padding(widget.Insets{Top: 6}),
		)),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(
			widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionStart,
			},
		)),
	)

	var err error
	var face font.Face
	if face, err = assets.LoadFont(18); err != nil {
		panic(err)
	}

	h.fps = widget.NewLabel(
		widget.LabelOpts.Text("FPS", face, &widget.LabelColor{
			Idle: color.White,
		}),
		widget.LabelOpts.TextOpts(
			widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
		),
	)

	container.AddChild(h.fps)
	root.AddChild(container)
	h.ui.Container.AddChild(root)
}

func (h *HUD) initButtons(actions Actions) {
	h.play = custom.NewIconButton(assets.PlayIcon, actions.Play, custom.NewShortcut("Play the world", assets.KeySpaceIcon))
	h.pause = custom.NewIconButton(assets.PauseIcon, actions.Pause, custom.NewShortcut("Pause the world", assets.KeySpaceIcon))
	h.zoomIn = custom.NewIconButton(assets.PlusIcon, actions.ZoomIn, custom.NewShortcut("Increase cell's size", assets.KeyPlusIcon))
	h.zoomOut = custom.NewIconButton(assets.MinusIcon, actions.ZoomOut, custom.NewShortcut("Decrease cell's size", assets.KeyMinusIcon))
	h.fast = custom.NewIconButton(assets.FastIcon, actions.Fast, custom.NewShortcut("Increase generation speed", assets.KeyGreaterIcon))
	h.slow = custom.NewIconButton(assets.SlowIcon, actions.Slow, custom.NewShortcut("Decrease generation speed", assets.KeyLowerIcon))
	h.step = custom.NewIconButton(assets.StepIcon, actions.Step, custom.NewShortcut("Advance one generation", assets.KeyPeriodIcon))
	h.restart = custom.NewIconButton(assets.RestartIcon, actions.Restart, custom.NewShortcut("Regenerate the world", assets.KeyRIcon))

	root := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	container := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
			widget.RowLayoutOpts.Spacing(2),
		)),
		widget.ContainerOpts.WidgetOpts(
			widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionCenter,
			}),
		),
	)

	container.AddChild(h.restart)
	container.AddChild(h.play)
	container.AddChild(h.step)
	container.AddChild(h.pause)
	container.AddChild(h.zoomIn)
	container.AddChild(h.zoomOut)
	container.AddChild(h.slow)
	container.AddChild(h.fast)

	root.AddChild(container)
	h.ui.Container.AddChild(root)
}

func (h *HUD) initResolutions(actions Actions) {
	root := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
	)

	container := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout()),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(
			widget.AnchorLayoutData{
				HorizontalPosition: widget.AnchorLayoutPositionEnd,
			},
		)),
	)

	var resContainer *widget.Container
	h.resDown, resContainer, h.resUp, h.resText = custom.NewSelector("Change the resolution", state.res.String(), actions.ResolutionDown, actions.ResolutionUp)

	container.AddChild(resContainer)
	root.AddChild(container)
	h.ui.Container.AddChild(root)
}

func (h *HUD) updateFps() {
	h.fps.Label = "FPS: " + strconv.Itoa(int(ebiten.ActualFPS()))
}

func (h *HUD) UpdateResText() {
	h.resText.Label = state.res.String()
}

func (h *HUD) handleDisableButtons() {
	h.play.GetWidget().Disabled = !state.Paused()
	h.pause.GetWidget().Disabled = state.Paused()
	h.step.GetWidget().Disabled = !state.Paused()
	h.zoomIn.GetWidget().Disabled = !state.CanZoomIn()
	h.zoomOut.GetWidget().Disabled = !state.CanZoomOut()
	h.fast.GetWidget().Disabled = !state.CanFast()
	h.slow.GetWidget().Disabled = !state.CanSlow()
	h.resDown.GetWidget().Disabled = !state.CanResolutionDown()
	h.resUp.GetWidget().Disabled = !state.CanResolutionUp()
}

func (h *HUD) Update() {
	h.updateFps()
	h.handleDisableButtons()
	h.ui.Update()
}

func (h *HUD) Draw(screen *ebiten.Image) {
	h.ui.Draw(screen)
}
