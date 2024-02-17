package custom

import (
	"bytes"
	"game-of-life/internal/assets"
	"image"
	_ "image/png"

	eimage "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewIconButton(icon assets.Icon, clickedHandler func(), toolTipContent widget.PreferredSizeLocateableWidget) *widget.Button {
	var img image.Image
	var err error

	if img, _, err = image.Decode(bytes.NewReader(icon.Idle)); err != nil {
		panic(err)
	}
	idleImage := ebiten.NewImageFromImage(img)
	iconSize := idleImage.Bounds().Dx()
	buttonSize := iconSize * 2

	if img, _, err = image.Decode(bytes.NewReader(icon.Disabled)); err != nil {
		panic(err)
	}
	disabledImage := ebiten.NewImageFromImage(img)

	buttonImage := &widget.ButtonImage{
		Idle:     eimage.NewNineSliceSimple(idleImage, 0, iconSize),
		Hover:    eimage.NewNineSliceSimple(idleImage, 0, iconSize),
		Pressed:  eimage.NewNineSliceSimple(idleImage, 0, iconSize),
		Disabled: eimage.NewNineSliceSimple(disabledImage, 0, iconSize),
	}

	button := widget.NewButton(
		widget.ButtonOpts.Image(buttonImage),
		widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) { clickedHandler() }),
		widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			MaxWidth:  buttonSize,
			MaxHeight: buttonSize,
		})),
	)

	if toolTipContent != nil {
		tooltip := NewTooltip(toolTipContent)
		button.Configure(widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.ToolTip(tooltip)))
	}

	return button
}
