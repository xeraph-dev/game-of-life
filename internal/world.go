package internal

type World struct {
	Pixels                 []byte
	cells                  [][]Cell
	offset                 int
	virtualOffset          int
	yMin, yMax, xMin, xMax int
}

func (w *World) Init() {
	res := state.Resolution()
	width := res.Width
	height := res.Height
	zoom := state.Zoom()

	w.offset = width * 40
	w.virtualOffset = 10

	w.initPixels(width, height)
	w.initCells(width, height, zoom)

	w.computeVirtualConstraints()

	for cy := range w.cells {
		for cx := range w.cells[cy] {
			neighbors := w.cellNeighbors(cy, cx)
			pixels := w.cellPixels(cy, cx, width, zoom)
			w.cells[cy][cx].Init(pixels, neighbors)
		}
	}
}

func (w *World) initPixels(width, height int) {
	w.Pixels = make([]byte, width*height*4)
}

func (w *World) initCells(width, height, zoom int) {
	w.cells = make([][]Cell, height/zoom+w.virtualOffset*2)
	for cy := range w.cells {
		w.cells[cy] = make([]Cell, width/zoom+w.virtualOffset*2)
	}
}

func (w *World) computeVirtualConstraints() {
	w.yMin = w.virtualOffset
	w.yMax = len(w.cells) - w.virtualOffset
	w.xMin = w.virtualOffset
	w.xMax = len(w.cells[0]) - w.virtualOffset
}

func (w *World) cellIsVirtual(cy, cx int) bool {
	return cy < w.yMin || cx < w.xMin || cy >= w.yMax || cx >= w.xMax
}

func (w *World) initCellPixels(zoom int) (pixels [][]Pixel) {
	pixels = make([][]Pixel, zoom)
	for py := range pixels {
		pixels[py] = make([]Pixel, zoom)
	}
	return
}

func (w *World) cellPixels(cy, cx, width, zoom int) (pixels [][]Pixel) {
	pixels = w.initCellPixels(zoom)

	if w.cellIsVirtual(cy, cx) {
		return
	}

	cy -= w.virtualOffset
	cx -= w.virtualOffset

	for py := range pixels {
		for px := range pixels[py] {
			pos := py*width + cy*width*zoom + px + cx*zoom
			if pos < w.offset {
				continue
			}
			pixels[py][px].pixels = w.Pixels[pos*4 : pos*4+4]
		}
	}
	return
}

func (w *World) cellNeighbors(cy, cx int) (neighbors [8]*Cell) {
	ni := 0
	for y := cy - 1; y <= cy+1; y++ {
		for x := cx - 1; x <= cx+1; x++ {
			if x == cx && y == cy {
				continue
			}
			if y >= 0 && x >= 0 && y < len(w.cells) && x < len(w.cells[cy]) {
				neighbors[ni] = &w.cells[y][x]
			}
			ni++
		}
	}

	return
}

func (w *World) Update() {
	for cy := range w.cells {
		for cx := range w.cells[cy] {
			if w.cellIsVirtual(cy, cx) {
				w.cells[cy][cx].InVirtual()
			} else {
				w.cells[cy][cx].OutVirtual()
			}

			w.cells[cy][cx].Update(w.virtualOffset)
		}
	}
}

func (w *World) Draw() {
	for cy := range w.cells {
		for cx := range w.cells[cy] {
			w.cells[cy][cx].Draw()
		}
	}
}
