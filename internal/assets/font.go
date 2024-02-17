package assets

import (
	"fmt"
	"sync"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/goregular"
)

var ttfFont *truetype.Font
var face font.Face
var fontM sync.RWMutex

func LoadFont(size float64) (font.Face, error) {
	fontM.Lock()
	defer fontM.Unlock()
	var err error
	if ttfFont == nil {
		if ttfFont, err = truetype.Parse(goregular.TTF); err != nil {
			err = fmt.Errorf("loading ttf font: %w", err)
			return nil, err
		}
	}
	if face != nil {
		return face, err
	}

	opts := &truetype.Options{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingFull,
	}
	face = truetype.NewFace(ttfFont, opts)

	return face, err
}
