package custom

import (
	"image"
	"image/color"

	eimage "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
)

func NewTooltip(content widget.PreferredSizeLocateableWidget) *widget.ToolTip {
	tooltipImage := eimage.NewNineSliceColor(color.NRGBA{R: 40, G: 40, B: 40, A: 255})

	container := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewAnchorLayout(
			widget.AnchorLayoutOpts.Padding(widget.Insets{Top: 4, Left: 8, Right: 8, Bottom: 4}),
		)),
		widget.ContainerOpts.AutoDisableChildren(),
		widget.ContainerOpts.BackgroundImage(tooltipImage),
	)

	container.AddChild(content)

	return widget.NewToolTip(
		widget.ToolTipOpts.Content(container),
		widget.ToolTipOpts.Position(widget.TOOLTIP_POS_WIDGET),
		widget.ToolTipOpts.Offset(image.Point{0, 4}),
		widget.ToolTipOpts.WidgetOriginHorizontal(widget.TOOLTIP_ANCHOR_MIDDLE),
		widget.ToolTipOpts.WidgetOriginVertical(widget.TOOLTIP_ANCHOR_END),
		widget.ToolTipOpts.ContentOriginHorizontal(widget.TOOLTIP_ANCHOR_MIDDLE),
		widget.ToolTipOpts.ContentOriginVertical(widget.TOOLTIP_ANCHOR_START),
	)
}
