package internal

type Cell struct {
	pixels    [4]*byte
	neighbors [8]*Cell
	alive     bool
	wasAlive  bool
}

func (self *Cell) Init(pixels [4]*byte, neighbors [8]*Cell) {
	self.pixels = pixels
	self.neighbors = neighbors
	self.alive = randomBool()
	self.wasAlive = self.alive
}

func (self *Cell) Update() (err error) {
	alives := 0
	for _, cell := range self.neighbors {
		if cell != nil && cell.wasAlive {
			alives++
		}
	}
	if self.alive {
		self.alive = alives == 2 || alives == 3
	} else {
		self.alive = alives == 3
	}
	return
}

func (self *Cell) Draw() {
	for _, pixel := range self.pixels {
		*pixel = self.byte()
	}
	self.wasAlive = self.alive
}

func (self Cell) byte() byte {
	if self.alive {
		return 0xff
	} else {
		return 0x0
	}
}
