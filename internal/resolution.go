package internal

import (
	"fmt"
	"slices"

	"github.com/hajimehoshi/ebiten/v2"
)

type Resolution struct {
	Width, Height int
}

func NewResolution(width, height int) (r Resolution) {
	r.Width = width
	r.Height = height
	return
}

func NewResolutionFromString(s string) (r Resolution, err error) {
	if _, err = fmt.Sscanf(s, "%dx%d", &r.Width, &r.Height); err != nil {
		err = fmt.Errorf("invalid resolution format: %w", err)
		return
	}
	return
}

func (r Resolution) String() string {
	return fmt.Sprintf("%dx%d", r.Width, r.Height)
}

func (r Resolution) Greater(r2 Resolution) bool {
	return r.Width > r2.Width && r.Height > r2.Height
}

func (r Resolution) Lower(r2 Resolution) bool {
	return r.Width < r2.Width && r.Height < r2.Height
}

func (r *Resolution) Up() {
	for _, str := range ResolutionList {
		ro := Resolutions[str]
		if r.Width < ro.Width && r.Height < ro.Height {
			r.Width = ro.Width
			r.Height = ro.Height
			return
		}
	}
}

func (r *Resolution) Down() {
	list := slices.Clone(ResolutionList)
	slices.Reverse(list)
	for _, str := range list {
		ro := Resolutions[str]
		if r.Width > ro.Width && r.Height > ro.Height {
			r.Width = ro.Width
			r.Height = ro.Height
			return
		}
	}
}

func InitializeResolutionsMap(rs []Resolution) (adjust map[string]Resolution, list []string) {
	adjust = make(map[string]Resolution)
	width, height := ebiten.ScreenSizeInFullscreen()
	for _, r := range rs {
		if r.Width <= width && r.Height <= height {
			adjust[r.String()] = r
			list = append(list, r.String())
		}
	}
	return adjust, list
}
