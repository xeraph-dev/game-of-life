package internal

import (
	"math/rand"
)

type Cell struct {
	pixels    [][]Pixel
	neighbors [8]*Cell
	alive     bool
	wasAlive  bool
}

func (c *Cell) Init(pixels [][]Pixel, neighbors [8]*Cell) {
	c.alive = rand.Intn(100) < 10
	c.wasAlive = c.alive

	c.pixels = pixels
	c.neighbors = neighbors
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
			if c.alive {
				px.White()
			} else {
				px.Black()
			}
		}
	}
	c.wasAlive = c.alive
}
