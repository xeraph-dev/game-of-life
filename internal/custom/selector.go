package custom

import (
	"game-of-life/internal/assets"

	"github.com/ebitenui/ebitenui/widget"
)

func NewSelector(tooltip string, content string, onDown func(), onUp func()) (*widget.Button, *widget.Container, *widget.Button, *widget.Text) {
	container := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
			widget.RowLayoutOpts.Spacing(8),
		)),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.ToolTip(NewTooltip(NewText(tooltip)))),
	)

	leftButton := NewIconButton(assets.LeftIcon, func() {
		onDown()
	}, nil)

	rightButton := NewIconButton(assets.RightIcon, func() {
		onUp()
	}, nil)

	text := NewText(content)

	container.AddChild(leftButton)
	container.AddChild(text)
	container.AddChild(rightButton)

	return leftButton, container, rightButton, text
}
