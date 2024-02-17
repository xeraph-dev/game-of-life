package internal

type World struct {
	Pixels []byte
	cells  [][]Cell
}

func (w *World) Init(width, height, zoom int) {
	offset := width * 40

	w.Pixels = make([]byte, width*height*4)

	w.cells = make([][]Cell, height/zoom)
	for cy := range w.cells {
		w.cells[cy] = make([]Cell, width/zoom)
	}

	for cy := range w.cells {
		for cx := range w.cells[cy] {
			var neighbors [8]*Cell
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

			pixels := make([][]Pixel, zoom)
			for py := range pixels {
				pixels[py] = make([]Pixel, zoom)
			}

			for py := range pixels {
				for px := range pixels[py] {
					pos := py*width + cy*width*zoom + px + cx*zoom
					if pos < offset {
						continue
					}
					pixels[py][px].pixels = w.Pixels[pos*4 : pos*4+4]
				}
			}

			w.cells[cy][cx].Init(pixels, neighbors)
		}
	}
}

func (w *World) Update() {
	for cy := range w.cells {
		for cx := range w.cells[cy] {
			w.cells[cy][cx].Update()
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
