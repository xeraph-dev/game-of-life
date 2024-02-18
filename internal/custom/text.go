package custom

import (
	"game-of-life/internal/assets"
	"image/color"

	"github.com/ebitenui/ebitenui/widget"
	"golang.org/x/image/font"
)

func NewText(text string) *widget.Text {
	var err error
	var face font.Face

	if face, err = assets.LoadFont(16); err != nil {
		panic(err)
	}

	return widget.NewText(
		widget.TextOpts.Text(text, face, color.White),
		widget.TextOpts.Position(widget.TextPositionCenter, widget.TextPositionCenter),
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(
			widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			},
		)),
	)
}
