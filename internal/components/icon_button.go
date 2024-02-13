package custom

import (
	"image/color"

	"github.com/ebitenui/ebitenui/image"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
)

func NewIconButton() *widget.Button {
	ttfFont, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic(err)
	}

	face := truetype.NewFace(ttfFont, &truetype.Options{
		Size:    8,
		DPI:     50,
		Hinting: font.HintingFull,
	})

	buttonImage := &widget.ButtonImage{
		Idle:    image.NewNineSliceColor(color.NRGBA{R: 170, G: 170, B: 180, A: 255}),
		Hover:   image.NewNineSliceColor(color.NRGBA{R: 130, G: 130, B: 150, A: 255}),
		Pressed: image.NewNineSliceColor(color.NRGBA{R: 100, G: 100, B: 120, A: 255}),
	}

	buttonTextColor := &widget.ButtonTextColor{
		Idle: color.NRGBA{0xdf, 0xf4, 0xff, 0xff},
	}

	padding := widget.Insets{
		Left:   5,
		Right:  5,
		Top:    5,
		Bottom: 5,
	}

	clickedHandler := func(args *widget.ButtonClickedEventArgs) {
		println("button clicked")
	}

	return widget.NewButton(
		widget.ButtonOpts.Image(buttonImage),
		widget.ButtonOpts.Text("Play", face, buttonTextColor),
		widget.ButtonOpts.TextPadding(padding),
		widget.ButtonOpts.ClickedHandler(clickedHandler),
	)
}
