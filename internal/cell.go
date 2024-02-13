package internal

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Cell struct {
	pixels    [][]*Pixel
	neighbors [8]*Cell
	alive     bool
	wasAlive  bool
}

func NewCell() (c *Cell) {
	c = new(Cell)

	c.alive = rand.Intn(100) < 10
	c.wasAlive = c.alive

	c.pixels = make([][]*Pixel, state.Zoom)
	for py := range c.pixels {
		c.pixels[py] = make([]*Pixel, state.Zoom)
	}

	return
}

func (c *Cell) Update() (err error) {
	alive := 0
	for _, cell := range c.neighbors {
		if cell != nil && cell.wasAlive {
			alive++
		}
	}
	if c.alive {
		c.alive = alive == 2 || alive == 3
	} else {
		c.alive = alive == 3
	}
	return
}

func (c *Cell) Draw() {
	for _, pixels := range c.pixels {
		for _, px := range pixels {
			if px == nil {
				continue
			}
			if c.alive {
				px.White()
			} else {
				px.Black()
			}
		}
	}
	c.wasAlive = c.alive
}
