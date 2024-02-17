package custom

import (
	"bytes"
	"game-of-life/internal/assets"
	"image"
	"image/color"

	eimage "github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

func NewShortcut(text string, icon assets.Icon) *widget.Container {
	container := widget.NewContainer(
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionHorizontal),
			widget.RowLayoutOpts.Spacing(8),
		)),
	)

	container.AddChild(NewShortcutLabel(text))
	container.AddChild(NewShortcutIcon(icon))

	return container
}

func NewShortcutLabel(text string) *widget.Text {

	var err error
	var face font.Face

	if face, err = assets.LoadFont(16); err != nil {
		panic(err)
	}

	return widget.NewText(
		widget.TextOpts.Text(text, face, color.White),
		widget.TextOpts.WidgetOpts(widget.WidgetOpts.LayoutData(
			widget.RowLayoutData{
				Position: widget.RowLayoutPositionCenter,
			},
		)),
	)
}

func NewShortcutIcon(icon assets.Icon) *widget.Container {
	var err error
	var img image.Image

	if img, _, err = image.Decode(bytes.NewReader(icon.Idle)); err != nil {
		panic(err)
	}
	idleImage := ebiten.NewImageFromImage(img)
	iconSize := idleImage.Bounds().Dx()
	image := eimage.NewNineSliceSimple(idleImage, 0, iconSize)

	return widget.NewContainer(
		widget.ContainerOpts.BackgroundImage(image),
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(
			widget.RowLayoutData{
				Position:  widget.RowLayoutPositionCenter,
				MaxHeight: 28,
				MaxWidth:  28,
			},
		)),
	)
}
