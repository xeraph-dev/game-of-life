package internal

type Pixel struct {
	pixels []byte
}

func (p *Pixel) White() {
	for i := range p.pixels {
		p.pixels[i] = 0xff
	}
}

func (p *Pixel) Black() {
	for i := range p.pixels {
		p.pixels[i] = 0x0
	}
}
