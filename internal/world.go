package internal

type World struct {
	Pixels []byte
	cells  [][]*Cell
}

func (w *World) Init(width, height, zoom int) {
	offset := width * 40

	w.Pixels = make([]byte, width*height*4)

	w.cells = make([][]*Cell, height/zoom)
	for cy := range w.cells {
		w.cells[cy] = make([]*Cell, width/zoom)
		for cx := range w.cells[cy] {
			w.cells[cy][cx] = NewCell(zoom)
		}
	}

	for cy, cells := range w.cells {
		for cx, cell := range cells {
			for py := range cell.pixels {
				for px := range cell.pixels[py] {
					pos := py*width + cy*width*zoom + px + cx*zoom
					if pos < offset {
						continue
					}
					cell.pixels[py][px] = NewPixel()
					pixel := cell.pixels[py][px]
					pixel.pixels = w.Pixels[pos*4 : pos*4+4]
				}
			}

			ni := 0
			for y := cy - 1; y <= cy+1; y++ {
				for x := cx - 1; x <= cx+1; x++ {
					if x == cx && y == cy {
						continue
					}
					if y >= 0 && x >= 0 && y < len(w.cells) && x < len(w.cells[cy]) {
						cell.neighbors[ni] = w.cells[y][x]
					}
					ni++
				}
			}

		}
	}
}

func (w *World) Update() (err error) {
	for _, cells := range w.cells {
		for _, cell := range cells {
			if err = cell.Update(); err != nil {
				return
			}
		}
	}
	return
}

func (w *World) Draw() {
	for _, cells := range w.cells {
		for _, cell := range cells {
			cell.Draw()
		}
	}
}
