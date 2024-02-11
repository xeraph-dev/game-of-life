package internal

type World struct {
	cells  [][]Cell
	pixels []byte
}

func (self *World) Init() {
	width := ScreenWidth / state.zoom
	height := ScreenHeight / state.zoom

	self.cells = make([][]Cell, height)
	for i := range self.cells {
		self.cells[i] = make([]Cell, width)
	}
	self.pixels = make([]byte, width*height*4)

	for cy := range self.cells {
		for cx := range self.cells[cy] {
			pixels := [4]*byte{}

			for i := 0; i < 4; i++ {
				pixels[i] = &self.pixels[cy*len(self.cells[cy])*4+cx*4+i]
			}

			neighbors := [8]*Cell{}
			ni := 0
			for y := cy - 1; y <= cy+1; y++ {
				for x := cx - 1; x <= cx+1; x++ {
					if x == cx && y == cy {
						continue
					}
					if y >= 0 && x >= 0 && y < len(self.cells) && x < len(self.cells[cy]) {
						neighbors[ni] = &self.cells[y][x]
					}
					ni++
				}
			}

			self.cells[cy][cx].Init(pixels, neighbors)
		}
	}
}

func (self World) Pixels() []byte {
	return self.pixels
}

func (self *World) Update() (err error) {
	for cy := range self.cells {
		for cx := range self.cells[cy] {
			self.cells[cy][cx].Update()
		}
	}
	return
}

func (self *World) Draw() {
	for cy := range self.cells {
		for cx := range self.cells[cy] {
			self.cells[cy][cx].Draw()
		}
	}
}
